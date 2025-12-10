variable "project_name" {
  description = "Name of the project (used for resource naming)"
  type        = string
  default     = "web-template"
}

variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment (development, staging, production)"
  type        = string
  default     = "production"
}

variable "enable_database" {
  description = "Whether to create an RDS database"
  type        = bool
  default     = true
}

# ECS Configuration
variable "ecs_cpu" {
  description = "CPU units for ECS task (256, 512, 1024, 2048, 4096)"
  type        = string
  default     = "256"
}

variable "ecs_memory" {
  description = "Memory for ECS task in MB (512, 1024, 2048, 4096, 8192)"
  type        = string
  default     = "512"
}

variable "ecs_desired_count" {
  description = "Desired number of ECS tasks"
  type        = number
  default     = 1
}

# Database Configuration
variable "db_instance_class" {
  description = "RDS instance class"
  type        = string
  default     = "db.t3.micro"
}

variable "db_name" {
  description = "Name of the database"
  type        = string
  default     = "appdb"
}

variable "db_username" {
  description = "Database master username"
  type        = string
  default     = "dbadmin"
}

variable "db_password" {
  description = "Database master password"
  type        = string
  sensitive   = true
  default     = "ChangeThisPassword123!"
}

# Optional: Domain Configuration
# Uncomment and configure if you want to use a custom domain
# variable "domain_name" {
#   description = "Domain name for the application"
#   type        = string
#   default     = ""
# }
