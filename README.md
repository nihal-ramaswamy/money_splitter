# Money Splitter App
App to split money between friends and family. Based on the [Splitwise](https://www.splitwise.com/) app.

--- 

## Table of Contents
- [Installation and Setup Instructions](#installation-and-setup-instructions)  
- [Running the app](#running-the-app)  

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
3. [Postgres](https://www.postgresql.org/download/)

---

## Running the app
You can use docker to run the app or run it directly on your machine.
### Running with Docker
- Run `./start.sh` to start the server
- Ensure to run `chmod +x start.sh` if you get a permission error
- Ensure that the [`docker-compose`](https://docs.docker.com/compose/) cli-plugin is in the same path mentioned in the `start.sh` file
### Running on your machine
- Run `make run` to start the server

---
