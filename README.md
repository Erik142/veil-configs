# Veil Configs: Your Nebula Configuration Sidekick 🚀

Welcome to Veil Configs, the ultimate solution for dynamically serving Nebula VPN configurations! Tired of manually distributing and updating your Nebula client configurations? Veil Configs automates this tedious process, ensuring your network is always up-to-date and securely connected.

## ✨ Features

*   **Dynamic Configuration Delivery:** Serve Nebula configurations to your clients on demand via gRPC.
*   **Centralized Management:** Manage all your Nebula client configurations from a single, easy-to-use server.
*   **Go-Powered Performance:** Built with Go for blazing-fast performance and reliability.
*   **Cobra & Viper Integration:** Enjoy robust command-line interface (CLI) and flexible configuration management through YAML files, environment variables, and CLI flags.
*   **Logrus Logging:** Comprehensive and customizable logging with Logrus for easy debugging and monitoring.
*   **Unit Tested:** Thoroughly tested codebase to ensure stability and correctness.

## 🚀 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have the following installed:

*   [Go (1.21 or higher)](https://golang.org/doc/install)
*   [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/) (if you plan to regenerate protobuf files)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Erik142/veil-configs.git
    cd veil-configs
    ```
2.  **Download Go modules:**
    ```bash
    go mod tidy
    ```

### Running the Server

The Nebula Config Server serves configurations to clients.

```bash
# Run with default address (localhost:50051)
go run cmd/server/main.go

# Run with a custom address
go run cmd/server/main.go --address :8080

# You can also build and run the binary
go build -o cmd/server/server cmd/server/main.go
./cmd/server/server --address :50051
```

### Running the Client

The Nebula Config Client fetches configurations from the server.

```bash
# Fetch config for client1 and save to nebula_config_client1.yaml
go run cmd/client/main.go --client-id client1

# Fetch config for client2 and save to a custom file
go run cmd/client/main.go --client-id client2 --output-file my_client2_config.yaml

# You can also build and run the binary
go build -o cmd/client/client cmd/client/main.go
./cmd/client/client --client-id client1
```

## ⚙️ Configuration

Veil Configs uses [Viper](https://github.com/spf13/viper) for flexible configuration. You can configure both the server and client using:

*   **Command-line flags:** (e.g., `--address :8080`, `--client-id client1`)
*   **Environment variables:** (e.g., `CLIENT_SERVER_ADDRESS=localhost:50051`)
*   **Configuration files:** Create `.server.yaml` or `.client.yaml` in your home directory or specify a custom path using the `--config` flag.

### Example `.server.yaml`

```yaml
server:
  address: ":50051"
```

### Example `.client.yaml`

```yaml
client:
  server_address: "localhost:50051"
  client_id: "client1"
  output_file: "my_custom_config.yaml"
```

## 📂 Project Structure

```
.
├── .git/                   # Git version control
├── .github/                # GitHub Actions workflows
│   └── workflows/
│       └── ci.yml
├── .gitignore              # Specifies intentionally untracked files to ignore
├── go.mod                  # Go module dependencies
├── go.sum                  # Checksums for Go module dependencies
├── nebula_config_client1.yaml # Example Nebula client configuration
├── README.md               # Project README file
├── renovate.json           # Renovate bot configuration
├── cmd/
│   ├── client/             # Client application entry point and Cobra commands
│   │   ├── app/            # Client Cobra application logic
│   │   │   ├── app_test.go
│   │   │   └── app.go
│   │   └── main.go         # Client main entry point
│   └── server/             # Server application entry point and Cobra commands
│       ├── app/            # Server Cobra application logic
│       │   ├── app_test.go
│       │   └── app.go
│       └── main.go         # Server main entry point
├── internal/
│   ├── client/             # Internal client logic (gRPC communication, file saving)
│   │   └── client.go
│   └── server/             # Internal server logic (gRPC server implementation)
│       └── server.go
└── pkg/
    ├── config/             # Nebula configuration structs and in-memory store
    │   ├── config_test.go
    │   └── config.go
    └── proto/              # Protocol Buffer definitions
        ├── dummy.go                # Dummy file to ensure directory is recognized
        └── nebula_config.proto
```

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
