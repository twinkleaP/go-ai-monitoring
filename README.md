# go-ai-monitoring

Overview

Go microservice collects metrics :-
AI service (Python + FastAPI + IsolationForest)
+ Go service (metrics collector + inference client)

AI service (Python) detects anomalies.

Deployed on Kubernetes, provisioned with Terraform.

Monitored with Prometheus + Grafana.

Architecture - Go service → AI service → DB → Grafana.

Tech Stack

Backend: Go (Gin/Chi/GRPC)

AI: Python (FastAPI + scikit-learn IsolationForest)

Infra: Terraform, AWS (EKS, RDS)

CI/CD: GitHub Actions

Monitoring: Prometheus, Grafana
