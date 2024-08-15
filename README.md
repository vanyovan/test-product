# Product
Product service backend application

## Introduction

Product service is a simple product management system that allows users to perform operations like create, delete, view, and update product

## Prerequisites

- Go programming language (version 1.20.4 or later)
- SQLite3 database (version 3.42.0 or later)

## Installation

git clone https://github.com/vanyovan/test-product.git

## Getting Started

1. Database using sqlite3 in db file named **database.db**

2. Configure the database connection details in the config.json file.

3. Run the application:
   ```bash
   go run ./cmd/main.go
   ```
   
   or you can use makefile
   ```bash
   make run
   ```
   
## API Endpoints

The following API endpoints are available:

- `POST /api/v1/product`: Create product to database.
- `GET /api/v1/product`: Retrieves all products in database.
- `PATCH /api/v1/product/{id}`: Update product.
- `DELETE /api/v1/product/{id}`: Delete product.

## Database

The application uses an SQLite3 database named **database.db**. The database contains the following tables:

1. `mst_user`: Stores information about all products.

   Columns:
   - `id` (TEXT): Product ID.
   - `name` (TEXT): Product name.
   - `description` (TEXT): Product description.
   - `price` (FLOAT): Product price.
   - `variety` (TEXT): Product variety.
   - `rating` (FLOAT): Product rating.
   - `stock` (INT): Product stock.


## Testing

To run the unit tests, use the following command:

```bash
make test
```
