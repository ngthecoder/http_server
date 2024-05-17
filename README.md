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
The /echo/{string} endpoint allows clients to send a string as part of the request. The server responds with a 200 OK status code, echoing the received string in the response body. The response also includes Content-Type and Content-Length headers to ensure proper parsing by the client.

### User-Agent Endpoint
The /user-agent endpoint reads the User-Agent header from the client's request and returns its value in the response body. This allows clients to verify that their user agent information is correctly received and processed by the server.

### Concurrent Connections
The server supports concurrent client connections, enabling it to handle multiple requests simultaneously. This improves the server's performance and responsiveness under load.

### Files Endpoint
The /files/{filename} endpoint finds and returns up to 1014 bytes of the contents of the file in the responce body from the direcotry provided with --directory flag.

The serveFile function handles GET requests to serve files from the specified directory. If the file is found, it returns the file contents with a 200 OK status. If the file is not found, it returns a 404 Not Found status. If there is an internal server error, it returns a 500 Internal Server Error status. The function reads and returns up to 1024 bytes of the file content.

The saveFile function handles POST requests to save the request body as a file in the specified directory. If the file is successfully saved, it returns a 200 OK status with a message indicating the file was saved. If there is an error creating or writing to the file, it returns a 500 Internal Server Error status.

### Gzip Compression for Echo Endpoint
The server supports gzip compression for the /echo/{string} endpoint. When a client sends a GET request to this endpoint with the Accept-Encoding: gzip header, the server responds with the requested string gzip-encoded. This feature enhances data transmission efficiency especially when the request body is learge by reducing the size of the response, which is particularly beneficial for clients that support gzip encoding.

## Running the Server
To run the server, use the following command:
```
go build -o ./build/http_server ./app && ./build/http_server --directory <directory path>
```
The server will start listening on port 4221, ready to handle incoming HTTP requests as per the described functionalities.

## Conclusion
This Go-based HTTP server demonstrates fundamental concepts of HTTP request handling, including basic responses, path processing, echoing messages, reading headers, and supporting concurrent connections. It serves as a foundational example for building more complex and feature-rich HTTP servers.