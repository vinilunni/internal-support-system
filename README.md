# 🏢 Internal Asset & Ticket Management System

A comprehensive microservices-based system for managing IT assets and support tickets within an organization. Built with Go backends, React frontend, and Kubernetes deployment.

## 🏗️ Architecture

### Microservices
- **Auth Service** (`port 8080/9090`): OAuth2 authentication, JWT token management
- **User Service** (`port 9091`): Employee profile management
- **Asset Service** (`port 9092`): IT asset inventory and tracking
- **Assignment Service** (`port 9094`): Asset assignment and return management
- **Ticket Service** (`port 9093`): Support ticket creation and management
- **Notification Service** (`port 9095`): Email and in-app notifications
- **BFF (Backend-for-Frontend)** (`port 8000`): REST API gateway

### Communication
- **External**: REST API (Frontend ↔ BFF)
- **Internal**: gRPC (Microservice ↔ Microservice)

### Technology Stack
- **Backend**: Go 1.21, gRPC, Protocol Buffers
- **Database**: PostgreSQL/SQLite (configurable)
- **Authentication**: OAuth2 (Google), JWT
- **Containerization**: Docker, Docker Compose
- **Orchestration**: Kubernetes (k3s/Minikube)

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- Protocol Buffers compiler (`protoc`)
- Make

### Development Setup

1. **Clone and setup**:
   ```bash
   git clone <repository-url>
   cd internal-support-system
   make dev-setup
   ```

2. **Configure environment**:
   ```bash
   cp .env.example .env
   # Edit .env with your Google OAuth credentials
   ```

3. **Run with Docker Compose**:
   ```bash
   make run-docker
   ```

4. **Or run in development mode**:
   ```bash
   make run-dev
   ```

### Available Commands

```bash
make help          # Show all available commands
make proto         # Generate protobuf files
make build         # Build all services
make test          # Run tests
make run-dev       # Run services locally
make run-docker    # Run with Docker Compose
make clean         # Clean build artifacts
```

## 📁 Project Structure

```
internal-support-system/
├── proto/                     # Protocol Buffer definitions
│   ├── auth/
│   ├── user/
│   ├── asset/
│   ├── assignment/
│   ├── ticket/
│   └── notification/
├── pkg/                       # Shared packages
│   ├── config/
│   └── jwt/
├── services/                  # Microservices
│   ├── auth-service/
│   ├── user-service/
│   ├── asset-service/
│   ├── assignment-service/
│   ├── ticket-service/
│   ├── notification-service/
│   └── bff/
├── k8s/                       # Kubernetes manifests
├── scripts/                   # Build and deployment scripts
├── docker-compose.yml
├── Makefile
└── README.md
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | Database connection string | `sqlite://app.db` |
| `JWT_SECRET` | JWT signing secret | `your-super-secret-key` |
| `GOOGLE_CLIENT_ID` | Google OAuth client ID | - |
| `GOOGLE_CLIENT_SECRET` | Google OAuth client secret | - |
| `ALLOWED_DOMAIN` | Allowed email domain | `yourcompany.com` |

### OAuth Setup

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select existing
3. Enable Google+ API
4. Create OAuth 2.0 credentials
5. Add authorized redirect URIs:
   - `http://localhost:8080/auth/google/callback`
6. Update environment variables with client credentials

## 📚 API Documentation

### Authentication Endpoints

- `GET /auth/google` - Initiate Google OAuth
- `GET /auth/google/callback` - OAuth callback
- `POST /auth/refresh` - Refresh access token
- `POST /auth/logout` - Logout user
- `GET /auth/me` - Get current user info

### REST API Endpoints (BFF)

All API endpoints require authentication via `Authorization: Bearer <token>` header.

#### Users
- `GET /api/v1/users` - List users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/{id}` - Get user
- `PUT /api/v1/users/{id}` - Update user
- `DELETE /api/v1/users/{id}` - Delete user

#### Assets
- `GET /api/v1/assets` - List assets
- `POST /api/v1/assets` - Create asset
- `GET /api/v1/assets/{id}` - Get asset
- `PUT /api/v1/assets/{id}` - Update asset
- `DELETE /api/v1/assets/{id}` - Delete asset
- `GET /api/v1/assets/user/{userId}` - Get user's assets

#### Tickets
- `GET /api/v1/tickets` - List tickets
- `POST /api/v1/tickets` - Create ticket
- `GET /api/v1/tickets/{id}` - Get ticket
- `PUT /api/v1/tickets/{id}` - Update ticket
- `POST /api/v1/tickets/{id}/close` - Close ticket
- `POST /api/v1/tickets/{id}/comments` - Add comment
- `GET /api/v1/tickets/{id}/comments` - Get comments

#### Assignments
- `GET /api/v1/assignments` - List assignments
- `POST /api/v1/assignments/assign` - Assign asset
- `POST /api/v1/assignments/{id}/unassign` - Unassign asset
- `GET /api/v1/assignments/user/{userId}` - Get user assignments

## 🐳 Docker Deployment

### Build Images
```bash
docker-compose build
```

### Run Services
```bash
docker-compose up -d
```

### View Logs
```bash
docker-compose logs -f <service-name>
```

## ☸️ Kubernetes Deployment

### Prerequisites
- Minikube or k3s cluster
- kubectl configured

### Deploy
```bash
kubectl apply -f k8s/
```

### Access Services
```bash
kubectl port-forward svc/bff 8000:8000
```

## 🔍 Monitoring & Observability

### Health Checks
- Auth Service: `http://localhost:8080/health`
- BFF: `http://localhost:8000/health`

### Metrics
- Prometheus metrics available at `/metrics` endpoint on each service

## 🧪 Testing

### Run Tests
```bash
make test
```

### Integration Tests
```bash
go test -tags=integration ./...
```

## 🛠️ Development

### Adding a New Service

1. Create service directory: `services/new-service/`
2. Define protobuf: `proto/new-service/new-service.proto`
3. Generate code: `make proto`
4. Implement service logic
5. Add to docker-compose.yml
6. Update BFF routing

### Database Migrations

Each service manages its own database schema using GORM auto-migration. For production, consider using proper migration tools.
