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
	Emails  []string `json:"emails" binding:"required"` // this is now roll no too
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

		var evnt ProformaEvent
		err = fetchEvent(ctx, req.EventID, &evnt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var srcIDs []uint
		srcIDs, req.Emails, err = rc.FetchStudentRCIDs(ctx, rid, req.Emails)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if evnt.Name == string(ApplicationSubmitted) {
			if len(srcIDs) != 1 {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Only one student can be force enrolled at a time"})
				return
			}

			rsid, resume, err := rc.FetchFirstResume(ctx, srcIDs[0])
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Resume not found" + err.Error()})
				return
			}

			err = createApplicationResume(ctx, &ApplicationResume{
				StudentRecruitmentCycleID: srcIDs[0],
				ProformaID:                pid,
				ResumeID:                  rsid,
				Resume:                    resume,
			})
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
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

		if evnt.Name == string(PIOPPOACCEPTED) || evnt.Name == string(Recruited) {
			err = rc.UpdateStudentType(ctx, proforma.CompanyRecruitmentCycleID, req.Emails, string(Recruited))
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			msg := "Dear student" + "\n\nCongratulations\n"
			msg += "You have been recruited by " + proforma.CompanyName
			msg += " in the profile of " + proforma.Profile

			mail_channel <- mail.GenerateMails(req.Emails, "Congratulations", msg)
		} else if evnt.Name == string(WALKIN) {
			msg := "Dear student" + "\n\n"
			msg += "You have been selected for the walk in interview for the profile " + proforma.Profile + " by " + proforma.CompanyName + "."

			mail_channel <- mail.GenerateMails(req.Emails, "Update on Application", msg)
		} else {
			msg := "Dear student" + "\n\n"
			msg += "You have advanced to the stage of " + evnt.Name + " for the recruitment process of profile "
			msg += proforma.Profile + " by " + proforma.CompanyName + "."

			mail_channel <- mail.GenerateMails(req.Emails, "Update on Application", msg)
		}

		ctx.JSON(http.StatusOK, gin.H{"success": "Students added successfully"})
	}
}

func deleteStudentFromEventHandler(ctx *gin.Context) {
	eventID, err := util.ParseUint(ctx.Param("eid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var evnt ProformaEvent
	err = fetchEvent(ctx, eventID, &evnt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentID, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	rcID, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid r ID"})
		return
	}

	if evnt.Name == string(PIOPPOACCEPTED) || evnt.Name == string(Recruited) {
		if err := rc.UnRecruitStudent(ctx, studentID, rcID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := deleteStudentFromEvent(ctx, eventID, studentID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student deleted from event successfully"})
}

func deleteAllStudentsFromEventHandler(ctx *gin.Context) {
	eventID, err := util.ParseUint(ctx.Param("eid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var evnt ProformaEvent
	err = fetchEvent(ctx, eventID, &evnt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sids, _ := getStudentIDByEventID(ctx, eventID)

	if evnt.Name == string(PIOPPOACCEPTED) || evnt.Name == string(Recruited) {
		if err := rc.UnRecruitAll(ctx, sids); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := deleteAllStudentsFromEvent(ctx, eventID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All students deleted from event successfully"})
}
