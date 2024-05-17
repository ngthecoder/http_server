# HTTP Server
This project is a basic HTTP server implemented in Go, designed to handle various types of HTTP requests and respond appropriately. The server provides several key functionalities, including responding to basic HTTP requests, echoing messages, reading request headers, and supporting concurrent connections.

## Features
### TCP Server
The server listens on port 4221, utilizing the TCP protocol to establish and manage client connections. This forms the foundation for handling HTTP requests.

### Basic HTTP Response
The server responds to HTTP requests with a 200 OK status code. This response includes only the status line, adhering to the HTTP/1.1 specification.

### URL Path Handling
The server processes the URL path from incoming HTTP requests. Depending on the path, it responds with either a 200 OK status code for known paths or a 404 Not Found status code for unknown paths.

### Echo Endpoint
The /echo/{str} endpoint allows clients to send a string as part of the request. The server responds with a 200 OK status code, echoing the received string in the response body. The response also includes Content-Type and Content-Length headers to ensure proper parsing by the client.

### User-Agent Endpoint
The /user-agent endpoint reads the User-Agent header from the client's request and returns its value in the response body. This allows clients to verify that their user agent information is correctly received and processed by the server.

### Concurrent Connections
The server supports concurrent client connections, enabling it to handle multiple requests simultaneously. This improves the server's performance and responsiveness under load.

## Running the Server
To run the server, use the following command:
```
make run
```
The server will start listening on port 4221, ready to handle incoming HTTP requests as per the described functionalities.

## Conclusion
This Go-based HTTP server demonstrates fundamental concepts of HTTP request handling, including basic responses, path processing, echoing messages, reading headers, and supporting concurrent connections. It serves as a foundational example for building more complex and feature-rich HTTP servers.