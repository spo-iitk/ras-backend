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

	user_email := middleware.GetUserID(ctx)
	if user_email == "" {
		return 0, errors.New("unauthorized")
	}

	sIDs, err := rc.FetchStudentRCIDs(ctx, rid, []string{user_email})
	if err != nil {
		return 0, err
	}

	if len(sIDs) != 1 || sIDs[0] == 0 {
		return 0, errors.New("RCID not found")
	}

	return sIDs[0], nil
}
