# finance-project

# Finance Project Challenge

## Overview

This project consists of two APIs that work together to process matrix data.

### Go API (Input API)

The Go API receives a matrix as input, validates it, performs a QR decomposition using Gonum, and sends the resulting matrices to the Node.js API.

### Node.js API (Output API)

The Node.js API receives the QR decomposition matrices, calculates statistical information, and returns the results.

---

## Architecture

```text
Client
  |
  v
Go API (Fiber)
  |
  | QR Decomposition
  v
Node.js API (Express)
  |
  | Statistics Calculation
  v
Response
```

---

## Technologies Used

### Go API

* Go 1.24
* Fiber
* Gonum

### Node.js API

* Node.js
* Express
* Dotenv

### Infrastructure

* Docker
* Docker Compose

---

## Project Structure

```text
finance-project/
│
├── input-api/
│   ├── handlers/
│   ├── models/
│   ├── services/
│   ├── Dockerfile
│   └── main.go
│
├── output-api/
│   ├── controllers/
│   ├── routes/
│   ├── services/
│   ├── Dockerfile
│   └── server.js
│
└── docker-compose.yml
```

---

## Features

### Go API

* Receives matrix data
* Validates matrix structure
* Performs QR decomposition
* Sends results to Node.js API

### Node.js API

* Receives Q and R matrices
* Calculates:

  * Maximum value
  * Minimum value
  * Average value
  * Sum of all values
  * Diagonal matrix validation

---

## Running the Project with Docker

From the project root directory:

```bash
docker compose up --build
```

The services will be available at:

* Go API: http://localhost:3000
* Node API: http://localhost:4000

---

## API Endpoints

### Go API

#### POST /received-matrix

Request:

```json
{
  "matrix": [
    [1, 2],
    [3, 4]
  ]
}
```

Response:

```json
{
  "incomingMatrix": [
    [0.1, 0.5],
    [0.2, 0.8]
  ],
  "qrDecomposition": {
    "q":[  [
                -0.4472135954999581,
                -0.8944271909999159
            ],
            [
                -0.8944271909999159,
                0.447213595499958
            ]
        ],
    "r":[
            [
                -0.223606797749979,
                -0.9391485505499118
            ],
            [
                0,
                -0.08944271909999146
            ]
        ]
  },
  "statistics": {
    "average": -0.3801315561749643,
    "hasDiagonalMatrix": false,
    "max": 0.447213595499958,
    "min": -0.9391485505499118,
    "sum": -3.0410524493997144
  }
}
```

---

## Validation Rules

* Matrix cannot be empty.
* Matrix must be rectangular.
* Invalid request bodies return HTTP 400.

---

## QR Decomposition

The Go API uses the Gonum library to perform QR decomposition.

A matrix A is decomposed into:

```text
A = Q × R
```

Where:

* Q is an orthogonal matrix.
* R is an upper triangular matrix.

---


The challenge description mentions matrix rotation in the overview and QR decomposition in the functional requirements. QR decomposition was implemented because it is the most specific requirement and is explicitly defined as part of the expected functionality.

Gonum was selected to perform the QR decomposition because it provides a reliable and production-ready implementation of numerical linear algebra algorithms.


