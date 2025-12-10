# Troubleshooting Guide

Common issues and their solutions when using this template.

## Local Development Issues

### Docker Compose Fails to Start

**Problem**: `docker-compose up` fails with network errors

**Solution**:
```bash
# Remove existing containers and networks
docker-compose down -v

# Rebuild without cache
docker-compose up --build --force-recreate
```

### Database Connection Refused

**Problem**: Backend can't connect to PostgreSQL

**Solutions**:
1. Check if PostgreSQL container is running:
   ```bash
   docker-compose ps
   ```

2. Verify DATABASE_URL in `.env`:
   ```bash
   DATABASE_URL=postgres://appuser:change_this_password@postgres:5432/appdb?sslmode=disable
   ```

3. Wait for PostgreSQL to be ready:
   ```bash
   docker-compose logs postgres
   # Look for "database system is ready to accept connections"
   ```

### Port Already in Use

**Problem**: "Port 3000 is already allocated"

**Solution**:
```bash
# Find and kill process using port 3000
lsof -ti:3000 | xargs kill -9

# Or change the port in docker-compose.yml
ports:
  - "3001:3000"  # Use 3001 externally
```

## AWS Deployment Issues

### Terraform Apply Fails

**Problem**: "Error: error creating ECS Service"

**Solutions**:

1. **Check AWS credentials**:
   ```bash
   aws sts get-caller-identity
   ```

2. **Verify sufficient permissions**: Your AWS user needs permissions for:
   - VPC, EC2, ECS, RDS, ECR, IAM, CloudWatch, Application Load Balancer

3. **Check for resource limits**:
   ```bash
   # Check ECS limits
   aws service-quotas get-service-quota --service-code ecs --quota-code L-3032A538
   ```

4. **Clean up partial deployment**:
   ```bash
   cd terraform
   terraform destroy
   terraform apply
   ```

### ECR Push Permission Denied

**Problem**: "denied: Your authorization token has expired"

**Solution**:
```bash
# Re-authenticate with ECR
aws ecr get-login-password --region us-east-1 | \
  docker login --username AWS --password-stdin <account-id>.dkr.ecr.us-east-1.amazonaws.com
```

### ECS Tasks Keep Stopping

**Problem**: Tasks start but immediately stop

**Solutions**:

1. **Check CloudWatch logs**:
   ```bash
   make logs
   # or
   aws logs tail /ecs/web-template --follow
   ```

2. **Common issues**:
   - Missing environment variables
   - Database connection failure
   - Application crashes on startup
   - Out of memory (increase `ecs_memory`)

3. **Check task definition**:
   ```bash
   aws ecs describe-task-definition --task-definition web-template
   ```

4. **Check stopped task reason**:
   ```bash
   aws ecs describe-tasks --cluster web-template-cluster --tasks <task-arn>
   ```

### Health Check Failing

**Problem**: ALB marks targets as unhealthy

**Solutions**:

1. **Verify health endpoint**:
   ```bash
   # SSH into ECS task or test locally
   curl http://localhost:3000/api/health
   ```

2. **Increase health check timeouts** in `terraform/main.tf`:
   ```hcl
   health_check {
     timeout             = 60
     interval            = 300
     healthy_threshold   = 2
     unhealthy_threshold = 10
   }
   ```

3. **Check security groups**: Ensure ALB can reach ECS tasks on port 3000

### Database Connection Timeout in Production

**Problem**: Application can't connect to RDS

**Solutions**:

1. **Check security groups**:
   - RDS security group should allow inbound on 5432 from ECS security group
   - Verify in AWS Console: EC2 > Security Groups

2. **Verify RDS endpoint**:
   ```bash
   cd terraform
   terraform output database_endpoint
   ```

3. **Check DATABASE_URL format**:
   ```
   postgres://username:password@endpoint:5432/dbname?sslmode=require
   ```

4. **Test connection from ECS task**:
   ```bash
   # Get task ID
   aws ecs list-tasks --cluster web-template-cluster
   
   # Execute command in task
   aws ecs execute-command --cluster web-template-cluster \
     --task <task-id> --container web-template \
     --command "/bin/sh" --interactive
   ```

## Performance Issues

### Slow Application Response

**Solutions**:

1. **Increase ECS resources** in `terraform/variables.tf`:
   ```hcl
   ecs_cpu    = "512"
   ecs_memory = "1024"
   ```

2. **Scale horizontally**:
   ```hcl
   ecs_desired_count = 2
   ```

