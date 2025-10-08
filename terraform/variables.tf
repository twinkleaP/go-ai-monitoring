variable "aws_region" {
  description = "AWS region"
  default     = "us-east-1"
}

variable "project_name" {
  description = "Project name"
  default     = "go-ai-monitoring"
}

variable "go_image" {
  description = "Docker image for Go service"
  default     = "tk2802/go-service:latest"
}

variable "ai_image" {
  description = "Docker image for AI service"
  default     = "tk2802/ai-service:latest"
}
