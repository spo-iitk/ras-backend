package application

import "github.com/gin-gonic/gin"

func fetchProformaByCompanyRC(ctx *gin.Context, cid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", cid).
		Select(
			"id",
			"created_at",
			"updated_at",
			"deleted_at",
			"eligibility",
			"company_id",
			"company_recruitment_cycle_id",
			"recruitment_cycle_id",
			"is_approved",
			"action_taken_by",
			"set_deadline",
			"hide_details",
			"active_hr_id",
			"nature_of_business",
			"tentative_job_location").
		Find(jps)
	return tx.Error
}

func fetchProformaByRC(ctx *gin.Context, rid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).
		Select(
			"id",
			"created_at",
			"updated_at",
			"deleted_at",
			"eligibility",
			"company_id",
			"company_recruitment_cycle_id",
			"recruitment_cycle_id",
			"is_approved",
			"set_deadline",
			"hide_details",
			"active_hr_id",
			"nature_of_business",
			"tentative_job_location").
		Find(jps)
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

func updateHideProforma(ctx *gin.Context, jp *hideProformaRequest) error {
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("id = ?", jp.ID).Update("hide_details", jp.HideDetails)
	return tx.Error
}

func updateProformaForCompany(ctx *gin.Context, jp *Proforma) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ? AND company_recruitment_cycle_id = ?", jp.ID, jp.CompanyRecruitmentCycleID).Updates(jp)
	return tx.RowsAffected == 1, tx.Error
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

func getEligibility(ctx *gin.Context, pid uint) (string, error) {
	var proforma Proforma
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("id = ?", pid).First(&proforma)
	return proforma.Eligibility, tx.Error
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
