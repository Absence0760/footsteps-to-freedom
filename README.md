# Web Development Template

A lightweight, production-ready template for building modern web applications.

## Tech Stack

### Frontend
- **HTMX** - Dynamic HTML without heavy JavaScript
- **Alpine.js** - Minimal JavaScript framework for interactivity
- **Tailwind CSS** - Utility-first CSS framework

### Backend
- **Go** with Fiber framework - Fast, lightweight HTTP framework
- **PostgreSQL** - Optional database (easily removable)

### Infrastructure
- **Docker & Docker Compose** - Local development
- **Terraform** - AWS deployment automation
- **AWS Services**: ECS Fargate, RDS, ALB, VPC

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Go 1.21+ (for local development without Docker)
- Terraform 1.5+ (for AWS deployment)
- AWS CLI configured (for deployment)

### Local Development

1. **Clone and setup**:
```bash
cp .env.example .env
# Edit .env with your settings
```

2. **Run with Docker**:
```bash
docker-compose up
```

3. **Access the application**:
- Frontend: http://localhost:3000
- API: http://localhost:3000/api

### Without Database

If you don't need a database, simply remove the PostgreSQL service from `docker-compose.yml` and comment out database-related code in `backend/main.go`.

## Project Structure

```
.
├── backend/              # Go backend application
│   ├── main.go          # Application entry point
│   ├── handlers/        # HTTP handlers
│   ├── models/          # Data models
│   └── static/          # Frontend files served by backend
├── frontend/            # Frontend source files
│   ├── index.html       # Main HTML file
│   └── styles/          # Custom styles
├── terraform/           # Infrastructure as Code
│   ├── main.tf          # Main Terraform configuration
│   ├── variables.tf     # Input variables
│   └── outputs.tf       # Output values
├── docker-compose.yml   # Local development setup
├── Dockerfile           # Application container
└── .env.example         # Environment variables template
```

## Deployment to AWS

### Setup

1. **Configure AWS credentials**:
```bash
aws configure
```

2. **Initialize Terraform**:
```bash
cd terraform
terraform init
```

3. **Review and apply**:
```bash
terraform plan
terraform apply
```

### What Gets Deployed

- **VPC** with public/private subnets across 2 AZs
- **Application Load Balancer** for traffic distribution
- **ECS Fargate** cluster running your containerized app
- **RDS PostgreSQL** database (optional)
- **ECR** repository for Docker images
- **Security Groups** with proper access controls

### Custom Domain (Optional)

To use a custom domain:
1. Add your domain to Route53
2. Update `terraform/variables.tf` with your domain name
3. Uncomment the Route53 and ACM sections in `terraform/main.tf`

## Customization Guide

### For Each New Client/Project

1. **Clone this template**:
```bash
git clone <this-repo> client-project-name
cd client-project-name
```

2. **Update branding**:
   - Edit `frontend/index.html` with client branding
   - Update `backend/main.go` application name
   - Modify Tailwind colors in HTML

3. **Configure environment**:
```bash
cp .env.example .env
# Update with client-specific values
```

4. **Deploy**:
```bash
cd terraform
terraform init
terraform apply -var="project_name=client-name"
```

## Environment Variables

See `.env.example` for all available configuration options.

Key variables:
- `APP_PORT` - Application port (default: 3000)
- `DATABASE_URL` - PostgreSQL connection string
- `ENV` - Environment (development/production)

## Development Tips

### Adding New Routes

Edit `backend/main.go`:
```go
app.Get("/api/new-endpoint", handlers.NewHandler)
```

### Adding Database Models

Create new model in `backend/models/` and add migration logic in `main.go`.

### Frontend Updates

Edit `frontend/index.html` directly. The backend serves these static files in production.

## Production Considerations

- [ ] Set up proper secrets management (AWS Secrets Manager)
- [ ] Configure auto-scaling policies
- [ ] Set up monitoring and logging (CloudWatch)
- [ ] Enable HTTPS with ACM certificate
- [ ] Configure backup policies for RDS
- [ ] Set up CI/CD pipeline

## Cost Optimization

For small projects:
- Use `t3.micro` or `t4g.micro` instances
- Consider Aurora Serverless v2 for database
- Use CloudFront for static assets
- Enable RDS autoscaling

## Support

For issues or questions:
1. Check the README
2. Review Terraform outputs
3. Check Docker logs: `docker-compose logs`

## License

MIT License - customize for your company
