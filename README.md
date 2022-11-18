# Basic Dating App

`go 1.18`

## Run the App

`go run cmd/main.go`

Application will start at `http://localhost:3000/`

## Limitations

This app was created as a tech test, in the interest of time some things have been left as essentially dummy code, to enable producing most of the solution in a shorter space of time, documenting the main limitations here:

- Only Handler layer logic has been implemented properly, the DB layer simply responds with dummy code as opposed to actually querying the DB and returning a real response. The method signatures themselves would remain unchanged and the DB layer could and would be implemented independently
- No testing is done on any part of the application
- Passwords are not encrypted etc. (they're not even stored) and just compared as plain text
- authentication is done using an `access-token` cookie, rather than using an `Authorization` header, not necessarily a limitation just clarification of the feature
- filtering results not implemented

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
