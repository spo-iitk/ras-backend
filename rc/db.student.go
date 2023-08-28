package rc

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func fetchAllStudents(ctx *gin.Context, rid uint, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(students)
	return tx.Error
}

func fetchAllUnfrozenEmails(ctx *gin.Context, rid uint) ([]string, error) {
	var emails []string
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).
		Where("recruitment_cycle_id = ? AND is_frozen = ?", rid, false).Pluck("email", &emails)
	return emails, tx.Error
}

func fetchStudentByEmailAndRC(ctx *gin.Context, email string, rid uint, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("email = ? AND recruitment_cycle_id = ?", email, rid).First(student)
	return tx.Error
}

func FetchStudent(ctx *gin.Context, sid uint, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).First(student, sid)
	return tx.Error
}

func FetchStudentBySRID(ctx *gin.Context, sid []uint, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id IN ?", sid).Find(students).Order("id ASC")
	return tx.Error
}

func updateStudent(ctx *gin.Context, student *StudentRecruitmentCycle) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", student.ID).Updates(student)
	return tx.RowsAffected > 0, tx.Error
}

func freezeStudentsToggle(ctx *gin.Context, ids []string, frozen bool) (bool, error) {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("email IN ? OR roll_no IN ?", ids, ids).Update("is_frozen", frozen)
	return tx.RowsAffected > 0, tx.Error
}

func deleteStudent(ctx *gin.Context, sid string) error {
	tx := db.WithContext(ctx).Where("id = ?", sid).Delete(&StudentRecruitmentCycle{})
	return tx.Error
}

func createStudents(ctx *gin.Context, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(students)
	return tx.Error
}

func UpdateStudentType(ctx *gin.Context, cid uint, emails []string, action string) error {
	var c CompanyRecruitmentCycle
	tx := db.WithContext(ctx).Where("id = ?", cid).First(&c)
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).
		Where("recruitment_cycle_id = ? AND (email IN ? OR roll_no IN ?)", c.RecruitmentCycleID, emails, emails).
		Updates(
			&StudentRecruitmentCycle{
				Type:     StudentRecruitmentCycleType(action),
				IsFrozen: true,
				Comment:  action + " by " + c.CompanyName,
			})
	return tx.Error
}

func FetchStudentRCIDs(ctx *gin.Context, rid uint, ids []string) ([]uint, []string, error) {
	var (
		students       []StudentRecruitmentCycle
		studentIDs     []uint
		filteredEmails []string
	)

	tx := db.WithContext(ctx).
		Where("recruitment_cycle_id = ? AND (email IN ? OR roll_no IN ?) AND is_frozen = ?", rid, ids, ids, false).
		Select("id", "email").Find(&students)

	for i := range students {
		studentIDs = append(studentIDs, students[i].ID)
		filteredEmails = append(filteredEmails, students[i].Email)
	}

	return studentIDs, filteredEmails, tx.Error
}

func FetchStudentRCID(ctx *gin.Context, rid uint, email string) (uint, error) {
	var student StudentRecruitmentCycle

	tx := db.WithContext(ctx).
		Where("recruitment_cycle_id = ? AND email = ? AND is_frozen = ?", rid, email, false).
		Select("id").First(&student)
	return student.ID, tx.Error
}

func FetchStudentEmailBySRCID(ctx *gin.Context, srcIDs []uint) ([]string, error) {
	var studentEmails []string

	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("id IN ?", srcIDs).Pluck("email", &studentEmails)
	return studentEmails, tx.Error
}

func getRegisteredStudentCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("recruitment_cycle_id = ?", rid).Count(&count)
	return int(count), tx.Error
}

type StatsBranchResponse struct {
	ProgramDepartmentID uint `json:"program_department_id"`
	Total               uint `json:"total"`
	PreOffer            uint `json:"pre_offer"`
	Recruited           uint `json:"recruited"`
}

func FetchRegisteredStudentCountByBranch(ctx *gin.Context, rid uint, res *[]StatsBranchResponse) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).
		Where("recruitment_cycle_id = ?", rid).Group("program_department_id").
		Select("program_department_id, count(*) as total, 0 as pre_offer, 0 as recruited").
		Scan(&res)
	return tx.Error
}

func GetStudentEligible(ctx *gin.Context, sid uint, eligibility string, cpiEligibility float64) (bool, error) {

	var primaryID int
	var secondaryID int

	var student StudentRecruitmentCycle

	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("id = ?", sid).First(&student)

	if tx.Error != nil {
		return false, tx.Error
	}

	primaryID = int(student.ProgramDepartmentID)
	secondaryID = int(student.SecondaryProgramDepartmentID)

	if !student.IsVerified {
		return false, errors.New("student not verified")
	}

	if student.CPI < cpiEligibility {
		return false, errors.New("cpi cutoff doesnot match")
	}

	if eligibility[primaryID] != '1' && eligibility[secondaryID] != '1' {
		return false, errors.New("student branch not eligible")
	}

	return true, nil
}

func FetchStudents(ctx *gin.Context, ids []uint, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id IN ?", ids).Find(students)
	return tx.Error
}

func deregisterAllStudentsWithRCID(ctx *gin.Context,rcid uint) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("recruitment_cycle_id = ? AND type = ?",rcid,AVAILABLE).Updates(StudentRecruitmentCycle{IsFrozen: true,Type: DEREGISTERED})
	return tx.Error
}