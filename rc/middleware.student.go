package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ensureActiveStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok, err := extractStudentRCID(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "student not logged in"})
			return
		}

		ctx.Set("student_rc_id", id)
	}
}

func getStudentRCID(ctx *gin.Context) uint {
	return uint(ctx.GetInt("student_rc_id"))
}
