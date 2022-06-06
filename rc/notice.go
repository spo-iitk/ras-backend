package rc

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getAllNotices(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var notices []Notice

	err := fetchAllNotices(ctx, rid, &notices)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": notices})
}

func postNotice(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var notice Notice

	err := ctx.BindJSON(&notice)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseUint(rid, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	notice.RecruitmentCycleID = uint(id)
	notice.CreatedBy = middleware.GetUserID(ctx)

	err = createNotice(ctx, &notice)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	nid := gin.H{"id": notice.ID}
	ctx.JSON(200, gin.H{"data": nid})
}

func deleteNotice(ctx *gin.Context) {
	nid := ctx.Param("nid")

	err := removeNotice(ctx, nid)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "success"})
}

func postReminder(ctx *gin.Context) {
	// rid := ctx.Param("rid")
	nid := ctx.Param("nid")

	var notice Notice
	err := fetchNotice(ctx, nid, &notice)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	notice.LastReminderAt = time.Now().UnixMilli()
	err = updateNotice(ctx, &notice)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	// !TODO: send email
	// get all student
	// mial then

	ctx.JSON(200, gin.H{"status": "mail send"})
}