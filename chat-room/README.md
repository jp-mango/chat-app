# TCP Chat Application

This repository contains a simple TCP chat application written in Go. The project is structured into two main components: a server and a client.

## Current Implementation

- **Server (`/server/main.go`)**:

  - Listens for incoming TCP connections on port 8080.
  - Accepts connections and handles them concurrently.
  - Broadcasts messages from any client to all connected clients.

- **Client (`/client/main.go`)**:
  - Connects to the server on `localhost:8080`.
  - Sends messages to the server and receives broadcast messages.

## File Structure

```shell
/chat-room
│
├── server
│   ├── tcp_server.go
│   ├── go.mod
│   └── tcp_server.exe
│
├── client
│   ├── tcp-client.go
│   ├── go.mod
│   └── tcp-client.exe
└── README.md
```

## Planned Improvements

- **Robust Error Handling**: Enhance error handling on both server and client sides.
- **User Authentication**: Implement a user login system for personalized experience.
- **Message Encryption**: Secure the message transmission using encryption.
- **User Interface**: Develop a more user-friendly interface for the client.
- **Scalability**: Improve server architecture to handle more simultaneous connections efficiently.
- **Private Messaging**: Enable users to send private messages to each other in addition to broadcasting to all users.
- **Graceful Shutdown**: Handle server shutdowns gracefully, ensuring all connections are closed properly.
- **Command Parsing**: Add special commands users can type for actions like changing rooms, listing users, etc.
- **Notifications and Alerts**: Add user notifications for events like new messages, user join/leave, etc.

## Running the Application

- Start the server: Navigate to `/tcp-chat-app/server` and run `./tcp_server`.
- Start the client: In a new terminal, navigate to `/tcp-chat-app/client` and run `./tcp_client`.
- You can open multiple instances of the client to simulate a chat room.

---

Enjoy exploring and expanding this basic TCP chat application!
