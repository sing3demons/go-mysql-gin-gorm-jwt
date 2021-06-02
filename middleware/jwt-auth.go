package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sing3demons/golanh-api/helper"
	"github.com/sing3demons/golanh-api/service"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Panicln("Claim[user_id]: ", claims["user_id"])
			log.Panicln("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Panicln(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}
