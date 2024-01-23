package middleware

import (
	"net/http"
	"service-order/pkg"
	"service-order/schemas"

	"github.com/gin-gonic/gin"

)

func Auth() gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {

		var errorResponse schemas.SchemaUnathorizatedError

		errorResponse.Status = "Forbidden"
		errorResponse.Code = http.StatusForbidden
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "Authorization is required for this endpoint"

		if ctx.GetHeader("Authorization") == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, errorResponse)
		}

		token, err := pkg.VerifyTokenHeader(ctx, "JWT_SECRET")

		errorResponse.Status = "Unathorizated"
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "accessToken invalid or expired"

		if err != nil {
			defer ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
		} else {
			ctx.Set("user", token.Claims)
			ctx.Next()
		}
	})
}
