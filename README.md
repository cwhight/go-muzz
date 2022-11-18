# Basic Dating App

`go 1.18`

## Run the App

`go run cmd/main.go`

Application will start at `http://localhost:3000/`

## Endpoints

```
POST /user/create

Success Response - 201

{
        "id": <uuid>,
        "email": <string>,
        "password": <string>,
        "name": <string>,
        "gender": <string>,
        "age": <int>
}
``` 

```
POST /login

Request:

{
  password: <string>,
  email: <string>
}

Success Response - 200:

Set-Cookie: access-token=<jwt>
```




