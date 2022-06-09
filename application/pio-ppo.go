package application

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
)

func getEmptyProformaByCID(ctx *gin.Context, cid uint, jp *JobProforma) error {
	var companyRC rc.CompanyRecruitmentCycle
	err := rc.FetchCompanyByID(ctx, cid, &companyRC)
	if err != nil {
		return err
	}

	jp.CompanyRecruitmentCycleID = companyRC.ID
	jp.RecruitmentCycleID = companyRC.RecruitmentCycleID
	jp.CompanyID = companyRC.CompanyID
	jp.IsApproved = sql.NullBool{}

	return firstOrCreateEmptyPerfoma(ctx, jp)
}

type pioppoRequest struct {
	cid    uint
	emails []string
}

func postPPOPIOHandler(ctx *gin.Context) {
	var req pioppoRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = rc.UpdateStudentType(ctx, req.cid, req.emails)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp JobProforma
	err = getEmptyProformaByCID(ctx, req.cid, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	studentIDs, err := rc.FetchStudentRCIDs(ctx, jp.RecruitmentCycleID, req.emails)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var event = JobProformaEvent{
		JobProformaID: jp.ID,
		Name:          "PIO-PPO",
	}
	err = createEvent(ctx, &event)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var ses []EventStudent

	for _, studentID := range studentIDs {
		ses = append(ses, EventStudent{
			JobProformaEventID:        event.ID,
			StudentRecruitmentCycleID: studentID,
			Present:                   true,
		})
	}

	err = createStudentEvents(ctx, &ses)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "updated student pioppo"})
}
