package rc

import "github.com/gin-gonic/gin"

func fetchAllRCs(ctx *gin.Context, rc *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Find(&rc)
	return tx.Error
}

func createRC(ctx *gin.Context, rc *RecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(rc)
	return tx.Error
}

func fetchRC(ctx *gin.Context, rid string, rc *RecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id = ?", rid).First(rc)
	return tx.Error
}
