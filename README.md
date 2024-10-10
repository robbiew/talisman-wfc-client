# Talisman BBS WFC Client

This client application connects to the Talisman BBS WFC Server, handles user authentication, and streams real-time updates from the server log.

## Features

- **Interactive Authentication**: Dynamically handles username and password prompts from the server.
- **Real-Time Log Streaming**: Once authenticated, the client receives continuous updates from the server in real time.
- **Configurable via Command-Line Flags**: The server's address is configurable via command-line flags.

## Requirements

- Go 1.16 or higher
- A running instance of the Talisman BBS WFC Server [GitHub link](https://github.com/robbiew/talisman-wfc-server)

## Compilation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/talisman-wfc-client.git
   cd talisman-wfc-client
   ```

2. Compile the project:
   ```bash
   go build -o talisman-wfc-client main.go
   ```

This will create an executable named `talisman-wfc-client`.

## Usage

Once compiled, you can run the client by passing the required flags.

### Required Flags

- `--server`: The address of the server you want to connect to (e.g., `localhost:8080`).

### Example Command

```bash
./talisman-wfc-client --server localhost:8080
```

This command connects to a server running on `localhost` on port `8080`.

### How It Works

1. When you run the client, it first connects to the server specified by the `--server` flag.
2. The server sends a prompt for the username, which the client displays. The user enters their username, and it is sent to the server.
3. Next, the server prompts for the password. The user enters the password, and it is sent to the server.
4. Once the server verifies the credentials and the required security level, the client will continuously receive and display log updates from the server.

### Sample Output

When the client is running, it will display something like:

```
Username: jdoe
Password: ****
Authentication successful!
Node 1: User: jdoe, Location: Main Menu
Node 2: User: waiting for caller, Location: -
...
```

If authentication fails, the client will display an error message returned by the server.

## License

This project is licensed under the MIT License.



[def]: http://