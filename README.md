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
  createdb grocery_db

  ```

- Update the database connection string in `internal/config/config.go` if necessary.

1. Generate GraphQL code:
   go run [github.com/99designs/gqlgen](http://github.com/99designs/gqlgen) generate

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
  createUser(input: { name: "John Doe", email: "john@example.com" }) {
    id
    name
    email
  }
}
```

1. Query all users:

```graphql
query {
  users {
    id
    name
    email
  }
}
```

1. Query a specific user:

```graphql
query {
  user(id: "1") {
    id
    name
    email
  }
}
```

## Start the Entity

### Step 1: Create struct

### Go into `internal/model/product.go`

```
type Product struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

```

### Step 2: Create repository

### Go into `internal/repository/product.go`

```
type ProductRepository interface {
    CreateProduct(ctx context.Context, product *models.Product) error
    GetProductByID(ctx context.Context, id string) (*models.Product, error)
    UpdateProduct(ctx context.Context, product *models.Product) error
    DeleteProduct(ctx context.Context, id string) error
}

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *models.Product) error {
    // Implement create product logic
}

// Implement other repository methods

```

### Step 3: Create service

### Go into `internal/service/product.go`

```
type ProductService interface {
    CreateProduct(ctx context.Context, product *models.Product) error
    GetProductByID(ctx context.Context, id string) (*models.Product, error)
    UpdateProduct(ctx context.Context, product *models.Product) error
    DeleteProduct(ctx context.Context, id string) error
}

type productService struct {
    repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
    return &productService{repo: repo}
}

func (s *productService) CreateProduct(ctx context.Context, product *models.Product) error {
    // Implement create product logic
    return s.repo.CreateProduct(ctx, product)
}

// Implement other service methods

```

## GRAPHQL

### Step 4: Create resolver

### Go into `graphql/resolver/product.go`

```
type ProductResolver struct {
    service ProductService
}

func NewProductResolver(service ProductService) *ProductResolver {
    return &ProductResolver{service: service}
}

func (r *ProductResolver) CreateProduct(ctx context.Context, args struct{ Product *models.Product }) (*models.Product, error) {
    product, err := r.service.CreateProduct(ctx, args.Product)
    return product, err
}

// Implement other resolver methods

```

### Step 5: Create schema

### Go into `graph/schema.graphqls`

### After that run `go run github.com/99designs/gqlgen generate`

```
type Product {
    id: ID!
    name: String!
    description: String!
    price: Float!
    createdAt: Time!
    updatedAt: Time!
}

type Query {
    getProductByID(id: ID!): Product
}

type Mutation {
    createProduct(product: ProductInput!): Product!
    updateProduct(id: ID!,

```

### Step 6: Define input type

### Go into `graphql/schema.graphqls`

```
input ProductInput {
    name: String!
    description: String!
    price: Float!
}

```

### Step 7: Implement resolver

### Go into `graphql/resolver/product_resolver.go`

```
func (r *ProductResolver) GetProductByID(ctx context.Context, args struct{ ID string }) (*models.Product, error) {
    product, err := r.service.GetProductByID(ctx, args.ID)
    return product, err
}

func (r *ProductResolver) UpdateProduct(ctx context.Context, args struct{ ID string; Product *models.Product }) (*models.Product, error) {
    product, err := r.service.UpdateProduct(ctx, args.Product)
    return product, err
}

func (r *ProductResolver) DeleteProduct(ctx context.Context, args struct{ ID string }) (bool, error) {
    err := r.service.DeleteProduct(ctx, args.ID)
    return err == nil, err
}

```

### Step 8: Combine resolver into graphql

### Go into `cmd/server/main.go`

```
func main() {
    // ... other setup code

    productRepo := repository.NewProductRepository(db)
    productService := service.NewProductService(productRepo)
    productResolver := resolver.NewProductResolver(productService)

    schema := graphql.MustParseSchema(schema.Schema, &resolver.Resolver{
        ProductResolver: productResolver,
        // Other resolvers
    })

    // ... server setup code
}

```

Project Structure

- cmd/: Contains the main application entry points.
- internal/: Contains the internal packages of the application.
- graph/: Contains GraphQL-related files.
- migrations/: Contains database migration files (if any).
- scripts/: Contains utility scripts.

## 
```
.
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── db.go
│   ├── handler
│   │   └── graphql.go
│   ├── model
│   │   └── user.go
│   ├── repository
│   │   └── user_repository.go
│   └── service
│       └── user_service.go
├── pkg
│   └── utils
│       └── utils.go
├── graph
│   ├── generated
│   ├── model
│   ├── resolvers
│   └── schema.graphqls
├── migrations
│   └── 001_create_users_table.sql
├── scripts
│   └── seed.go
├── go.mod
├── go.sum
└── README.md
```