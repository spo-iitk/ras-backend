package rc

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getStudentEnrollment(ctx *gin.Context) {
	rid := ctx.Param("rid")

	sid, err := GetStudentRecruitmentCycleID(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var questions []RecruitmentCycleQuestion
	var answers []RecruitmentCycleQuestionsAnswer

	err = fetchStudentQuestions(ctx, rid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = fetchStudentAnswers(ctx, strconv.FormatUint(uint64(sid), 10), &answers)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"questions": questions, "answers": answers})
}

func GetStudentRecruitmentCycleID(ctx *gin.Context, rid string) (uint, error) {
	var student StudentRecruitmentCycle

	email := middleware.GetUserID(ctx)

	err := fetchStudent(ctx, email, rid, &student)
	if err != nil {
		return 0, err
	}

	return student.ID, err
}

func postEnrollmentAnswer(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var answer RecruitmentCycleQuestionsAnswer

	err := ctx.BindJSON(&answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	answer.StudentRecruitmentCycleID, err = GetStudentRecruitmentCycleID(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = createStudentAnswer(ctx, &answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	aid := answer.ID
	ctx.JSON(200, gin.H{"data": aid})
}
