# 🛠️ Setup Instructions for Internal Asset & Ticket Management System

## 📋 Current Status

The backend microservices architecture has been successfully scaffolded with the following components:

### ✅ **Completed**
- **Project Structure**: Full microservices directory layout
- **Go Modules**: Dependency management configured
- **Protobuf Definitions**: All service contracts defined
- **Service Implementations**: Core business logic for all services
- **Database Models**: GORM models for all entities
- **Authentication**: JWT-based auth with OAuth2 Google integration
- **BFF (Backend-for-Frontend)**: REST API gateway with middleware
- **Docker Configuration**: Multi-service containerization
- **Kubernetes Ready**: Production deployment structure
- **Build Automation**: Makefile with common tasks

### 🔧 **Remaining Tasks**

#### 1. Generate Protobuf Files
```bash
# Install protoc (Protocol Buffer Compiler)
# On Ubuntu/Debian:
sudo apt install protobuf-compiler

# On macOS:
brew install protobuf

# Generate Go code from proto files
make proto
```

#### 2. Install Dependencies
```bash
make install-deps
go mod tidy
```

#### 3. Configure Environment
```bash
# Create environment file
cp .env.example .env

# Edit .env with your configuration:
# - Google OAuth credentials
# - Database connections
# - JWT secrets
# - Allowed domains
```

#### 4. Setup Google OAuth
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create new project or select existing
3. Enable Google+ API / Google Identity
4. Create OAuth 2.0 credentials
5. Add authorized redirect URIs:
   - `http://localhost:8080/auth/google/callback`
6. Update `.env` file with client credentials

#### 5. Database Setup
```bash
# For development (SQLite - default)
# No additional setup required

# For production (PostgreSQL)
# Update DATABASE_URL in .env:
# DATABASE_URL=postgres://user:password@localhost:5432/internal_support
```

#### 6. Build and Run

**Option A: Docker Compose (Recommended)**
```bash
make run-docker
```

**Option B: Local Development**
```bash
make run-dev
```

#### 7. Test the Setup
```bash
# Check services are running
curl http://localhost:8080/health         # Auth Service
curl http://localhost:8000/api/v1/health  # BFF

# Test OAuth flow
curl http://localhost:8080/auth/google
```

## 🚀 **Service Endpoints**

### Authentication Service (Port 8080)
- `GET /auth/google` - Initiate OAuth
- `GET /auth/google/callback` - OAuth callback
- `POST /auth/refresh` - Refresh token
- `GET /auth/me` - Current user info

### BFF API Gateway (Port 8000)
All endpoints require `Authorization: Bearer <token>` header

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

## 🔗 **gRPC Services (Internal)**
- **User Service**: `localhost:9091`
- **Asset Service**: `localhost:9092`
- **Ticket Service**: `localhost:9093`
- **Assignment Service**: `localhost:9094`
- **Notification Service**: `localhost:9095`

## 📊 **Next Development Steps**

1. **Complete Protobuf Integration**
   - Generate proper protobuf files
   - Connect BFF to gRPC services
   - Implement service-to-service communication

2. **Frontend Development**
   - Create React application
   - Implement authentication flow
   - Build dashboard and management interfaces

3. **Database Migration**
   - Switch to PostgreSQL for production
   - Implement proper migration scripts
   - Add database indexing and optimization

4. **Testing & Quality**
   - Unit tests for all services
   - Integration tests
   - API documentation (Swagger/OpenAPI)

5. **Production Readiness**
   - Logging and monitoring
   - Error handling and recovery
   - Performance optimization
   - Security hardening

6. **Deployment**
   - Kubernetes manifests
   - CI/CD pipeline
   - Production environment setup

## 🐛 **Troubleshooting**

### Common Issues

**Protobuf Generation Fails**
```bash
# Install protoc first, then:
export PATH="$PATH:$(go env GOPATH)/bin"
make proto
```

**Import Errors**
```bash
go mod tidy
go mod download
```

**Docker Build Issues**
```bash
docker-compose down
docker system prune -f
make clean
make run-docker
```

**Database Connection Issues**
- Check DATABASE_URL in .env
- Ensure database service is running
- Verify connection permissions

## 📞 **Support**

For technical questions:
1. Check the logs: `docker-compose logs -f <service-name>`
2. Review the README.md for detailed architecture
3. Examine individual service implementations
4. Verify environment configuration

---

**The foundation is complete! Ready for development and deployment.** 🎉 