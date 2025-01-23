# ChatApp

ChatApp is a simple client-server chat application built using Go and Vue.js with Wails. This application allows multiple users to connect to a chatroom and exchange messages in real-time.

## Features

- Real-time messaging
- User-friendly interface
- Light and dark themes
- Easy setup and configuration

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Node.js and npm
- Wails CLI

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/chatapp.git
   cd chatapp
   ```

2. Install Go dependencies:

   ```sh
   go mod tidy
   ```

3. Install frontend dependencies:

   ```sh
   cd frontend
   npm install
   ```

### Running the Server

The server should be run in a separate process. One person should act as the host and run the server. Others should change the IP in `global/chatroom.go` to match the host's IP.

1. Open `global/chatroom.go` and set the `HOST` and `PORT` constants:

   ```go
   const (
       HOST = "your_host_ip" // everyone should put the hosts IP
       PORT = "9090"
   )
   ```

2. (the host should:) Start the server:

   ```sh
   go run server/server.go
   ```

### Running the Client

1. Build the frontend:

   ```sh
   cd frontend
   npm run build
   ```

2. Start the client:

   ```sh
   wails dev
   ```

### Usage

1. Open the application in your browser.
2. Enter a username and connect to the chatroom.
3. Start sending messages!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Wails](https://wails.io/)
- [Vue.js](https://vuejs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [DaisyUI](https://daisyui.com/)

Enjoy chatting!
