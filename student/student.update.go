package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func updateStudentHandler(ctx *gin.Context) {
	var updateStudentRequest Student

	if err := ctx.ShouldBindJSON(&updateStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := middleware.GetUserID(ctx)

	if updateStudentRequest.SecondaryProgramDepartmentID > updateStudentRequest.ProgramDepartmentID && util.IsDoubleMajor(updateStudentRequest.SecondaryProgramDepartmentID) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Secondary program department and primary program department seems to be interchanged"})
		return
	}

	updated, err := updateStudentByEmail(ctx, &updateStudentRequest, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Student not found or forbidden"})
		return
	}

	logrus.Infof("A student with email %s is updated", email)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}
