package company

import "github.com/gin-gonic/gin"

func getAllHR(ctx *gin.Context, HRs *[]CompanyHR, cid uint) error {
	tx := db.WithContext(ctx).Where("company_id = ?", cid).Find(HRs)
	return tx.Error
}

func addHR(ctx *gin.Context, HR *CompanyHR) error {
	tx := db.WithContext(ctx).Create(HR)
	return tx.Error
}

func deleteHR(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Delete(&CompanyHR{}, "id = ?", id)
	return tx.Error
}

func FetchCompanyIDByEmail(ctx *gin.Context, email string) (uint, error) {
	var hr CompanyHR
	tx := db.WithContext(ctx).Where("email = ?", email).First(&hr)
	return hr.CompanyID, tx.Error
}
