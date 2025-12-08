package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AMFarhan21/fres"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func forbidden(e echo.Context) error {
	return e.JSON(http.StatusUnauthorized, fres.Response.StatusUnauthorized(http.StatusUnauthorized))
}

func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			header := strings.Split(e.Request().Header.Get("Authorization"), " ")

			if len(header) < 2 {
				log.Print("Error on len(header) middleware")
				return forbidden(e)
			}

			if header[0] != "Bearer" {
				log.Print("Error on != Bearer middleware")
				return forbidden(e)
			}

			claims := jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(header[1], claims, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					log.Printf("unexpected signing method %v", t.Header["alg"])
					return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
				}

				return []byte(jwtSecret), nil
			})
			if err != nil {
				log.Printf("Error on ParseWithClaims %s:", err.Error())
				return forbidden(e)
			}

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok && method != jwt.SigningMethodHS256 {
				log.Print("Error on token method")
				return forbidden(e)
			}

			expAt, err := claims.GetExpirationTime()
			if err != nil {
				log.Printf("Error on middleware expAt %s:", err.Error())
				return forbidden(e)
			}

			if time.Now().After(expAt.Time) {
				log.Print("Your token already expired")
				return forbidden(e)
			}

			user_id := claims["id"].(string)
			role := claims["role"].(string)

			e.Set("id", user_id)
			e.Set("role", role)

			return next(e)
		}
	}
}

func ACLMiddleware(rolesMap map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			role, _ := e.Get("role").(string)

			if rolesMap[role] {
				return next(e)
			}

			return forbidden(e)
		}
	}
}
