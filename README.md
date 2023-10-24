# go-base-web

### Overview
**WIP**

`go-base-web` is a foundational Go web project aimed at providing a solid starting point for building scalable and maintainable web applications. It incorporates best practices and idiomatic Go code.

### Features
- Graceful shutdown
- Pre-configured logging using [zap](https://pkg.go.dev/go.uber.org/zap)
- HTTP request routing with [chi](https://pkg.go.dev/github.com/go-chi/chi/v5)
- Middleware support

### Prerequisites
- Go 1.18 or higher due to Generics, can be used with 1.16 if you don't plan on using Generics
- [zap](https://pkg.go.dev/go.uber.org/zap) for logging
- [chi](https://pkg.go.dev/github.com/go-chi/chi/v5) for routing

## Usage

1. Clone this repository:
    ```sh
        git clone https://github.com/yourusername/go-base-web.git
    ```
2. Navigate into the project directory:
    ```sh
        cd go-base-web
    ```
3. Build the project:
    ```sh
        go build -o go-base-web
    ```

## Usage
1. Start the server:
    ```sh
    ./mywebapp
    ```

2. Navigate to `http://localhost:8000` in your web browser to interact with the application.

## Configuration
The application listens on port 8000 by default. You can change this by editing the `addr` variable in `main.go`.

## Contributing
I'm relatively new to Go, so if you have any suggestion, please feel free to submit a pull request or create an issue.

Contributions are welcome!

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
