package application

import "github.com/gin-gonic/gin"

func fetchProformaByCompanyRC(ctx *gin.Context, cid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", cid).Find(jps)
	return tx.Error
}

func fetchProformaByRC(ctx *gin.Context, rid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(jps)
	return tx.Error
}

func fetchProforma(ctx *gin.Context, pid uint, jp *Proforma) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).First(jp)
	return tx.Error
}

func createProforma(ctx *gin.Context, jp *Proforma) error {
	tx := db.WithContext(ctx).Create(jp)
	return tx.Error
}

func updateProforma(ctx *gin.Context, jp *Proforma) error {
	tx := db.WithContext(ctx).Where("id = ?", jp.ID).Updates(jp)
	return tx.Error
}

func deleteProforma(ctx *gin.Context, pid uint) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).Delete(Proforma{})
	return tx.Error
}

func firstOrCreateEmptyPerfoma(ctx *gin.Context, jp *Proforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", jp.CompanyRecruitmentCycleID).FirstOrCreate(jp)
	return tx.Error
}
