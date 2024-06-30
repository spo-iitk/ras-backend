package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
)

func postStudentDocumentHandler(ctx *gin.Context) {
	var document StudentDocument
	email := middleware.GetUserID(ctx)

	if err := ctx.ShouldBindJSON(&document); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var student Student
	err := getStudentByEmail(ctx, &student, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	document.StudentID = student.ID
	err = saveDocument(ctx, &document)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("Document for student %d uploaded", student.ID)
	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully uploaded document"})
}

func getStudentDocumentHandler(ctx *gin.Context) {
	email := middleware.GetUserID(ctx)
	var student Student
	err := getStudentByEmail(ctx, &student, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var documents []StudentDocument
	err = getDocumentsByStudentID(ctx, &documents, student.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, documents)
}
