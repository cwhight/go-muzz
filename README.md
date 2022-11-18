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

### Authenticated Endpoints

Require valid `access-token` cookie in request - unauthorised requests return `401`

```
GET /profiles

Success Response - 200:

{
        [
                {
                    "id": <integer>,
                    "name": <string>,
                    "gender": <string>,
                    "age": <int>
                }, 
                {
                    "id": <integer>,
                    "name": <string>,
                    "gender": <string>,
                    "age": <int>
                },
                {
                    …
        ]
}
```

```
POST /swipe

Request: 

{
        userId: <uuid>, //swiping user
        profileId: <uuid>, //swiped user
        preference <YES/NO>
}

Success Response  - 201:

{
        matched: <bool>
        matchId: <uuid> // only present if mtached = true
}
```
