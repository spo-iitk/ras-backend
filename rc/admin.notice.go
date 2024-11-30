package rc

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
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

func postNoticeHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var notice Notice
		err = ctx.ShouldBindJSON(&notice)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = CreateNotice(ctx, rid, &notice)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// As a custom plugin by RAS God Harshit Raj, feel free to remove this after 2023. :P
		// plugins.NewNoticeNotification(mail_channel, notice.ID, notice.RecruitmentCycleID, notice.Title, notice.Description, notice.CreatedBy)

		ctx.JSON(http.StatusOK, gin.H{"status": "notice created"})
	}
}

func CreateNotice(ctx *gin.Context, id uint, notice *Notice) error {
	notice.RecruitmentCycleID = uint(id)
	notice.LastReminderAt = 0
	notice.CreatedBy = middleware.GetUserID(ctx)
	return createNotice(ctx, notice)
}

func putNoticeHandler(ctx *gin.Context) {
	var editNoticeRequest Notice

	err := ctx.ShouldBindJSON(&editNoticeRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if editNoticeRequest.RecruitmentCycleID != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recruitment cycle id is not allowed"})
		return
	}

	if editNoticeRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	err = updateNotice(ctx, &editNoticeRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, editNoticeRequest)
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
		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		nid := ctx.Param("nid")

		var notice Notice
		err = fetchNotice(ctx, nid, &notice)
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

		emails, err := fetchAllUnfrozenEmails(ctx, rid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mailBody := notice.Description
		if notice.Deadline > 0 {
			deadlineTime := time.Unix(int64(notice.Deadline)/1000, 0)
                        deadlineStr := deadlineTime.Format("02 Jan 2006 15:04")
			mailBody += "\n\nDeadline: " + deadlineStr
		}

		mail_channel <- mail.GenerateMails(emails, "Notice: "+notice.Title, mailBody)

		ctx.JSON(http.StatusOK, gin.H{"status": "mail sent"})
	}
}
