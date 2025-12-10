.PHONY: help dev build push deploy destroy logs clean

# Variables
PROJECT_NAME ?= web-template
AWS_REGION ?= us-east-1

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev: ## Start local development environment with hot reload
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build

dev-down: ## Stop local development environment
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml down

prod: ## Start production environment locally
	docker-compose up --build

prod-down: ## Stop production environment
	docker-compose down

dev-logs: ## View logs from local development
	docker-compose logs -f

build: ## Build Docker image locally
	docker build -t $(PROJECT_NAME):latest .

test-local: ## Test the application locally without Docker Compose
	cd backend && go run main.go

init-terraform: ## Initialize Terraform
	cd terraform && terraform init

plan: ## Run Terraform plan
	cd terraform && terraform plan -var="project_name=$(PROJECT_NAME)" -var="aws_region=$(AWS_REGION)"

deploy: ## Deploy to AWS (runs Terraform apply + Docker build/push + ECS update)
	./deploy.sh $(PROJECT_NAME) $(AWS_REGION)

deploy-infra-only: ## Deploy only infrastructure with Terraform
	cd terraform && terraform apply -var="project_name=$(PROJECT_NAME)" -var="aws_region=$(AWS_REGION)"

push: ## Build and push Docker image to ECR
	@echo "Getting ECR URL..."
	$(eval ECR_URL := $(shell cd terraform && terraform output -raw ecr_repository_url))
	@echo "Logging into ECR..."
	aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin $(ECR_URL)
	@echo "Building image..."
	docker build -t $(PROJECT_NAME):latest .
	@echo "Tagging image..."
	docker tag $(PROJECT_NAME):latest $(ECR_URL):latest
	@echo "Pushing image..."
	docker push $(ECR_URL):latest

update-service: ## Force ECS service to redeploy
	$(eval CLUSTER := $(shell cd terraform && terraform output -raw ecs_cluster_name))
	$(eval SERVICE := $(shell cd terraform && terraform output -raw ecs_service_name))
	aws ecs update-service --cluster $(CLUSTER) --service $(SERVICE) --force-new-deployment --region $(AWS_REGION)

logs: ## View CloudWatch logs
	aws logs tail /ecs/$(PROJECT_NAME) --follow --region $(AWS_REGION)

status: ## Check ECS service status
	$(eval CLUSTER := $(shell cd terraform && terraform output -raw ecs_cluster_name))
	$(eval SERVICE := $(shell cd terraform && terraform output -raw ecs_service_name))
	aws ecs describe-services --cluster $(CLUSTER) --services $(SERVICE) --region $(AWS_REGION)

url: ## Get the application URL
	@cd terraform && terraform output -raw alb_url

destroy: ## Destroy all AWS resources
	cd terraform && terraform destroy -var="project_name=$(PROJECT_NAME)" -var="aws_region=$(AWS_REGION)"

clean: ## Clean up local files and Docker resources
	docker-compose down -v
	rm -rf backend/tmp
	find . -name "*.log" -delete

setup-client: ## Setup for a new client project (usage: make setup-client CLIENT=client-name)
	@if [ -z "$(CLIENT)" ]; then \
		echo "Error: CLIENT variable not set. Usage: make setup-client CLIENT=client-name"; \
		exit 1; \
	fi
	@echo "Setting up project for client: $(CLIENT)"
	@sed -i '' 's/web-template/$(CLIENT)/g' terraform/variables.tf backend/main.go || \
	 sed -i 's/web-template/$(CLIENT)/g' terraform/variables.tf backend/main.go
	@echo "Updated project name to: $(CLIENT)"
	@echo "Don't forget to update:"
	@echo "  - frontend/index.html with client branding"
	@echo "  - .env file with client-specific values"

db-migrate: ## Run database migrations (when connected to local dev)
	@echo "Running migrations..."
	cd backend && go run main.go

fmt: ## Format Go code
	cd backend && go fmt ./...

tidy: ## Tidy Go modules
	cd backend && go mod tidy

download-assets: ## Download frontend assets (Tailwind, HTMX, Alpine.js) for local use
	./download-assets.sh

clean-assets: ## Remove downloaded frontend assets (forces CDN fallback)
	rm -rf frontend/assets/css frontend/assets/js
	@echo "Assets removed. Pages will now use CDN versions."
