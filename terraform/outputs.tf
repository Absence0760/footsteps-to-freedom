output "alb_dns_name" {
  description = "DNS name of the Application Load Balancer"
  value       = aws_lb.main.dns_name
}

output "alb_url" {
  description = "URL to access the application"
  value       = "http://${aws_lb.main.dns_name}"
}

output "ecr_repository_url" {
  description = "URL of the ECR repository"
  value       = aws_ecr_repository.main.repository_url
}

output "ecs_cluster_name" {
  description = "Name of the ECS cluster"
  value       = aws_ecs_cluster.main.name
}

output "ecs_service_name" {
  description = "Name of the ECS service"
  value       = aws_ecs_service.main.name
}

output "database_endpoint" {
  description = "Endpoint of the RDS database"
  value       = var.enable_database ? aws_db_instance.main[0].endpoint : "Database not enabled"
}

output "cloudwatch_log_group" {
  description = "Name of the CloudWatch log group"
  value       = aws_cloudwatch_log_group.main.name
}

output "deployment_instructions" {
  description = "Instructions for deploying the application"
  value = <<-EOT
  
  Next Steps:
  ===========
  
  1. Build and push your Docker image:
     
     aws ecr get-login-password --region ${var.aws_region} | docker login --username AWS --password-stdin ${aws_ecr_repository.main.repository_url}
     docker build -t ${var.project_name} .
     docker tag ${var.project_name}:latest ${aws_ecr_repository.main.repository_url}:latest
     docker push ${aws_ecr_repository.main.repository_url}:latest
  
  2. Update ECS service to use the new image:
     
     aws ecs update-service --cluster ${aws_ecs_cluster.main.name} --service ${aws_ecs_service.main.name} --force-new-deployment --region ${var.aws_region}
  
  3. Access your application:
     
     http://${aws_lb.main.dns_name}
  
  4. View logs:
     
     aws logs tail /ecs/${var.project_name} --follow --region ${var.aws_region}
  
  EOT
}
