# Go TCP Chat Server

This is a simple TCP Chat Server built using Go. It allows multiple users to connect to the server to write and read each other.

## Features

- In real time text chatting over a TCP connection to the server.

## Installation

### Prerequisites

- Have Go installed in your PC
- Have a tool like telnet for connecting over TCP to the server

### Steps

1. Clone this repository:

   ```bash
   git clone https://github.com/NoOPeEKS/go-tcp-server.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-tcp-server
   ```

3. Build the application:

   ```bash
   go build -o server
   ```

4. Run the application:

   ```bash
   ./server
   ```

## Usage

- To connect to the server from a different terminal:

  ```bash
    telnet 127.0.0.1 6969
  ```

- Proceed to start chatting.

## Contributing

Contributions are welcome! If you'd like to improve this project, feel free to fork the repository and submit a pull request.

## Future improvements

Future lines of improvements are:

- Implement a client application instead of using telnet.
- Unit tests to correct any malfunctions of the application.
- Implement a banning method to ban spammers.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit/).
