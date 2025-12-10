# Architecture Overview

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         Internet                            │
└────────────────────┬────────────────────────────────────────┘
                     │
                     ▼
         ┌───────────────────────┐
         │  Application Load     │
         │  Balancer (ALB)       │
         │  Port 80/443          │
         └───────────┬───────────┘
                     │
                     ▼
         ┌───────────────────────┐
         │   Target Group        │
         │   Health Check:       │
         │   /api/health         │
         └───────────┬───────────┘
                     │
         ┌───────────┴───────────┐
         │                       │
         ▼                       ▼
┌─────────────────┐    ┌─────────────────┐
│  ECS Fargate    │    │  ECS Fargate    │
│  Task 1         │    │  Task 2         │
│                 │    │  (Optional)     │
│  ┌───────────┐  │    │  ┌───────────┐  │
│  │ Go Backend│  │    │  │ Go Backend│  │
│  │ + Frontend│  │    │  │ + Frontend│  │
│  │ Port 3000 │  │    │  │ Port 3000 │  │
│  └─────┬─────┘  │    │  └─────┬─────┘  │
└────────┼────────┘    └────────┼────────┘
         │                      │
         └──────────┬───────────┘
                    │
                    ▼
         ┌───────────────────────┐
         │   RDS PostgreSQL      │
         │   (Optional)          │
         │   Port 5432           │
         └───────────────────────┘
```

## Technology Stack

### Frontend Layer
- **HTMX**: Enables dynamic HTML updates without heavy JavaScript
- **Alpine.js**: Lightweight JavaScript for interactivity (~15KB)
- **Tailwind CSS**: Utility-first CSS framework (CDN)

### Backend Layer
- **Go 1.21**: High-performance compiled language
- **Fiber Framework**: Express-inspired web framework for Go
- **GORM**: ORM for database operations

### Database Layer (Optional)
- **PostgreSQL 16**: Relational database
- **Managed by**: AWS RDS

### Infrastructure Layer
- **Docker**: Containerization
- **AWS ECS Fargate**: Serverless container orchestration
- **Application Load Balancer**: Traffic distribution
- **ECR**: Container registry
- **CloudWatch**: Logging and monitoring

## Data Flow

### Request Flow
1. User makes request → ALB
2. ALB forwards to healthy ECS task
3. Go backend receives request
4. If static file → Serve from `/static`
5. If API call → Process and query database (if needed)
6. Return response to user

### HTMX Flow
1. User interaction triggers HTMX
2. HTMX sends AJAX request to API
3. Server returns JSON
4. JavaScript transforms JSON to HTML
5. HTMX swaps HTML into DOM

## Deployment Flow

```
Local Development
       │
       ▼
Docker Build
       │
       ▼
Push to ECR
       │
       ▼
ECS Task Definition Updated
       │
       ▼
ECS Service Deploys New Tasks
       │
       ▼
ALB Health Checks Pass
       │
       ▼
Traffic Routed to New Tasks
       │
       ▼
Old Tasks Drained & Stopped
```

## Network Architecture

### VPC Structure
- **CIDR**: 10.0.0.0/16
- **Public Subnets**: 
  - 10.0.1.0/24 (AZ1)
  - 10.0.2.0/24 (AZ2)
- **Private Subnets**: 
  - 10.0.10.0/24 (AZ1)
  - 10.0.11.0/24 (AZ2)

### Security Groups

**ALB Security Group**
- Inbound: 80, 443 from 0.0.0.0/0
- Outbound: All

**ECS Security Group**
- Inbound: 3000 from ALB SG
- Outbound: All

**RDS Security Group** (if enabled)
- Inbound: 5432 from ECS SG
- Outbound: All

## Scaling Strategy

### Vertical Scaling
Increase resources in `terraform/variables.tf`:
```hcl
ecs_cpu    = "512"  # or 1024, 2048, 4096
ecs_memory = "1024" # or 2048, 4096, 8192
```

### Horizontal Scaling
Increase task count:
```hcl
ecs_desired_count = 2  # or more
```

### Auto-Scaling (Not Included)
Can be added with ECS Service Auto Scaling:
- Target Tracking: CPU or Memory utilization
- Step Scaling: Custom CloudWatch metrics
- Scheduled Scaling: Predictable patterns

## High Availability

### Current Setup
- Multi-AZ deployment
- ALB distributes traffic across AZs
- Database in private subnet (if enabled)

### Production Recommendations
- Enable RDS Multi-AZ
- Use at least 2 ECS tasks
- Configure auto-scaling
- Set up CloudWatch alarms
- Enable container insights

## Security Features

### Network Security
- Private subnets for database
- Security groups with least privilege
- VPC isolation

### Application Security
- CORS configuration in backend
- HTTPS ready (add ACM certificate)
- Environment variables for secrets
- No hardcoded credentials

### Recommendations
- Use AWS Secrets Manager
- Enable WAF on ALB
- Set up AWS Shield
- Enable VPC Flow Logs
- Use IAM roles, not access keys

## Monitoring & Logging

### CloudWatch Logs
- Log Group: `/ecs/{project_name}`
- Retention: 7 days (configurable)
- Streams: Per container

### Metrics
- ECS: CPU, Memory utilization
- ALB: Request count, latency
- RDS: Connections, queries

### Recommended Alerts
- ECS task stopped unexpectedly
- ALB unhealthy target count
- High CPU/Memory usage
- Database connection errors

## Cost Breakdown

### Fixed Costs
- **ALB**: ~$20/month
- **ECR Storage**: ~$1-5/month
- **CloudWatch Logs**: ~$1-5/month

### Variable Costs
- **ECS Fargate**: 
  - 256 CPU, 512 MB: ~$10-15/month per task
  - 512 CPU, 1024 MB: ~$20-30/month per task
- **RDS**:
  - db.t3.micro: ~$15-20/month
  - db.t3.small: ~$30-40/month
- **Data Transfer**: Varies by traffic

### Total (Minimal Setup)
1 ECS task (256/512) + db.t3.micro: **~$50-60/month**

### Without Database
1 ECS task (256/512): **~$35-40/month**

## Disaster Recovery

### Backup Strategy
- **Database**: Automated RDS snapshots (configurable)
- **Application**: Infrastructure as Code (Terraform)
- **Container Images**: Versioned in ECR

### Recovery Procedures
1. **Database**: Restore from RDS snapshot
2. **Application**: Redeploy from ECR image
3. **Infrastructure**: Re-apply Terraform

### RTO/RPO
- **RTO** (Recovery Time): ~15-30 minutes
- **RPO** (Point): Depends on backup frequency

## Performance Considerations

### Optimization Tips
1. Use connection pooling for database
2. Enable HTTP/2 on ALB
3. Implement caching (Redis/ElastiCache)
4. Use CloudFront for static assets
5. Optimize Docker image size
6. Enable compression in Go

### Benchmarks (Expected)
- **Single Task (256/512)**:
  - ~500-1000 req/sec
  - ~50-100ms latency
- **Load Balancing (2+ tasks)**:
  - Scales linearly with task count
  - Better fault tolerance

## Development vs Production

### Development
```hcl
ecs_desired_count = 1
db_instance_class = "db.t3.micro"
skip_final_snapshot = true
```

### Production
```hcl
ecs_desired_count = 2
db_instance_class = "db.t3.small"
skip_final_snapshot = false
multi_az = true
```