3. **Check database performance**:
   - Upgrade RDS instance class
   - Enable connection pooling
   - Add database indexes

4. **Enable CloudFront** for static assets (not included in template)

### High AWS Costs

**Solutions**:

1. **Use smaller instances**:
   ```hcl
   ecs_cpu           = "256"
   ecs_memory        = "512"
   db_instance_class = "db.t4g.micro"
   ```

2. **Use Aurora Serverless** instead of standard RDS

3. **Reduce desired task count**:
   ```hcl
   ecs_desired_count = 1
   ```

4. **Set up auto-scaling** based on CPU/memory

5. **Delete unused resources**:
   ```bash
   # Remove old ECR images
   aws ecr list-images --repository-name web-template
   aws ecr batch-delete-image --repository-name web-template --image-ids imageTag=old-tag
   ```

## Frontend Issues

### HTMX Requests Not Working

**Problem**: HTMX requests return errors

**Solutions**:

1. **Check CORS configuration** in `backend/main.go`:
   ```go
   app.Use(cors.New(cors.Config{
       AllowOrigins: "*",
       AllowHeaders: "Origin, Content-Type, Accept",
   }))
   ```

2. **Verify API endpoints**:
   ```bash
   curl http://localhost:3000/api/items
   ```

3. **Check browser console** for JavaScript errors

### Styles Not Loading

**Problem**: Tailwind CSS not working

**Solutions**:

1. **Check CDN connection**:
   - Ensure internet connectivity
   - Try using a different CDN or self-hosted Tailwind

2. **Verify HTML structure**:
   ```html
   <script src="https://cdn.tailwindcss.com"></script>
   ```

## Database Issues

### Migration Errors

**Problem**: GORM auto-migration fails

**Solutions**:

1. **Check database permissions**:
   ```sql
   GRANT ALL PRIVILEGES ON DATABASE appdb TO appuser;
   ```

2. **Manually create tables**:
   ```bash
   # Connect to database
   docker-compose exec postgres psql -U appuser -d appdb
   
   # Check tables
   \dt
   ```

3. **Drop and recreate database** (development only):
   ```bash
   docker-compose down -v
   docker-compose up
   ```

### Database Out of Disk Space

**Problem**: RDS runs out of storage

**Solutions**:

1. **Increase allocated storage** in `terraform/main.tf`:
   ```hcl
   allocated_storage = 50  # Increase from 20
   ```

2. **Enable storage autoscaling**:
   ```hcl
   max_allocated_storage = 100
   ```

3. **Clean up old data**

## Debugging Tips

### Enable Verbose Logging

**Backend** (`backend/main.go`):
```go
app.Use(logger.New(logger.Config{
    Format: "[${time}] ${status} - ${method} ${path}\n",
}))
```

**Terraform**:
```bash
export TF_LOG=DEBUG
terraform apply
```

### Access ECS Task

```bash
# List tasks
aws ecs list-tasks --cluster web-template-cluster

# Get task details
aws ecs describe-tasks --cluster web-template-cluster --tasks <task-arn>

# Execute command (requires ECS Exec enabled)
aws ecs execute-command --cluster web-template-cluster \
  --task <task-id> --container web-template \
  --command "/bin/sh" --interactive
```

### Check Docker Image

```bash
# List images in ECR
aws ecr describe-images --repository-name web-template

# Pull and test locally
docker pull <ecr-url>:latest
docker run -p 3000:3000 -e APP_PORT=3000 <ecr-url>:latest
```

## Getting Help

1. **Check logs first**:
   ```bash
   make logs
   docker-compose logs
   ```

2. **Verify resources**:
   ```bash
   make status
   ```

3. **Review Terraform state**:
   ```bash
   cd terraform
   terraform show
   ```

4. **Test individual components**:
   - Backend: `cd backend && go run main.go`
   - Frontend: Open `frontend/index.html` in browser
   - Database: `docker-compose exec postgres psql -U appuser`

## Emergency Recovery

### Complete Reset (Development)

```bash
# Stop everything
docker-compose down -v

# Clean Docker
docker system prune -a

# Clean Terraform
cd terraform
rm -rf .terraform terraform.tfstate*

# Start fresh
terraform init
terraform apply
```

### Complete Reset (AWS)

```bash
# Destroy all AWS resources
make destroy

# Wait for complete deletion, then redeploy
make deploy
```

**Warning**: This will delete all data!
