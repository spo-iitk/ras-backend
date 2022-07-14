package application

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func getEmptyProformaByCID(ctx *gin.Context, cid uint, jp *Proforma) error {
	var companyRC rc.CompanyRecruitmentCycle
	err := rc.FetchCompany(ctx, cid, &companyRC)
	if err != nil {
		return err
	}

	jp.CompanyRecruitmentCycleID = companyRC.ID
	jp.RecruitmentCycleID = companyRC.RecruitmentCycleID
	jp.CompanyName = companyRC.CompanyName
	jp.IsApproved = sql.NullBool{Bool: false, Valid: true}
	jp.ActionTakenBy = middleware.GetUserID(ctx)
	jp.Role = string(PIOPPOACCEPTED)

	return firstOrCreateEmptyPerfoma(ctx, jp)
}

type pioppoRequest struct {
	Cid    uint     `json:"cid" binding:"required"`
	Emails []string `json:"emails" binding:"required"`
}

func postPPOPIOHandler(ctx *gin.Context) {
	var req pioppoRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = rc.UpdateStudentType(ctx, req.Cid, req.Emails, string(PIOPPOACCEPTED))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma
	err = getEmptyProformaByCID(ctx, req.Cid, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	studentIDs, err := rc.FetchStudentRCIDs(ctx, rid, req.Emails)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event = ProformaEvent{
		ProformaID: jp.ID,
		Name:       "PIO-PPO",
	}
	err = createEvent(ctx, &event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ses []EventStudent

	for _, studentID := range studentIDs {
		ses = append(ses, EventStudent{
			ProformaEventID:           event.ID,
			StudentRecruitmentCycleID: studentID,
			Present:                   true,
		})
	}

	err = createEventStudents(ctx, &ses)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated student pioppo"})
}
