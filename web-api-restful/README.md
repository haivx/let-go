## Final Project

## 1. Structure

```
.
├── README.md
├── main.go        -> Main entry
├── controller     -> function handle API
├── services       -> function handle specific logic
├── util           -> common function
├── db             -> database
├── db/seed        -> initial database
├── db/migration   -> migration database
├── model          -> define entity
├── repo           -> handle logic database
├── router         -> define router
├── config         -> common config
├── dto            -> validate incoming requests
├── middleware     -> define middleware
├── script         -> setup script
└── tmp            -> air (hot reload for development)
├── go.mod
├── go.sum
```

## 2. Requirement

Docker, docker-compose, air

## 3. Start

1. Run command

```
  docker-compose up
```

2. Init DB: 

Run file init_database.sql in /seed/init_database.sql

3. Login

```
  POST localhost:8080/api/v1/auth/login HTTP/1.1
  content-type: application/json

  {
      "password": "123a@123",
      "email": "liam_vo@gmail.com"
  }
```