# Merge lists API

A simple API dedicated to merging two ordered lists.
## Overview

This API provides functionality to merge two ordered lists and retrieve the merged result. It allows saving two lists via the **SaveLists** endpoint and subsequently merging the saved lists using the **Merge** endpoint.
## Installation

To run this API, ensure you have Go installed. Then, clone the repository and install the necessary dependencies:

```bash

git clone https://github.com/angelicalombas/merge-list-api.git
cd merge-list-api
go mod download
```

## Usage
### Running the server

Run the server using the following command:

```bash
go run main.go
```
By default, the server will start at **localhost:8080**.

## Swagger Documentation

The API is documented using Swagger. Once the server is running, access the Swagger UI by visiting http://localhost:8080/swagger/index.html.

The Swagger UI provides an interactive interface to explore and test the API endpoints.

## Endpoints
#### SaveLists

   - **Description**: Save two lists provided in the request body.
   - **HTTP Method**: POST
   - **Endpoint**: /SaveLists
   - **Request Body**:

```json
{
  "list1": [1, 2, 4],
  "list2": [1, 3, 4]
}
```
   - **Success Response**:
       - **Code**: 200
       - **Content**:
```json
{
    "message": "Lists saved successfully."
}
```

#### Merge

   - **Description**: Merge previously saved lists.
   - **HTTP Method**: GET
   - **Endpoint**: /Merge
   - **Success Response**:
       - **Code**: 200
       - **Content**: 
```json
[1,1,2,3,4,4]
```

## Running Unit Tests

To run unit tests for this project, execute the following command:

```bash
go test ./...
```
## Dependencies

This project uses several external dependencies, including:

   - **github.com/gin-gonic/gin**
   - **github.com/swaggo/files**
   - **github.com/swaggo/gin-swagger**

## Code Structure

The codebase is structured into different packages:

```
.
├── controllers                         # implements the logic for API endpoints
├── docs                                # auto-generated docs on the usage of custom provider
├── models                              # contains data structures and functions for list manipulation
├── routes                              # handles HTTP requests and defines API endpoints
├── deployment.yaml                     # configuration file for Kubernetes
├── Dockerfile.yaml                     # configuration file for Docker
├── go.mod                              # project dependencies
└── main.go                             # Entry point for local development
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://choosealicense.com/licenses/mit/) file for details.