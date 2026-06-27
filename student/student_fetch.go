package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getStudentResumeHandler(ctx *gin.Context) {
	sid, err := strconv.ParseUint(ctx.Param("sid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var documents []StudentDocument
	err = getDocumentsByStudentID(ctx, &documents, uint(sid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(documents) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No documents found"})
		return
	}

	ctx.JSON(http.StatusOK, documents)
}
