# Messaging and Chatbot Service

![gophers_chatting](assets/gopher_chat.jpg)

## Project details
This project features a simple real-time messaging and AI chatbot backend service written in Go. The scope of the project
includes:
- a cookie-based JWT implementation for login authentication
- a Gorilla WebSocket for full-duplex communication over a TCP connection
- a handler to process message streaming from the Cohere API
- a Gin router to handle HTTP routes and middleware for the REST API
- a password encryption and verification utility to securely store passwords

## Tech stack
- Go
- PostgreSQL
- Gin
- Gorilla WebSocket
- Cohere API
- Docker

## References
- [psql tutorial](https://tomcam.github.io/postgres/)
- [Go Password Hashing](https://gowebexamples.com/password-hashing/)
- [Cohere API docs](https://docs.cohere.com/reference/about)
