package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

/*
Admin Authentication function takes the []byte secret key and echo handler function as input and returns the echo handler function

This function allows only registered admins to do the required operations
*/
func AdminAuthorize(jwtSecret []byte, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := ExtractTokenFromHeader(c.Request())
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		fmt.Println("Claims", token.Claims)

		isAdmin := claims["admin"]
		if isAdmin == "admin" {
			c.Set("token", token)
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
	}
}

/*
Member Authentication function takes the []byte of secret key and echo handler function as input and returns the echo handler function

This function allows only registered members to place their orders
*/
func MemberAuthorize(jwtSecret []byte, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := ExtractTokenFromHeader(c.Request())
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		isAdmin := claims["member"]
		if isAdmin == "member" {
			c.Set("token", token)
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
	}
}

// ExtractTokenFromHeader function is used to extract the token part after the Bearer part
func ExtractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Expected format: "Bearer <token>"
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ""
	}

	return tokenParts[1]
}
