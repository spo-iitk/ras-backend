package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func updateStudentByIDHandler(ctx *gin.Context) {
	var updateStudentRequest Student

	if err := ctx.ShouldBindJSON(&updateStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateStudentRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter student ID"})
		return
	}

	if updateStudentRequest.SecondaryProgramDepartmentID > updateStudentRequest.ProgramDepartmentID && util.IsDoubleMajor(updateStudentRequest.SecondaryProgramDepartmentID) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Secondary program department and primary program department seems to be interchanged"})
		return
	}

	updated, err := updateStudentByID(ctx, &updateStudentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	logrus.Infof("A student with id %d is updated by %s", updateStudentRequest.ID, ctx.GetString("userID"))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func verifyStudentHandler(ctx *gin.Context) {
	var verifyStudentRequest Student

	if err := ctx.ShouldBindJSON(&verifyStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verifyStudentRequest.ID = sid
	updated, err := verifyStudent(ctx, &verifyStudentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if verifyStudentRequest.IsVerified {
		logrus.Infof("A student with id %d is verified", verifyStudentRequest.ID)
		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully verified"})
	} else {
		logrus.Infof("A student with id %d is unverified", verifyStudentRequest.ID)
		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully unverified"})
	}
}

func makeStudentEdiatableHandler(ctx *gin.Context) {
	var editableStudentRequest Student

	if err := ctx.ShouldBindJSON(&editableStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = updateIsEditableWithID(ctx, sid, true)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("A student with id %d is made editable by %v", editableStudentRequest.ID, user)
	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully made student editable"})
}
