
# GRPC-USERMANAGER

## Requirements
- Golang 1.20
- Docker

## Description
- This repository provides implementation of a GRPC based server-client model wherein client fetches user details from a server

## Installation
### Server
- `cd` into root directory of this project (i.e. parallel to README.md)
- Run following command to create docker image for grpc-server
```
docker build -t grpc-server .
```
- Run the create image as container
```
docker run -p 50051:50051 grpc-server
```
- If the container has loaded properly, following log will be printed
```
server listening at: [::]:50051
```
- Server has started by now

### Client
- The same can be performed to run client
- For testing the server, client can be run locally by issuing following command from the root dir of this project
```
go run client/main.go
```

## Endpoints
- Server provides 2 endpoints:
GetUser: accepts a struct with Id field
GetUsers: accepts a struct with an array of Ids

- Responses:
GetUser: responds with UserDetails struct or error
GetUsers: responds with array of UserDetails struct or error
