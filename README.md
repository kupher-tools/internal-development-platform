---
# Internal Developer Platform (IDP)

A cloud-native Internal Developer Platform (IDP) built in Golang to provide self-service APIs for platform engineering workflows.

This platform aims to simplify and standardize developer operations such as:

* Container repository provisioning
* Source code repository provisioning
* Code quality project setup
* CI/CD integrations
* Kubernetes onboarding
* Platform automation

The project follows production-grade Golang architecture and cloud-native engineering practices.

---

# Architecture Goals

The platform is designed with the following goals:

* Production-grade API architecture
* Clean separation of layers
* RESTful API design
* Config-driven behavior
* Cloud-native compatibility
* Kubernetes readiness
* Extensible provider integrations
* Strong observability foundation
* Platform engineering best practices

---

# Current Features

## Health Endpoints

* Liveness endpoint
* Readiness endpoint

## Container Repository APIs

* Create container repositories
* Swagger/OpenAPI documentation
* Structured logging middleware
* Error handling middleware

---

# Technology Stack

* Golang
* Chi Router
* Swagger / OpenAPI
* Viper Configuration
* Docker
* Kubernetes

---

# Project Structure

```text
internal-development-platform/
├── cmd/
├── configs/
├── docs/
├── internal/
│   ├── config/
│   ├── domain/
│   ├── errors/
│   ├── handler/
│   ├── middleware/
│   ├── response/
│   ├── routes/
│   └── service/
├── go.mod
├── go.sum
└── README.md
```

---

# API Endpoints

## Health APIs

| Method | Endpoint        | Description     |
| ------ | --------------- | --------------- |
| GET    | `/health/live`  | Liveness probe  |
| GET    | `/health/ready` | Readiness probe |

---

## Container Repository APIs

| Method | Endpoint                         | Description                 |
| ------ | -------------------------------- | --------------------------- |
| POST   | `/api/v1/container-repositories` | Create container repository |

---

# Swagger API Documentation

Swagger UI is available at:

```text
http://localhost:8080/swagger/index.html
```

Generated Swagger files are available inside:

```text
docs/
```

---

# Running Locally

## Clone Repository

```bash
git clone https://github.com/kupher-tools/internal-development-platform.git
cd internal-development-platform
```

---

## Install Dependencies

```bash
go mod download
```

---

## Run Application

```bash
go run cmd/main.go
```

Application starts on:

```text
http://localhost:8080
```

---

# Generate Swagger Docs

Install Swagger CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate docs:

```bash
swag init -g cmd/main.go
```

---

# Configuration

The application uses:

* YAML configuration
* Environment variable overrides
* Typed configuration structs

Example config:

```yaml
server:
  port: "8080"

scm:
  baseurl: "https://gitlab.example.com"
```

---

# Logging

The application includes request logging middleware for:

* Request path
* HTTP method
* Status code
* Request duration
* Success/failure visibility

---

# Future Roadmap

Planned platform capabilities:

* Source repository provisioning
* SonarQube project creation
* CI/CD pipeline bootstrap
* Namespace provisioning
* GitOps integration
* RBAC automation
* OpenTelemetry integration
* Async operations API
* Multi-provider support
* Platform templates

---

# API Design Philosophy

The platform follows resource-oriented API design.

Example:

```text
/api/v1/container-repositories
/api/v1/source-repositories
/api/v1/code-quality-projects
```

The APIs abstract underlying platform tools and expose platform capabilities instead of vendor-specific operations.

---

# Kubernetes Readiness

The platform is designed for Kubernetes-native deployments and supports:

* Liveness probes
* Readiness probes
* Environment variable configuration
* Secret-based credentials
* Stateless service design

---

# Development Principles

* Clean architecture
* Dependency injection
* Structured error handling
* Middleware-driven observability
* Provider abstraction
* Cloud-native patterns
* Extensible platform design

---

# License

This project is intended for learning, experimentation, and platform engineering exploration.
