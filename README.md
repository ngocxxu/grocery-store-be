# Go GraphQL PostgreSQL Project

This project is a backend application using Go, GraphQL, and PostgreSQL with GORM as the ORM.

## Prerequisites

- Go (version 1.16 or later)
- PostgreSQL
- Git

## Installation

1. Clone the repository:
   git clone https://github.com/ngocxxu/grocery-store-svelte-be.git
   cd grocery-store-svelte-be

2. Install dependencies:
   go mod tidy

3. Set up the database:

- Create a new PostgreSQL database:
  ```
  createdb graphql_db
  ```
- Update the database connection string in `internal/config/config.go` if necessary.

4. Generate GraphQL code:
   go run github.com/99designs/gqlgen generate

## Configuration

- Database configuration can be found in `internal/config/config.go`.
- GraphQL schema is defined in `graph/schema.graphqls`.

## Running the Application

1. Start the server:
   go run cmd/server/main.go

2. The GraphQL playground will be available at `http://localhost:8080`.

## Database Migrations

This project uses GORM's AutoMigrate feature. The database schema will be automatically updated when you run the application.

## Seeding the Database

To seed the database with initial data:
go run scripts/seed.go

## GraphQL Queries and Mutations

Here are some example queries and mutations you can run in the GraphQL playground:

1. Create a new user:
   ```graphql
   mutation {
     createUser(input: {name: "John Doe", email: "john@example.com"}) {
       id
       name
       email
     }
   }
   Query all users:
   ```

query {
users {
id
name
email
}
}
Query a specific user:

query {
user(id: "1") {
id
name
email
}
}

Project Structure
cmd/: Contains the main application entry points.
internal/: Contains the internal packages of the application.
graph/: Contains GraphQL-related files.
migrations/: Contains database migration files (if any).
scripts/: Contains utility scripts.

graph/
  ├── generated/
  │   └── generated.go
  ├── model/
  │   └── models_gen.go
  ├── schema.resolvers.go
  └── schema.graphqls
  |-- resolver.go