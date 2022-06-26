package application

import (
	"github.com/gin-gonic/gin"
)

func fetchProformasByCompanyForAdmin(ctx *gin.Context, cid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", cid).
		Select(
			"id",
			"created_at",
			"updated_at",
			"deleted_at",
			"company_recruitment_cycle_id",
			"recruitment_cycle_id",
			"company_id",
			"company_name",
			"is_approved",
			"action_taken_by",
			"eligibility",
			"cpi_cutoff",
			"deadline",
			"hide_details",
			"active_hr",
			"role",
			"tentative_job_location",
		).
		Find(jps)
	return tx.Error
}

func fetchProformaForCompany(ctx *gin.Context, pid uint, cid uint, jp *Proforma) error {
	tx := db.WithContext(ctx).Where("id = ? AND company_recruitment_cycle_id=?", pid, cid).
		Select(
			"id",
			"created_at",
			"updated_at",
			"company_recruitment_cycle_id",
			"recruitment_cycle_id",
			"company_id",
			"company_name",
			"is_approved",
			"eligibility",
			"deadline",
			"hide_details",
			"active_hr",
			"role",
			"tentative_job_location",
			"job_description",
			"cost_to_company",
			"package_details",
			"bond_details",
			"medical_requirements",
			"additional_eligibility",
			"message_for_cordinator",
		).
		Find(jp)
	return tx.Error
}

func fetchProformasByCompanyForCompany(ctx *gin.Context, cid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", cid).
		Select(
			"id",
			"created_at",
			"updated_at",
			"company_recruitment_cycle_id",
			"recruitment_cycle_id",
			"company_id",
			"company_name",
			"is_approved",
			"eligibility",
			"deadline",
			"hide_details",
			"active_hr",
			"role",
			"tentative_job_location",
		).
		Find(jps)
	return tx.Error
}

func fetchProformasForStudent(ctx *gin.Context, rid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).
		Where("recruitment_cycle_id = ? AND is_approved = ? AND deadline > 0", rid, true).
		Select(
			"id",
			"company_name",
			"eligibility",
			"deadline",
			"role",
			"cpi_cutoff",
		).
		Find(jps)
	return tx.Error
}

func fetchProformaByRCForAdmin(ctx *gin.Context, rid uint, jps *[]Proforma) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).
		Select(
			"id",
			"created_at",
			"updated_at",
			"deleted_at",
			"company_recruitment_cycle_id",
			"recruitment_cycle_id",
			"company_id",
			"company_name",
			"is_approved",
			"action_taken_by",
			"eligibility",
			"cpi_cutoff",
			"deadline",
			"hide_details",
			"active_hr",
			"role",
			"tentative_job_location",
		).
		Order("updated_at DESC").
		Find(jps)
	return tx.Error
}

func fetchProforma(ctx *gin.Context, pid uint, jp *Proforma) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).First(jp)
	return tx.Error
}

func fetchProformaForStudent(ctx *gin.Context, pid uint, jp *Proforma) error {
	tx := db.WithContext(ctx).
		Where("id = ? AND is_approved = ? AND deadline > 0", pid, true).
		Select(
			"id",
			"company_name",
			"eligibility",
			"deadline",
			"role",
			"cpi_cutoff",
			"tentative_job_location",
			"job_description",
			"cost_to_company",
			"package_details",
			"bond_details",
			"medical_requirements",
		).
		First(jp)
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
	tx := db.WithContext(ctx).Where("id = ?", pid).Delete(&Proforma{})
	return tx.Error
}

func deleteProformaByCompany(ctx *gin.Context, pid uint, cid uint) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ? AND company_recruitment_cycle_id = ?", pid, cid).Delete(Proforma{})
	return tx.RowsAffected > 0, tx.Error
}

func firstOrCreateEmptyPerfoma(ctx *gin.Context, jp *Proforma) error {
	tx := db.WithContext(ctx).
		Where("company_recruitment_cycle_id = ? AND role = ?", jp.CompanyRecruitmentCycleID, jp.Role).
		FirstOrCreate(jp)
	return tx.Error
}

func getEligibility(ctx *gin.Context, pid uint) (string, float64, uint, error) {
	var proforma Proforma
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("id = ?", pid).First(&proforma)
	return proforma.Eligibility, proforma.CPICutoff, proforma.CompanyRecruitmentCycleID, tx.Error
}

func fetchRolesCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("recruitment_cycle_id = ? AND is_approved = ?", rid, true).Count(&count)
	return int(count), tx.Error
}

func fetchRecruitedCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Joins("JOIN proformas ON proformas.id=proforma_events.proforma_id").Model(&ProformaEvent{}).
		Where("name IN ? AND recruitment_cycle_id = ?", []EventType{Recruited, PIOPPOACCEPTED}, rid).Count(&count)
	return int(count), tx.Error
}
