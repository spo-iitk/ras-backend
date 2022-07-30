package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func getStudentsByEventHandler(ctx *gin.Context) {
	eid, err := util.ParseUint(ctx.Param("eid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var students []EventStudent
	err = fetchStudentsByEvent(ctx, eid, &students)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentRCIDs []uint
	for _, student := range students {
		studentRCIDs = append(studentRCIDs, student.StudentRecruitmentCycleID)
	}

	var studentRCs []rc.StudentRecruitmentCycle
	err = rc.FetchStudents(ctx, studentRCIDs, &studentRCs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, studentRCs)
}

type postStudentsByEventRequest struct {
	EventID uint     `json:"event_id" binding:"required"`
	Emails  []string `json:"emails" binding:"required"`
}

func postStudentsByEventHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pid, err := util.ParseUint(ctx.Param("pid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var proforma Proforma
		err = fetchProforma(ctx, pid, &proforma)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req postStudentsByEventRequest
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		srcIDs, err := rc.FetchStudentRCIDs(ctx, rid, req.Emails)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var students []EventStudent
		for _, srcID := range srcIDs {
			students = append(students, EventStudent{
				ProformaEventID:           req.EventID,
				CompanyRecruitmentCycleID: proforma.CompanyRecruitmentCycleID,
				StudentRecruitmentCycleID: srcID,
			})
		}

		err = createEventStudents(ctx, &students)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var evnt ProformaEvent
		err = fetchEvent(ctx, req.EventID, &evnt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if evnt.Name == string(PIOPPOACCEPTED) || evnt.Name == string(Recruited) {
			err = rc.UpdateStudentType(ctx, proforma.CompanyRecruitmentCycleID, req.Emails, string(Recruited))
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		msg := "Dear student" + "\n\n"
		msg += "You have advanced to the stage of " + evnt.Name + " for the recruitment process of role "
		msg += proforma.Role + " by " + proforma.CompanyName + "."

		mail_channel <- mail.GenerateMails(req.Emails, "Update on Application", msg)

		if len(req.Emails) == len(students) {
			ctx.JSON(http.StatusOK, gin.H{"success": "Students added successfully"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"success": "Students added successfully but Some students were not added"})
		}
	}
}
