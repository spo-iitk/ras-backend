package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/student"
	"github.com/spo-iitk/ras-backend/util"
)

func getAllStudents(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var students []StudentRecruitmentCycle

	err := fetchAllStudents(ctx, rid, &students)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func getStudentByID(ctx *gin.Context) {
	rid := ctx.Param("rid")
	srid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var student StudentRecruitmentCycle

	err = fetchStudentByID(ctx, srid, rid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func putStudent(ctx *gin.Context) {
	var student StudentRecruitmentCycle

	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if student.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter student ID"})
		return
	}

	ok, err := updateStudent(ctx, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No such student exists"})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v updated a student with id %d", user, student.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "updated student"})
}

type bulkPostStudentRequest struct {
	Email []string `json:"email" binding:"required"`
}

func postStudents(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var emails bulkPostStudentRequest

	err = ctx.ShouldBindJSON(&emails)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	emailArr := emails.Email
	var students []StudentRecruitmentCycle
	var studentsGlobal []student.Student

	err = student.FetchStudents(ctx, &studentsGlobal, emailArr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, student := range studentsGlobal {
		students = append(students, StudentRecruitmentCycle{
			RecruitmentCycleID:           rid,
			StudentID:                    student.ID,
			Email:                        student.IITKEmail,
			Name:                         student.Name,
			CPI:                          student.CurrentCPI,
			ProgramDepartmentID:          student.ProgramDepartmentID,
			SecondaryProgramDepartmentID: student.SecondaryProgramDepartmentID,
		})
	}

	err = createStudents(ctx, &students)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)
	num := len(students)

	logrus.Infof("%v added %v new students", user, num)

	ctx.JSON(http.StatusOK, gin.H{"status": "added students"})
}

func deleteStudentByID(ctx *gin.Context) {
	srid := ctx.Param("sid")

	err := deleteStudent(ctx, srid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v deleted %v from RC", user, srid)

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted student"})
}
