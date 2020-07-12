# Golang REST api with jwt tokens

## Heroku link to app:
https://go-rest-jwt.herokuapp.com/

## Endpoints app

### UI:
/
/refresh
/remove

### API:
/api/users
/api/tokens

/api/token/get/{userId}
/api/token/refresh
/api/token/remove/t/{tokenId}
/api/token/remove/u/{userId}

### Used libraries:
github.com/joho/godotenv
github.com/sirupsen/logrus
github.com/gorilla/handlers
github.com/gorilla/mux
github.com/dgrijalva/jwt-go
go.mongodb.org/mongo-driver
