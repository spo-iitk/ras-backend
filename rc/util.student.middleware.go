package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ensureActiveStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _, err := extractStudentRCID(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("student_rc_id", int(id))
	}
}

func getStudentRCID(ctx *gin.Context) uint {
	return uint(ctx.GetInt("student_rc_id"))
}
