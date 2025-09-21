package rc

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/student"
	"gorm.io/gorm/clause"
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

func updateStudentBulk(ctx *gin.Context, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"program_department_id", "secondary_program_department_id", "cpi"}),
	}).Create(&students)
	return tx.Error
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
		Where("recruitment_cycle_id = ? AND program_department_id IS NOT NULL AND program_department_id != 0", rid).Group("program_department_id").
		Select("program_department_id, count(*) as total, 0 as pre_offer, 0 as recruited").
		Scan(&res)
	return tx.Error
}

func FetchRegisteredStudentCountBySecondaryBranch(ctx *gin.Context, rid uint, res *[]StatsBranchResponse) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).
		Where("recruitment_cycle_id = ? AND secondary_program_department_id IS NOT NULL AND secondary_program_department_id != 0", rid).Group("secondary_program_department_id").
		Select("secondary_program_department_id as program_department_id, count(*) as total, 0 as pre_offer, 0 as recruited").
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

func syncStudentDataRC(ctx *gin.Context, rid uint) error {
	var rcStudents []StudentRecruitmentCycle
	err := fetchAllStudents(ctx, rid, &rcStudents)
	if err != nil {
		return err
	}
	var emailIds []string
	for _, student := range rcStudents {
		emailIds = append(emailIds, student.Email)
	}
	var masterStudents []student.Student
	err = student.FetchStudents(ctx, &masterStudents, emailIds)
	if err != nil {
		return err
	}
	var masterStudentMap = make(map[string]student.Student)
	for _, masterStudent := range masterStudents {
		masterStudentMap[masterStudent.IITKEmail] = masterStudent
	}

	for idx := range rcStudents {
		var masterStudent, exists = masterStudentMap[rcStudents[idx].Email]
		if !exists {
			continue
		}
		rcStudents[idx].ProgramDepartmentID = masterStudent.ProgramDepartmentID
		rcStudents[idx].SecondaryProgramDepartmentID = masterStudent.SecondaryProgramDepartmentID
		rcStudents[idx].CPI = masterStudent.CurrentCPI
	}
	err = updateStudentBulk(ctx, &rcStudents)
	if err != nil {
		return err
	}

	return nil
}

// func deregisterAllStudentsWithRCID(ctx *gin.Context,rcid uint) error {
// 	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("recruitment_cycle_id = ? AND type = ?",rcid,AVAILABLE).Updates(StudentRecruitmentCycle{IsFrozen: true,Type: DEREGISTERED})
// 	return tx.Error
// }

func UnRecruitStudent(ctx *gin.Context, sid uint, rid uint) error {
	var stu StudentRecruitmentCycle
	tx := db.WithContext(ctx).Where("id = ?", sid).First(&stu)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).
		Where("recruitment_cycle_id = ? AND  id = ?", rid, sid).
		Updates(map[string]interface{}{
			"type":      AVAILABLE,
			"is_frozen": false,
			"comment":   " ",
		})
	return tx.Error
}

func UnRecruitAll(ctx *gin.Context, sids []uint) error {
	tx := db.WithContext(ctx).Where("id IN (?)", sids).Model(&StudentRecruitmentCycle{}).
		Updates(map[string]interface{}{
			"type":      AVAILABLE,
			"is_frozen": false,
			"comment":   " ",
		})
	return tx.Error
}
