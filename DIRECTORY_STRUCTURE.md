# Directory Structure

```
web-template/
│
├── README.md                          # Main documentation
├── QUICKSTART.md                      # Fast setup guide for new clients
├── TROUBLESHOOTING.md                 # Common issues and solutions
├── ARCHITECTURE.md                    # System design and diagrams
├── CLIENT_PROPOSAL_TEMPLATE.md        # Business proposal template
├── DIRECTORY_STRUCTURE.md             # This file
│
├── .env.example                       # Environment variables template
├── .gitignore                         # Git ignore rules
├── Dockerfile                         # Container build instructions
├── docker-compose.yml                 # Local development environment
├── Makefile                           # Command shortcuts
├── deploy.sh                          # Automated deployment script
│
├── backend/                           # Go backend application
│   ├── main.go                        # Application entry point
│   ├── go.mod                         # Go dependencies
│   ├── go.sum                         # Go dependency checksums
│   │
│   ├── handlers/                      # HTTP request handlers
│   │   └── items.go                   # CRUD operations for items
│   │
│   └── models/                        # Database models
│       └── item.go                    # Item model definition
│
├── frontend/                          # Frontend files
│   └── index.html                     # Main HTML (HTMX + Alpine.js + Tailwind)
│
└── terraform/                         # Infrastructure as Code
    ├── main.tf                        # AWS resources definition
    ├── variables.tf                   # Input variables
    └── outputs.tf                     # Output values
```

## File Descriptions

### Root Level Documentation

| File | Purpose |
|------|---------|
| `README.md` | Complete documentation with setup, deployment, and usage instructions |
| `QUICKSTART.md` | Fast-track guide for setting up new client projects |
| `TROUBLESHOOTING.md` | Solutions for common issues during development and deployment |
| `ARCHITECTURE.md` | System architecture, data flow, and infrastructure diagrams |
| `CLIENT_PROPOSAL_TEMPLATE.md` | Professional proposal template for client presentations |

### Configuration Files

| File | Purpose |
|------|---------|
| `.env.example` | Template for environment variables (copy to `.env` and customize) |
| `.gitignore` | Specifies which files Git should ignore |
| `Dockerfile` | Instructions for building the Docker container |
| `docker-compose.yml` | Local development environment configuration |
| `Makefile` | Shortcuts for common commands (dev, deploy, logs, etc.) |
| `deploy.sh` | One-command deployment script for AWS |

### Backend (Go)

```
backend/
├── main.go              # Application entry point
│                        # - Initializes Fiber web framework
│                        # - Sets up middleware (CORS, logging, recovery)
│                        # - Configures database connection
│                        # - Defines API routes
│                        # - Serves static frontend files
│
├── go.mod               # Go module definition and dependencies
├── go.sum               # Dependency version lock file
│
├── handlers/            # HTTP request handlers
│   └── items.go         # CRUD operations:
│                        # - GetItems(): List all items
│                        # - GetItem(): Get single item by ID
│                        # - CreateItem(): Create new item
│                        # - UpdateItem(): Update existing item
│                        # - DeleteItem(): Delete item
│
└── models/              # Database models (GORM)
    └── item.go          # Item model with fields:
                         # - ID, CreatedAt, UpdatedAt, DeletedAt
                         # - Name, Description, Status
```

### Frontend

```
frontend/
└── index.html           # Single-page application
                         # - HTMX for dynamic updates
                         # - Alpine.js for interactivity
                         # - Tailwind CSS for styling
                         # - Examples of CRUD operations
                         # - Responsive design
```

### Terraform (Infrastructure)

```
terraform/
├── main.tf              # AWS infrastructure definition:
│                        # - VPC and networking (public/private subnets)
│                        # - Security groups
│                        # - Application Load Balancer
│                        # - ECS Fargate cluster and service
│                        # - ECR repository
│                        # - RDS PostgreSQL (optional)
│                        # - CloudWatch logs
│                        # - IAM roles
│
├── variables.tf         # Input variables:
│                        # - Project name
│                        # - AWS region
│                        # - ECS resources (CPU, memory)
│                        # - Database configuration
│                        # - Environment settings
│
└── outputs.tf           # Output values:
                         # - Application URL
                         # - ECR repository URL
                         # - Database endpoint
                         # - Deployment instructions
```

## How Files Work Together

### Local Development Flow

```
1. docker-compose.yml
   ↓
   Starts PostgreSQL and App containers
   ↓
2. Dockerfile
   ↓
   Builds Go application
   ↓
3. backend/main.go
   ↓
   Serves frontend/index.html
   ↓
4. User accesses http://localhost:3000
```

### AWS Deployment Flow

```
1. deploy.sh
   ↓
2. terraform/main.tf
   ↓
   Creates AWS infrastructure
   ↓
3. Dockerfile
   ↓
   Builds and pushes to ECR
   ↓
4. ECS pulls from ECR and runs container
   ↓
5. ALB routes traffic to ECS tasks
   ↓
6. User accesses application via ALB URL
```

### Request Flow

```
User Request
   ↓
ALB (Load Balancer)
   ↓
ECS Fargate Task
   ↓
backend/main.go
   ↓
   ├── Static files? → frontend/index.html
   │
   └── API call? → backend/handlers/*.go
                      ↓
                   backend/models/*.go
                      ↓
                   PostgreSQL (RDS)
```

## Adding New Features

### New API Endpoint

1. Create handler in `backend/handlers/`
2. Register route in `backend/main.go`
3. Update frontend to call new endpoint

### New Database Model

1. Create model in `backend/models/`
2. Add migration in `backend/main.go`
3. Create handler in `backend/handlers/`

### New Frontend Page

1. Update `frontend/index.html`
2. Add HTMX attributes for dynamic behavior
3. Use Alpine.js for client-side state

### New Infrastructure

1. Edit `terraform/main.tf`
2. Add variables in `terraform/variables.tf`
3. Add outputs in `terraform/outputs.tf`
4. Run `terraform apply`

## File Sizes (Approximate)

```
Total: ~50 KB (excluding dependencies)

Documentation:  ~25 KB
Backend Code:   ~10 KB
Frontend:       ~10 KB
Terraform:      ~15 KB
Config Files:   ~5 KB
```

## Generated/Ignored Files

These files are NOT in the repository but are generated:

```
.env                        # From .env.example (contains secrets)
backend/tmp/                # Temporary Go build files
terraform/.terraform/       # Terraform plugins
terraform/*.tfstate         # Terraform state (sensitive)
node_modules/               # If you add npm packages
```

## Customization Hotspots

Files you'll edit most often for new clients:

1. **frontend/index.html** - Branding, colors, text
2. **.env** - Client-specific configuration
3. **terraform/variables.tf** - Resource sizing
4. **backend/handlers/** - Business logic
5. **backend/models/** - Data structure

## Development Workflow

```
1. Edit code locally
   ↓
2. Test with: make dev
   ↓
3. Verify changes work
   ↓
4. Deploy with: make deploy
   ↓
5. Monitor with: make logs
```

## Minimal Setup (No Database)

To remove database support:

**Files to modify:**
1. `docker-compose.yml` - Remove postgres service
2. `terraform/variables.tf` - Set `enable_database = false`
3. `backend/main.go` - Comment out `initDatabase()`

**Files to delete** (optional):
1. `backend/models/` - No longer needed
2. Database-related handlers in `backend/handlers/`
