package rc

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getAllNoticesHandler(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var notices []Notice

	err := fetchAllNotices(ctx, rid, &notices)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notices)
}

func postNoticeHandler(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var notice Notice

	err := ctx.ShouldBindJSON(&notice)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseUint(rid, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notice.RecruitmentCycleID = uint(id)
	notice.CreatedBy = middleware.GetUserID(ctx)

	err = createNotice(ctx, &notice)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notice)
}

func deleteNoticeHandler(ctx *gin.Context) {
	nid := ctx.Param("nid")

	err := removeNotice(ctx, nid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "status"})
}

func postReminderHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid := ctx.Param("rid")
		nid := ctx.Param("nid")

		var notice Notice
		err := fetchNotice(ctx, nid, &notice)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if notice.LastReminderAt > time.Now().Add(-6*time.Hour).UnixMilli() {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Reminder already sent"})
			return
		}

		notice.LastReminderAt = time.Now().UnixMilli()
		err = updateNotice(ctx, &notice)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var students []StudentRecruitmentCycle

		err = FetchAllStudents(ctx, rid, &students)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var emails []string

		for _, student := range students {
			emails = append(emails, student.Email)
		}

		mail_channel <- mail.GenerateMails(emails, "Notice: "+notice.Title, notice.Description)

		ctx.JSON(http.StatusOK, gin.H{"status": "mail sent"})
	}
}
