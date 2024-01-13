# Money Splitter App
App to split money between friends and family. Based on the [Splitwise](https://www.splitwise.com/) app.

--- 

## Table of Contents
- [Installation and Setup Instructions](#installation-and-setup-instructions)  
- [Running the app](#running-the-app)  
- [API Documentation](#api-documentation)  
- [Todo](#todo)
---

## Installation and Setup Instructions
- Ensure go is installed on your machine
- Next, clone this repository and `cd` into it
- Copy the contents of `.env.example` into a new file called `.env` and fill in the required values
### Pre-requisites
It is recommended to use docker to run the app. If you don't have docker installed, you can install it from [here](https://docs.docker.com/get-docker/)

#### Running on your machine
1. [Go](https://golang.org/dl/)
2. [Make](https://www.gnu.org/software/make/)
3. [PostgreSQL](https://www.postgresql.org/download/)

---

## Running the app
You can use docker to run the app or run it directly on your machine.
### Running with Docker
- Run [`./start.sh`](./start.sh) to start the server
- Ensure to run `chmod +x start.sh` if you get a permission error
- Ensure that the [`docker-compose`](https://docs.docker.com/compose/) plugin is in the same path mentioned in the `start.sh` file
### Running on your machine
- Run `make run` to start the server

---

## API Documentation
### Health Check
```http
GET /health_check
```
A simple API to test out whether the server is up and running.

```http
GET /health_check_auth

Headers:
Authorization: <token>
```
A simple API to test out whether you are authorized.



### Authentication
#### Register
```http
POST /auth/register

body: {
  name: string
  "email": string,
  "password": string
  }
}
```
Endpoint to register a new user. Ensures that the user with same email is not already registered and then creates a new user with a `Status Accepted` response.
```json
{
  "id": string
}
```
#### Login 
```http
POST /auth/login

body: {
  "email": string,
  "password": string
}
```
Endpoint to login a user. Returns a JWT token which can be used to authenticate the user for other APIs.
```json
{
  "token": string
}
```
#### Logout 
```http
POST /auth/logout
```
Endpoint to log user out. Deletes JWT token from redis db to invalidate token. 
```json
{
  "message": "ok"
}
```


## Todo 
- [x] Add jwt redis authentication
- [ ] Add refresh token functionality
- [ ] Endpoints to
  - [ ] Create a group
  - [ ] Add a user to a group
  - [ ] Manage and view expenses
  - [ ] View balances

