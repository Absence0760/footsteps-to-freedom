#!/bin/bash

# Deployment script for AWS
# This script builds, pushes, and deploys your application to AWS ECS

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=====================================${NC}"
echo -e "${GREEN}Web Template Deployment Script${NC}"
echo -e "${GREEN}=====================================${NC}"
echo ""

# Check if AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo -e "${RED}Error: AWS CLI is not installed${NC}"
    exit 1
fi

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker is not installed${NC}"
    exit 1
fi

# Check if Terraform is installed
if ! command -v terraform &> /dev/null; then
    echo -e "${RED}Error: Terraform is not installed${NC}"
    exit 1
fi

# Get project name
PROJECT_NAME=${1:-web-template}
AWS_REGION=${2:-us-east-1}

echo -e "${YELLOW}Project Name: ${PROJECT_NAME}${NC}"
echo -e "${YELLOW}AWS Region: ${AWS_REGION}${NC}"
echo ""

# Step 1: Initialize Terraform if not already done
echo -e "${GREEN}Step 1: Initializing Terraform...${NC}"
cd terraform
if [ ! -d ".terraform" ]; then
    terraform init
fi

# Step 2: Apply Terraform configuration
echo -e "${GREEN}Step 2: Deploying infrastructure...${NC}"
terraform apply -var="project_name=${PROJECT_NAME}" -var="aws_region=${AWS_REGION}" -auto-approve

# Get outputs
ECR_URL=$(terraform output -raw ecr_repository_url)
CLUSTER_NAME=$(terraform output -raw ecs_cluster_name)
SERVICE_NAME=$(terraform output -raw ecs_service_name)

cd ..

# Step 3: Build Docker image
echo -e "${GREEN}Step 3: Building Docker image...${NC}"
docker build -t ${PROJECT_NAME}:latest .

# Step 4: Login to ECR
echo -e "${GREEN}Step 4: Logging into AWS ECR...${NC}"
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${ECR_URL}

# Step 5: Tag and push image
echo -e "${GREEN}Step 5: Pushing image to ECR...${NC}"
docker tag ${PROJECT_NAME}:latest ${ECR_URL}:latest
docker push ${ECR_URL}:latest

# Step 6: Update ECS service
echo -e "${GREEN}Step 6: Updating ECS service...${NC}"
aws ecs update-service \
    --cluster ${CLUSTER_NAME} \
    --service ${SERVICE_NAME} \
    --force-new-deployment \
    --region ${AWS_REGION}

echo ""
echo -e "${GREEN}=====================================${NC}"
echo -e "${GREEN}Deployment Complete!${NC}"
echo -e "${GREEN}=====================================${NC}"
echo ""
echo -e "${YELLOW}Your application is being deployed...${NC}"
echo -e "${YELLOW}It may take a few minutes for the service to become healthy.${NC}"
echo ""
echo -e "${GREEN}View deployment status:${NC}"
echo "aws ecs describe-services --cluster ${CLUSTER_NAME} --services ${SERVICE_NAME} --region ${AWS_REGION}"
echo ""
echo -e "${GREEN}View logs:${NC}"
echo "aws logs tail /ecs/${PROJECT_NAME} --follow --region ${AWS_REGION}"
echo ""
echo -e "${GREEN}Access your application at:${NC}"
terraform -chdir=terraform output -raw alb_url
echo ""
