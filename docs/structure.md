# Structure for project

```
├── cmd/                 # Service entry points
├── internal/            # Private application code
├── pkg/                 # Shared libraries
│   ├── appLogger/       # Logging utilities
│   ├── common/          # Common utilities
│   ├── constants/       # Constants definitions
│   ├── database/        # Database utilities
│   │   ├── postgres/    # PostgreSQL utilities
│   │   └── redis/       # Redis utilities
│   ├── dtos/            # Data Transfer Objects
│   ├── grpc/            # gRPC utilities
│   ├── queues/          # Queue utilities
│   ├── response/        # Response utilities
│   ├── settings/        # Configuration settings
│   ├── utils/           # General utilities
│   └── websocket/       # WebSocket utilities
├── configs/             # Configuration files
│   ├── development/     # Development environment configs
│   ├── production/      # Production environment configs
│   └── kubernetes/      # Kubernetes deployment configs
│       ├── deployments/
│       └── services/
├── docs/                # Documentation
│   ├── api/             # API documentation
│   ├── architecture/    # Architecture diagrams and documents
│   │   └── diagrams/
│   ├── devguide.md      # Development guidelines
│   ├── keyfeature.md    # Key features
│   ├── plan.md          # Implementation plan
│   └── structure.md     # Project structure
├── sqlc/                # SQLC generated code and SQL queries
│   ├── queries/         # SQL query files
│   │   ├── user/
│   │   ├── market-data/
│   │   ├── whale-tracking/
│   │   ├── portfolio/
│   │   └── notification/
│   └── generated/       # Generated Go code
│       ├── user/
│       ├── market-data/
│       ├── whale-tracking/
│       ├── portfolio/
│       └── notification/
├── tests/               # Test files
│   ├── unit/            # Unit tests
│   └── integration/     # Integration tests
├── scripts/             # Scripts for database migrations and seeding
├── sql/                 # SQL files
│   ├── queries/         # SQL query files
│   └── schema/          # Database schema
├── templates/           # Templates for services
│   ├── domain/
│   ├── gateway/
│   │   ├── .dockerignore
│   │   ├── infrastructure/
│   │   │   └── global/
│   │   └── proto/
│   │       └── pb/
├── Dockerfile           # Dockerfile for building the application
├── docker-compose.yaml  # Docker Compose file for local development
├── Makefile             # Makefile for build and deployment tasks
├── go.mod               # Go module file
├── go.sum               # Go dependencies file
└── README.md            # Project README file
```
