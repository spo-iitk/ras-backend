package rc

import (
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
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, students)
}

func getStudent(ctx *gin.Context) {
	rid := ctx.Param("rid")
	email := middleware.GetUserID(ctx)
	var student StudentRecruitmentCycle

	err := fetchStudent(ctx, email, rid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, student)
}

func putStudent(ctx *gin.Context) {
	var student StudentRecruitmentCycle

	err := ctx.BindJSON(&student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateStudent(ctx, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v updated a student with id %d", user, student.ID)

	ctx.JSON(200, gin.H{"status": "updated student"})
}

type bulkPostStudentRequest struct {
	Email []string `json:"email"`
}

func postStudents(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var emails bulkPostStudentRequest

	err := ctx.BindJSON(&emails)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	emailArr := emails.Email
	var students []StudentRecruitmentCycle
	var studentsGlobal []student.Student

	err = student.FetchStudents(ctx, &studentsGlobal, emailArr)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, student := range studentsGlobal {
		students = append(students, StudentRecruitmentCycle{
			RecruitmentCycleID:           util.ToUint(rid),
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
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)
	num := len(students)

	logrus.Infof("%v addedd %v new students", user, num)

	ctx.JSON(200, gin.H{"status": "created students", "data": students})
}
