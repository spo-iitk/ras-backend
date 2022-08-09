package application

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func extractStudentRCID(ctx *gin.Context) (uint, error) {
	rid_string := ctx.Param("rid")
	rid, err := util.ParseUint(rid_string)
	if err != nil {
		return 0, err
	}

	if !rc.IsRCActive(ctx, rid) {
		return 0, errors.New("recruitment cycle is not active")
	}

	user_email := middleware.GetUserID(ctx)
	if user_email == "" {
		return 0, errors.New("unauthorized")
	}

	studentrcid, err := rc.FetchStudentRCID(ctx, rid, user_email)
	if err != nil {
		return 0, err
	}

	if studentrcid == 0 {
		return 0, errors.New("RCID not found")
	}

	return studentrcid, nil
}
