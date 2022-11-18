package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/cwhight/go-muzz/internal/model"
)

const (
    accessTokenCookieName  = "access-token"
	jwtSecretKey = "some-secret-key"
)

func GetJWTSecret() string {
	return jwtSecretKey
}

type Claims struct {
	Profile model.Profile `json:"profile"`
	jwt.StandardClaims
}

func ParseCookie(cookie string) (*model.Profile, error) {
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(GetJWTSecret()), nil
    })

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)
	return &claims.Profile, nil
}

func GenerateTokensAndSetCookies(user *model.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)

	return nil
}

func generateAccessToken(user *model.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user *model.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		Profile:  user.Profile,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func JWTErrorChecker(err error, c echo.Context) error {
	return echo.NewHTTPError(http.StatusUnauthorized, "unauthorised access")
}
