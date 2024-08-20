# Real-Time Messaging and AI Chatbot Service

## Project details
This project is a simple real-time messaging and AI chatbot backend application written in Go. The scope of the project
includes:
- a layered-architecture with clear separation of concerns from the data access layer to API endpoints
- a cookie-based JWT implementation for login authentication
- a Gorilla WebSocket for full-duplex communication over a TCP connection
- a handler to process message streaming from the Cohere API
- a Gin router to handle HTTP routes and middleware for the REST API
- a password encryption and verification utility to securely store passwords

## Database schema
![Database schema](assets/db_schema.png)

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