package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

type DocumentRequest struct {
    Document     string       `json:"document" binding:"required"`
    DocumentType DocumentType `json:"document_type" binding:"required"`
}

func postStudentDocumentHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request DocumentRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}

	err = addStudentDocument(ctx, request.Document, sid, rid, request.DocumentType)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Document Added Successfully"})
}
func getStudentDocumentHandler(ctx *gin.Context) {
	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}
	var documents []StudentDocument
	err := fetchStudentDocuments(ctx, sid, &documents)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}   

	ctx.JSON(http.StatusOK, documents)
}
