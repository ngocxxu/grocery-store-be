# Go GraphQL PostgreSQL Project

This project is a backend application using Go, GraphQL, and PostgreSQL.

## Setup

1. Install dependencies:
   go mod tidy

2. Set up your PostgreSQL database and update the `DATABASE_URL` in `internal/config/config.go`.

3. Run database migrations:
   psql -U your_username -d your_database_name -a -f migrations/001_create_users_table.sql

4. (Optional) Seed the database:
   go run scripts/seed.go

5. Run the server:
   go run cmd/server/main.go

6. Open http://localhost:8080 in your browser to access the GraphQL playground.

## Usage

Example queries and mutations:

1. Create a new user:

```graphql
mutation {
  createUser(name: "Alice", email: "alice@example.com") {
    id
    name
    email
  }  }


2. Query all users:

query {
  users {
    id
    name
    email
  }
}

3. Query a specific user:

query {
  user(id: "1") {
    id
    name
    email
  }
}
```
## Database

This project uses GORM as an ORM. The database schema is automatically managed by GORM's AutoMigrate feature.

To manually create the database:

```sql
CREATE DATABASE graphql_db;
To run the application:

Ensure your PostgreSQL server is running.
Set the DATABASE_URL environment variable or update it in internal/config/config.go.
Run the application:

go run cmd/server/main.go
This will automatically create the necessary tables in your database.

To seed the database with initial data:


go run scripts/seed.go