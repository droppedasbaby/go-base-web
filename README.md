# go-base-web

### Overview
**WIP**

`go-base-web` is a foundational Go web project aimed at providing a solid starting point for building scalable and maintainable web applications. It incorporates best practices and idiomatic Go code.

### Features
- Graceful shutdown
- API documentation support using [http-swagger](https://github.com/swaggo/http-swagger)
- Pre-configured logging using [zap](https://github.com/uber-go/zap)
- HTTP request routing with [chi](https://github.com/go-chi/chi)
- Middleware support using chi
    - Includes logging
    - Include recovery
    - Adding Request ID
    - Adding Request IP

### Prerequisites
- Go 1.18 or higher due to Generics, can be used with 1.16 if you don't plan on using Generics
- [zap](https://github.com/uber-go/zap) for logging
- [chi](https://github.com/go-chi/chi) for routing
- [http-swagger](https://github.com/swaggo/http-swagger) for documentation

## Usage

1. Clone this repository:
    ```sh
        git clone https://github.com/yourusername/go-base-web.git
    ```
2. Navigate into the project directory:
    ```sh
        cd go-base-web
    ```
3. Place your api routes inside:
    ```sh
        api/routes.go
    ```
    Some example routes are already provided.
4. Place your operation handlers inside:
    ```sh
        api/*operations.go
    ```
    Some example handlers are already provided in: `api/testoperations.go`

## Swagger
Swagger documentation is automatically generated and served at `/swagger/index.html` by default.

You can change the path by editing the routes in `route/buildTopLevelRoutes.go`:
```go
    // ...
    r.Get("/docs/*", func(w http.ResponseWriter, r *http.Request) {
        http.StripPrefix("/docs", http.FileServer(http.Dir("./docs"))).ServeHTTP(w, r)
    })
    r.Get("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("/docs/swagger.json"),
    ))
    // ...
```

## Configuration
The application listens on port 8000 by default. You can change this by editing the `addr` variable in `orchestrateserver.go` OR by setting the `PORT` environment variable.

## Contributing
I'm relatively new to Go, so if you have any suggestion, please feel free to submit a pull request or create an issue.

Contributions are welcome!

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
