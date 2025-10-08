resource "aws_ecs_cluster" "main" {
  name = "${var.project_name}-cluster"
}

resource "aws_cloudwatch_log_group" "main" {
  name              = "/ecs/${var.project_name}"
  retention_in_days = 7
}

# Go Service Task Definition
resource "aws_ecs_task_definition" "go_service" {
  family                   = "go-service"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"

  container_definitions = file("${path.module}/go-service.json")
  execution_role_arn    = aws_iam_role.ecs_task_execution.arn
}

# AI Service Task Definition
resource "aws_ecs_task_definition" "ai_service" {
  family                   = "ai-service"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"

  container_definitions = file("${path.module}/ai-service.json")
  execution_role_arn    = aws_iam_role.ecs_task_execution.arn
}

# IAM Role for ECS tasks
resource "aws_iam_role" "ecs_task_execution" {
  name = "${var.project_name}-ecs-task-exec-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ecs-tasks.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_policy" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# AI Service
resource "aws_ecs_service" "ai_service" {
  name            = "ai-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.ai_service.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [aws_subnet.public.id]
    security_groups  = [aws_security_group.main.id]
    assign_public_ip = true
  }
}

# Go Service
resource "aws_ecs_service" "go_service" {
  name            = "go-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.go_service.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [aws_subnet.public.id]
    security_groups  = [aws_security_group.main.id]
    assign_public_ip = true
  }

  depends_on = [aws_ecs_service.ai_service]
}
