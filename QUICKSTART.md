# Quick Start Guide for New Client Projects

This guide will help you quickly set up this template for a new client project.

## 1. Clone the Template

```bash
# Clone your template repository
git clone <your-template-repo-url> client-project-name
cd client-project-name

# Remove the old git history and start fresh
rm -rf .git
git init
git add .
git commit -m "Initial commit from template"
```

## 2. Customize for Your Client

### Update Project Name

**Option A: Use the Makefile (macOS/Linux)**
```bash
make setup-client CLIENT=clientname
```

**Option B: Manual Updates**
1. Edit `terraform/variables.tf` - change `project_name` default value
2. Edit `backend/main.go` - change `AppName` in Fiber config
3. Edit `frontend/index.html` - update branding, company name, colors

### Update Environment Variables

```bash
cp .env.example .env
# Edit .env with client-specific values:
# - Database credentials
# - AWS settings
# - Application secrets
```

### Customize Frontend

Edit `frontend/index.html`:
- Update `<title>` tag
- Change company name in navigation
- Modify Tailwind colors to match client brand
- Update footer text

Example color scheme update:
```html
<!-- Change from blue to client color (e.g., purple) -->
<h1 class="text-2xl font-bold text-purple-600">Client Name</h1>
<button class="bg-purple-600 hover:bg-purple-700 ...">
```

## 3. Test Locally

```bash
# Start the development environment
make dev

# Or manually:
docker-compose up
```

Visit: http://localhost:3000

## 4. Deploy to AWS

### First Time Setup

```bash
# Configure AWS credentials
aws configure

# Initialize Terraform
make init-terraform

# Review what will be created
make plan
```

### Deploy

**Option A: One-Command Deploy (Recommended)**
```bash
make deploy
```

**Option B: Step-by-Step Deploy**
```bash
# 1. Deploy infrastructure
make deploy-infra-only

# 2. Build and push Docker image
make push

# 3. Update ECS service
make update-service
```

### Access Your Application

```bash
# Get the URL
make url
```

## 5. Common Customizations

### Remove Database (if not needed)

1. **Docker Compose**: Remove the `postgres` service from `docker-compose.yml`
2. **Terraform**: Set `enable_database = false` in `terraform/variables.tf`
3. **Backend**: Comment out database initialization in `backend/main.go`:
   ```go
   // initDatabase()
   ```

### Add Custom API Endpoints

1. Create a new handler in `backend/handlers/`:
   ```go
   package handlers
   
   import "github.com/gofiber/fiber/v2"
   
   func MyCustomHandler(c *fiber.Ctx) error {
       return c.JSON(fiber.Map{"message": "Custom response"})
   }
   ```

2. Register the route in `backend/main.go`:
   ```go
   api.Get("/custom", handlers.MyCustomHandler)
   ```

### Update Database Model

1. Edit `backend/models/item.go` or create new model files
2. Add auto-migration in `backend/main.go`:
   ```go
   DB.AutoMigrate(&models.YourNewModel{})
   ```

## 6. Cost Optimization

### Development/Small Projects
Edit `terraform/variables.tf`:
```hcl
ecs_cpu           = "256"
ecs_memory        = "512"
ecs_desired_count = 1
db_instance_class = "db.t3.micro"  # or db.t4g.micro
```

### Production/Higher Traffic
```hcl
ecs_cpu           = "512"
ecs_memory        = "1024"
ecs_desired_count = 2
db_instance_class = "db.t3.small"
```

## 7. Maintenance Commands

```bash
# View logs
make logs

# Check service status
make status

# Update application (after code changes)
make push && make update-service

# Destroy everything (be careful!)
make destroy
```

## 8. Adding Custom Domain (Optional)

1. Buy domain and add to Route53
2. Uncomment domain variables in `terraform/variables.tf`
3. Uncomment Route53 and ACM sections in `terraform/main.tf`
4. Update the variable:
   ```hcl
   domain_name = "client-domain.com"
   ```
5. Run `terraform apply`

## 9. Set Up CI/CD (Optional)

Create `.github/workflows/deploy.yml`:
```yaml
name: Deploy to AWS

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v1
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: us-east-1
      - name: Deploy
        run: make deploy
```

## 10. Checklist Before Client Handoff

- [ ] Update all branding in frontend
- [ ] Configure proper database passwords (not defaults!)
- [ ] Set up monitoring and alerts
- [ ] Enable HTTPS with SSL certificate
- [ ] Configure backup policies
- [ ] Set up proper IAM permissions
- [ ] Document any custom features
- [ ] Test all CRUD operations
- [ ] Load test the application
- [ ] Set up CI/CD pipeline

## Need Help?

- Check the main `README.md` for detailed documentation
- Review Terraform outputs: `cd terraform && terraform output`
- Check logs: `make logs`
- Verify service health: `make status`

## Estimated Costs (us-east-1)

**Minimal Setup** (1 task, db.t3.micro):
- ECS Fargate: ~$10-15/month
- RDS t3.micro: ~$15-20/month
- ALB: ~$20/month
- Data transfer: ~$5/month
- **Total: ~$50-60/month**

**No Database** (1 task):
- ECS Fargate: ~$10-15/month
- ALB: ~$20/month
- Data transfer: ~$5/month
- **Total: ~$35-40/month**

Costs will vary based on traffic and usage.
