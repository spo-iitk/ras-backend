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

func updateProformaForCompany(ctx *gin.Context, jp *Proforma) error {
	tx := db.WithContext(ctx).Where("id = ? AND company_recruitment_cycle_id = ?", jp.ID, jp.CompanyRecruitmentCycleID).Updates(jp)
	return tx.Error
}

func deleteProforma(ctx *gin.Context, pid uint) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).Delete(Proforma{})
	return tx.Error
}

func deleteProformaByCompany(ctx *gin.Context, pid uint, cid uint) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ? AND company_recruitment_cycle_id = ?", pid, cid).Delete(Proforma{})
	return tx.RowsAffected > 0, tx.Error
}

func firstOrCreateEmptyPerfoma(ctx *gin.Context, jp *Proforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", jp.CompanyRecruitmentCycleID).FirstOrCreate(jp)
	return tx.Error
}

func getRolesCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("recruitment_cycle_id = ? AND is_approved = ?", rid, true).Count(&count)
	return int(count), tx.Error
}

func getPPOPIOCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Joins("JOIN proformas ON proformas.id=proforma_events.proforma_id").Model(&ProformaEvent{}).
		Where("name IN ? AND recruitment_cycle_id = ?", []EventType{Recruited, PIOPPOACCEPTED}, rid).Count(&count)
	return int(count), tx.Error
}
