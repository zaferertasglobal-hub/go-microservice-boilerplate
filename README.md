# Go Microservice Boilerplate

## Project Documentation

### Features
- Lightweight and fast
- Highly scalable
- Built with best practices
- RESTful API
- Easy to configure
- Includes testing and CI/CD pipelines

### Project Structure
```
/
├── cmd/               # Main application entry point
├── config/            # Configuration files
├── pkg/              # Business logic and application logic
├── internal/          # Internal application code
├── api/               # API definitions and routes
├── migrations/        # Database migrations
└── Dockerfile         # Container configuration
```  

### Prerequisites
- Go (1.16 or newer)
- Docker (optional)
- Git

### Installation Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/zaferertasglobal-hub/go-microservice-boilerplate.git
   cd go-microservice-boilerplate
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running Instructions
- Run the application:
   ```bash
   go run main.go
   ```

### API Endpoints
- **GET /api/v1/health**: Check service health
- **POST /api/v1/resources**: Create a new resource
- **GET /api/v1/resources/{id}**: Get a resource by ID
- **PUT /api/v1/resources/{id}**: Update a resource by ID
- **DELETE /api/v1/resources/{id}**: Delete a resource by ID

### Examples
- Fetch all resources:
   ```bash
   curl -X GET http://localhost:8080/api/v1/resources
   ```
- Create a resource:
   ```bash
   curl -X POST -H 'Content-Type: application/json' -d '{"name":"Example"}' http://localhost:8080/api/v1/resources
   ```

### Configuration Details
- Configuration can be defined in `config/config.yml`.
- Environment variables can also be used for sensitive information.

### Database Schema
- **resources**:
  - `id`: Integer, primary key
  - `name`: String, unique
  - `created_at`: Timestamp
  - `updated_at`: Timestamp

### Development Guide
1. Follow coding standards and best practices.
2. Document your code.
3. Use meaningful commit messages.

### Testing
- To run tests:
   ```bash
   go test ./...
   ```

### Building
- To build the application:
   ```bash
   go build -o myapp
   ```

### Deployment
- Containerize the application using Docker:
   ```bash
   docker build -t myapp .
   docker run -p 8080:8080 myapp
   ```

### Troubleshooting
- Check logs for errors:
   ```bash
   go run main.go | grep ERROR
   ```

### Contributing Guidelines
- Fork the repository.
- Create a feature branch.
- Submit a pull request.

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Support Info
For support, please open an issue in the repository.

### Author Attribution
- Developed by Zafer Ertas.
