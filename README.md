# echo-swag

Sample of using echo swag for learning purpose

## Generating Swagger Documentation for Your Go Project

### 1. Install go-swag

Ensure that you have `go-swag` installed on your system. Run the following command:

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

### 2. Generate Swagger Documentation

Navigate to your project's root directory and run the following command to generate Swagger documentation:

Please adapt the paths, filenames, and port numbers according to your project's structure and configuration.
```bash
swag init -g ./cmd/main.go
```

This assumes that your main application file is located at `./cmd/main.go`. Adjust the path accordingly.

### 3. Run Your Server

Start your Go server. Make sure you have the necessary server setup in your code. For example:

```bash
go run ./cmd/main.go
```

Replace `./cmd/main.go` with the actual path to your main application file.

### 4. Open Swagger Documentation

Open your web browser and navigate to the Swagger documentation at:

```plaintext
http://localhost:{{port}}/swagger/index.html
```

Replace `port` with the actual port number your server is running on.

Now, you should be able to view your Swagger documentation and interact with it using the Swagger UI.

**Note:** Ensure that you have properly added Swagger comments to your Go code before running the `swag init` command. Customize the comments in your code to provide accurate information about your API.
```

Please adapt the paths, filenames, and port numbers according to your project's structure and configuration.
