package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/student"
	"github.com/spo-iitk/ras-backend/util"
)

func getAllStudentsHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var students []StudentRecruitmentCycle

	err = fetchAllStudents(ctx, rid, &students)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func getStudentHandler(ctx *gin.Context) {
	srid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var student StudentRecruitmentCycle

	err = FetchStudent(ctx, srid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func putStudentHandler(ctx *gin.Context) {
	var student StudentRecruitmentCycle

	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if student.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter student ID"})
		return
	}

	ok, err := updateStudent(ctx, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

type bulkFreezeStudentRequest struct {
	Emails []string `json:"email"`
	Frozen bool     `json:"frozen"`
}

func bulkFreezeStudentsHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req bulkFreezeStudentRequest

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ok, err := freezeStudentsToggle(ctx, req.Emails, req.Frozen)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No such student exists"})
			return
		}

		user := middleware.GetUserID(ctx)

		logrus.Infof("%v froze %v students", user, len(req.Emails))

		msg := "Dear student" + "\n\n"
		msg += "Your account has been "
		if req.Frozen {
			msg += "FROZEN"
		} else {
			msg += "UNFROZEN"
		}

		msg += " by the coordinators.\n\n"

		mail_channel <- mail.GenerateMails(req.Emails, "Action taken on Account", msg)

		ctx.JSON(http.StatusOK, gin.H{"status": "froze students"})
	}
}

type bulkPostStudentRequest struct {
	Email []string `json:"email" binding:"required"`
}

func postStudentsHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var emails bulkPostStudentRequest

		err = ctx.ShouldBindJSON(&emails)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		emailArr := emails.Email
		var students []StudentRecruitmentCycle
		var studentsGlobal []student.Student
		var regEmails []string

		err = student.FetchStudents(ctx, &studentsGlobal, emailArr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, student := range studentsGlobal {
			var secondaryProgramDepartmentID uint = 0
			if util.IsDoubleMajor(student.SecondaryProgramDepartmentID) {
				secondaryProgramDepartmentID = student.SecondaryProgramDepartmentID
			}

			regEmails = append(regEmails, student.IITKEmail)
			students = append(students, StudentRecruitmentCycle{
				RecruitmentCycleID:           rid,
				StudentID:                    student.ID,
				Email:                        student.IITKEmail,
				Name:                         student.Name,
				RollNo:                       student.RollNo,
				CPI:                          student.CurrentCPI,
				ProgramDepartmentID:          student.ProgramDepartmentID,
				SecondaryProgramDepartmentID: secondaryProgramDepartmentID,
			})
		}

		err = createStudents(ctx, &students)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMails(regEmails, "Registered in Recruitment Cycle", "Dear student,\n\nYou have been registered in a Recruitment Cycle.\n\n Please answer enrollment questions and proceed to the next steps.")

		user := middleware.GetUserID(ctx)
		num := len(students)
		reqNum := len(emailArr)

		logrus.Infof("%v added %v new students to RC %d", user, num, rid)

		if num != reqNum {
			ctx.JSON(http.StatusOK, gin.H{"status": "partially added student"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "added students"})
	}
}

func deleteStudentHandler(ctx *gin.Context) {
	srid := ctx.Param("sid")

	err := deleteStudent(ctx, srid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v deleted %v from RC", user, srid)

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted student"})
}
