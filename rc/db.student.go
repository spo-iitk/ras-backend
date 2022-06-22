package rc

import "github.com/gin-gonic/gin"

func fetchAllStudents(ctx *gin.Context, rid string, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(students)
	return tx.Error
}

func fetchStudent(ctx *gin.Context, email string, rid string, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("email = ? AND recruitment_cycle_id = ?", email, rid).First(student)
	return tx.Error
}

func fetchStudentByID(ctx *gin.Context, sid uint, rid string, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id = ? AND recruitment_cycle_id = ?", sid, rid).First(student)
	return tx.Error
}

func updateStudent(ctx *gin.Context, student *StudentRecruitmentCycle) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", student.ID).Updates(student)
	return tx.RowsAffected > 0, tx.Error
}

func freezeStudentsToggle(ctx *gin.Context, emails []string, frozen bool) (bool, error) {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("email IN ?", emails).Updates(
		&StudentRecruitmentCycle{
			IsFrozen: frozen,
		})
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

func UpdateStudentType(ctx *gin.Context, cid uint, emails []string) error {
	var c CompanyRecruitmentCycle
	tx := db.WithContext(ctx).Where("id = ?", cid).First(&c)
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("recruitment_cycle_id = ? AND email IN ?", c.RecruitmentCycleID, emails).Updates(
		&StudentRecruitmentCycle{
			Type:     PIOPPO,
			IsFrozen: true,
			Comment:  "PIO/PPO by " + c.CompanyName,
		})
	return tx.Error
}

func FetchStudentRCIDs(ctx *gin.Context, rid uint, emails []string) ([]uint, error) {
	var students []StudentRecruitmentCycle
	var studentIDs []uint

	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND email IN ?", rid, emails).Select("id").Find(&students).Pluck("id", &studentIDs)
	return studentIDs, tx.Error
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

func GetStudentEligible(ctx *gin.Context, eligibility string) (bool, error) {
	var primaryID int
	var secondaryID int

	var student StudentRecruitmentCycle
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).First(&student)

	if tx.Error != nil {
		return false, tx.Error
	}

	primaryID = int(student.ProgramDepartmentID)
	secondaryID = int(student.SecondaryProgramDepartmentID)

	if eligibility[primaryID] == '1' || eligibility[secondaryID] == '1' {
		return true, nil
	}

	return false, nil

}
