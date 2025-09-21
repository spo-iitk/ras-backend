package application

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
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
			"profile",
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
			"profile",
			"tentative_job_location",
			"job_description",
			"cost_to_company",
			"package_details",
			"bond_details",
			"medical_requirements",
			"additional_eligibility",
			"message_for_cordinator",
			"accommodation",
			"ppo_confirming_date",
			"min_hires",
			"total_hires",
			"skill_set",
			"cpi_criteria",
			"ctcinr",
			"ctcfr",
			"perks",
			"social_media",
			"type_of_org",
			"turnover",
			"website",
			"social_media",
			"postal_address",
			"total_employees",
			"internship_period",
			"head_office",
			"establishment_date",
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
			"profile",
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
			"profile",
			"cpi_cutoff",
		).
		Order("deadline DESC").
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
			"profile",
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
			"created_at",
			"deleted_at",
			"updated_at",
			"company_name",
			"eligibility",
			"profile",
			"role",
			"skill_set",
			"tentative_job_location",
			"internship_period",
			"job_description",
			"ctcinr",
			"ctcfr",
			"cost_to_company",
			"accommodation",
			"medical_requirements",
			"perks",
			"package_details",
			"bond_details",

			"base",
			"gross",
			"joining_bonus",
			"take_home",
			"relocation_bonus",
			"retention_bonus",
			"package_details",
			"bond_details",
			"first_ctc",
			"deductions",
		).
		First(jp)
	return tx.Error
}

func fetchProformaForEligibleStudent(ctx *gin.Context, rid uint, student *rc.StudentRecruitmentCycle, jps *[]Proforma) error {
	subQuery := db.WithContext(ctx).Model(&ApplicationResume{}).
		Where("student_recruitment_cycle_id = ?", student.ID).
		Select("proforma_id")

	tx := db.WithContext(ctx).
		Where(
			"recruitment_cycle_id = ? AND is_approved = ? AND deadline > ? AND cpi_cutoff <= ? AND id NOT IN (?)",
			rid, true, time.Now().UnixMilli(), student.CPI, subQuery,
		)
	var eligibilityClauses []string
	var eligibilityArgs []interface{}

	// Note: We add 1 because SQL SUBSTRING is 1-indexed
	eligibilityClauses = append(eligibilityClauses, "SUBSTRING(eligibility FROM ? FOR 1) = '1'")
	eligibilityArgs = append(eligibilityArgs, student.ProgramDepartmentID+1)

	eligibilityClauses = append(eligibilityClauses, "SUBSTRING(eligibility FROM ? FOR 1) = '1'")
	eligibilityArgs = append(eligibilityArgs, student.SecondaryProgramDepartmentID+1)

	if len(eligibilityClauses) > 0 {
		eligibilityCondition := strings.Join(eligibilityClauses, " OR ")
		tx = tx.Where(fmt.Sprintf("(%s)", eligibilityCondition), eligibilityArgs...)
	}

	// @AkshatGupta15
	err := tx.Select(
		"id",
		"company_name",
		"eligibility",
		"deadline",
		"role",
		"profile",
		"cpi_cutoff",
	).
		Order("deadline").
		Find(jps).Error

	return err
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
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("id = ?", jp.ID).
		Update("hide_details", jp.HideDetails).
		Update("action_taken_by", middleware.GetUserID(ctx))
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
	tx := db.WithContext(ctx).Where("id = ? AND company_recruitment_cycle_id = ?", pid, cid).Delete(&Proforma{})
	return tx.RowsAffected > 0, tx.Error
}

func firstOrCreatePPOProforma(ctx *gin.Context, jp *Proforma) error {
	tx := db.WithContext(ctx).
		Where("company_recruitment_cycle_id = ? AND role = ?", jp.CompanyRecruitmentCycleID, jp.Role).
		FirstOrCreate(jp)
	return tx.Error
}

func getEligibility(ctx *gin.Context, pid uint) (string, float64, uint, uint, error) {
	var proforma Proforma
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("id = ?", pid).First(&proforma)
	return proforma.Eligibility, proforma.CPICutoff, proforma.CompanyRecruitmentCycleID, proforma.Deadline, tx.Error
}

func fetchRolesCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&Proforma{}).Where("recruitment_cycle_id = ? AND is_approved = ?", rid, true).Count(&count)
	return int(count), tx.Error
}

func fetchRecruitedCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	queryProforma := db.WithContext(ctx).Model(&Proforma{}).
		Where("recruitment_cycle_id = ? AND is_approved = ?", rid, true).
		Select("id")
	queryEvents := db.WithContext(ctx).Model(&ProformaEvent{}).
		Where("name IN (?) AND proforma_id IN (?)", []EventType{Recruited, PIOPPOACCEPTED}, queryProforma).
		Select("id")
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Where("proforma_event_id IN (?)", queryEvents).
		Count(&count)
	return int(count), tx.Error
}
