output "ecs_cluster_name" {
  value = aws_ecs_cluster.main.name
}

output "go_service_name" {
  value = aws_ecs_service.go_service.name
}

output "ai_service_name" {
  value = aws_ecs_service.ai_service.name
}
