# Serendib Asia Service

A Go-based microservice built with Fiber framework for handling Serendib Asia business operations.

## Features

- RESTful API endpoints
- Swagger documentation
- Database integration with MySQL and PostgreSQL support
- Structured logging with Zap
- Configuration management with Viper
- Input validation
- Unit testing with coverage reporting
- Linting with golangci-lint

## Prerequisites

- Go 1.22.0 or higher
- MySQL or PostgreSQL database
- Make (for build automation)
- Swag (for Swagger documentation)

## Project Structure

```
.
├── app/            # Application-specific code
├── cmd/            # Main applications
├── internal/       # Private application code
├── pkg/            # Public library code
├── docs/           # Documentation
├── test/           # Test files
└── build/          # Build output
```

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/chazool/serendib_asia_service.git
   cd serendib_asia_service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the database:
   - Use the provided SQL files to set up your database:
     - `Schema.sql` - Database schema
     - `master_data.sql` - Master data
     - `sample_data.sql` - Sample data for testing

4. Build the service:
   ```bash
   make build
   ```

## Available Make Commands

- `make build` - Build the service (includes tests and swagger generation)
- `make test` - Run unit tests
- `make coverage` - Generate test coverage report
- `make swagger` - Generate Swagger documentation
- `make lint` - Run linter with auto-fix

## API Documentation

The API documentation is available in two formats:
1. Swagger UI - Accessible at `/swagger/*` when the service is running
2. Static HTML documentation - Generated in the `build` directory

## Testing

Run the test suite:
```bash
make test
```

Generate coverage report:
```bash
make coverage
```

## Linting

Run the linter:
```bash
make lint
```

## Configuration

The service uses Viper for configuration management. Configuration can be provided through:
- Environment variables
- Configuration files
- Command-line flags

## Dependencies

Major dependencies include:
- Fiber v2 - Web framework
- GORM - ORM library
- Zap - Logging
- Viper - Configuration
- Swagger - API documentation

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]
