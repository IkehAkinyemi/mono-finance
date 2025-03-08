# Mono Finance 🏦

A secure and robust banking API service built with Go, featuring user authentication, account management, money transfers, and more.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [Setup and Installation](#setup-and-installation)
  - [Prerequisites](#prerequisites)
  - [Local Development](#local-development)
  - [Using Docker](#using-docker)
  - [Database Migrations](#database-migrations)
- [Authentication](#authentication)
- [Deployment](#deployment)
- [CI/CD Pipeline](#cicd-pipeline)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

Mono Finance is a banking API service that provides a comprehensive solution for managing financial operations, including user authentication, account management, money transfers, and more. It's built with Go and follows modern backend development best practices.

## Architecture

The application follows a clean architecture pattern, separating concerns into different layers:

- **API Layer**: REST and gRPC interfaces
- **Service Layer**: Business logic implementation
- **Repository Layer**: Database interactions
- **Domain Layer**: Core business entities

The application can serve both REST (via Gin) and gRPC endpoints, with the gRPC gateway translating RESTful requests to gRPC calls.


## Features

- **User Management**:
  - User registration and authentication
  - JWT and PASETO token-based authentication
  - Email verification system

- **Account Management**:
  - Create and manage multiple accounts
  - Support for different currencies (USD, EUR, NGN, GBP)
  - Account balance tracking

- **Transaction Processing**:
  - Money transfers between accounts
  - Transaction history tracking
  - Atomic transactions with database consistency

- **API Support**:
  - RESTful API via Gin framework
  - gRPC API with protobuf
  - API documentation via Swagger/OpenAPI

- **Security**:
  - Password hashing
  - JWT/PASETO token authentication
  - Role-based access control

## Technologies Used

- **Backend**: Go (Golang)
- **Web Frameworks**:
  - Gin (HTTP REST)
  - gRPC (Protocol Buffers)
- **Database**:
  - PostgreSQL
  - SQL migrations with golang-migrate
- **Query Generation**: sqlc for type-safe SQL
- **Authentication**: JWT and PASETO tokens
- **API Documentation**: OpenAPI/Swagger
- **Asynchronous Processing**: Redis + Asynq
- **Containerization**: Docker
- **Deployment**:
  - AWS EKS (Elastic Kubernetes Service)
  - Amazon ECR
- **CI/CD**: GitHub Actions
- **Testing**: Go's testing package with testify

## API Endpoints

### REST API

#### User Management
- `POST /users`: Create a new user
- `POST /users/login`: Login a user
- `POST /tokens/renew_access`: Renew access token

#### Account Management
- `POST /accounts`: Create a new account
- `GET /accounts/:id`: Get account details
- `GET /accounts`: List all accounts
- `PUT /accounts/:id`: Update account details
- `DELETE /accounts/:id`: Delete an account

#### Transfers
- `POST /transfers`: Make a transfer between accounts

### gRPC API

The gRPC API provides the same functionality with Protocol Buffers:
- `CreateUser`: Create a new user
- `LoginUser`: Login a user
- `UpdateUser`: Update user details
- `VerifyEmail`: Verify user email

## Project Structure

```
.
├── api/                  # REST API handlers
├── db/                   
│   ├── migrations/       # Database migration files
│   ├── mock/             # Mock store for testing
│   ├── query/            # SQL queries
│   └── sqlc/             # Generated Go code from SQL
├── gapi/                 # gRPC API handlers
├── mail/                 # Email sending functionality
├── pb/                   # Protocol buffer generated files
├── proto/                # Protocol buffer definitions
├── token/                # JWT and PASETO token functionality
├── utils/                # Utility functions
├── validator/            # Input validation
├── worker/               # Async task processing
├── eks/                  # Kubernetes deployment files
├── app.env               # Environment configuration
├── Dockerfile            # Docker build instructions
├── docker-compose.yaml   # Docker compose configuration
├── Makefile              # Build automation
├── sqlc.yaml             # SQLC configuration
└── go.mod                # Go module definition
```

## Database Schema

The database consists of the following main tables:

- **users**: Store user information
- **accounts**: Store account information
- **entries**: Record all changes to account balances
- **transfers**: Record transfers between accounts
- **sessions**: Store user sessions
- **verify_emails**: Store email verification information

## Setup and Installation

### Prerequisites

- Go 1.20 or higher
- PostgreSQL 12 or higher
- Redis (for async processing)
- Docker and Docker Compose (optional)

### Local Development

1. Clone the repository:
   ```
   git clone https://github.com/IkehAkinyemi/mono-finance.git
   cd mono-finance
   ```

2. Set up the environment variables (copy from app.env.example if available):
   ```
   cp app.env.example app.env
   ```

3. Install dependencies:
   ```
   go mod download
   ```

4. Run PostgreSQL and Redis (you can use the provided Docker commands):
   ```
   make postgres
   make redis
   ```

5. Create the database:
   ```
   make createdb
   ```

6. Run database migrations:
   ```
   make migrateup
   ```

7. Start the server:
   ```
   make server
   ```

### Using Docker

```
docker-compose up
```

This will start the PostgreSQL database, Redis, and the API server.

### Database Migrations

To apply migrations:
```
make migrateup
```

To revert the last migration:
```
make migratedown1
```

To create a new migration:
```
make new_migration name=migration_name
```

## Authentication

The application uses two types of tokens:
- **Access Token**: Short-lived token used for API authentication
- **Refresh Token**: Long-lived token used to obtain new access tokens

To authenticate, include the access token in the Authorization header:
```
Authorization: Bearer <access_token>
```

## Deployment

The application is configured for deployment to AWS EKS:

1. Build and push the Docker image to Amazon ECR
2. Apply Kubernetes configurations in the `eks/` directory
3. Configure ingress and certificates

The GitHub Actions workflow in `.github/workflows/deploy.yml` automates this process.

## CI/CD Pipeline

The project includes GitHub Actions workflows for:
- Running tests on pull requests (`.github/workflows/test.yml`)
- Deploying to production on merge to main (`.github/workflows/deploy.yml`)

## Testing

The project includes comprehensive unit tests. To run tests:

```
make test
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the BSD 3-Clause License.