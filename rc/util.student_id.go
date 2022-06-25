package rc

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func extractStudentRCID(ctx *gin.Context) (uint, bool, error) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		return 0, false, err
	}

	email := middleware.GetUserID(ctx)

	var student StudentRecruitmentCycle
	err = fetchStudentByEmailAndRC(ctx, email, rid, &student)
	if err != nil {
		return 0, false, err
	}

	if student.IsFrozen {
		return 0, false, errors.New("student frozen")
	}

	return student.ID, student.IsVerified, err
}
