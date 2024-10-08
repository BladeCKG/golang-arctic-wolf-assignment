# Risk Management API - Versioned with Gin

This is a simple Risk Management application built using the Go Gin framework. The application allows you to create, retrieve, and manage risks via HTTP endpoints. The application also implements API version control, with different versions of the same endpoints for demonstration.

## Features

- The application exposes three main operations:
  - **Create a new risk** (POST)
  - **Retrieve all risks** (GET)
  - **Retrieve a specific risk by ID** (GET)
- Data is stored in-memory (no external databases required).
- Response data is exchanged in JSON format.

## API Endpoints

### Version 1 (`/v1`)

- **GET /v1/risks**: Retrieves a list of all risks.
- **POST /v1/risks**: Creates a new risk with a `state`, `title`, and `description`.
- **GET /v1/risks/{id}**: Retrieves a specific risk by its unique ID (UUID).

## Data Structure

Each risk consists of the following fields:

- **ID**: A UUID automatically generated when the risk is created.
- **State**: A string value that can be one of the following: `open`, `closed`, `accepted`, `investigating`.
- **Title**: A short description of the risk.
- **Description**: A detailed explanation of the risk.

### Example Risk Object

```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "state": "open",
  "title": "Security Risk",
  "description": "Vulnerability detected in web application."
}
```

# Setup Instructions

## 1. Install Go

Ensure that you have Go installed on your system. You can download and install Go from here.

## 2. Clone the Repository

Clone the repository where this project resides.

```bash
git clone git@github.com:BladeCKG/golang-arctic-wolf-assignment.git
cd golang-arctic-wolf-assignment
```

## 3. Install Dependencies

This project uses the Gin web framework. To install the required dependencies, use the following command:

```bash
go mod tidy
```

This will install all dependencies specified in the go.mod file, including github.com/gin-gonic/gin.

## 4. Running the Application

You can run the application using the following command:

```bash
go run main.go
```

By default, the application will start on localhost:8080.

## 5. Example Usage (with curl)

Create a New Risk (POST /v1/risks)

```bash
curl -X POST http://localhost:8080/v1/risks \
 -d '{"state": "open", "title": "Risk 1", "description": "Description for Risk 1"}' \
 -H "Content-Type: application/json"
```

Retrieve All Risks (GET /v1/risks)

```bash
curl http://localhost:8080/v1/risks
```

Retrieve a Specific Risk by ID (GET /v1/risks/{id})

```bash
curl http://localhost:8080/v1/risks/{id}
```

Replace {id} with the actual UUID of a risk.

# Testing the Application

## 1. Running Unit Tests

The application includes unit tests to verify its functionality. These tests are located in the main_test.go file.

To run the tests:

```bash
go test
```

## 2. Test Cases

TestGetRisksEmpty: Verifies that an empty risk list returns an empty array.
TestCreateRisk: Tests the creation of a new risk via the POST /v1/risks endpoint.
TestGetRiskByID: Tests retrieving a specific risk by ID after it's been created.
TestGetRiskByIDNotFound: Tests the behavior when trying to retrieve a non-existent risk.
The tests simulate HTTP requests and check the expected responses without the need for a running server, using Go's httptest package.

Example Output

```bash
PASS
ok command-line-arguments 0.005s
```

# Notes

Data Persistence: The application uses in-memory storage for risks.

Versioning: The versioning in this application is designed to demonstrate how different API versions can coexist.

# License

This project is licensed under the MIT License. See the LICENSE file for details.
