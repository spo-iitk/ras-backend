package student

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func FirstOrCreateStudent(ctx *gin.Context, student *Student) error {
	tx := db.WithContext(ctx).Where("iitk_email = ?", student.IITKEmail).FirstOrCreate(student)
	return tx.Error
}

func getStudentByID(ctx *gin.Context, student *Student, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(student)
	return tx.Error
}

func FetchStudentsByID(ctx *gin.Context, id []uint, students *[]Student) error {
	tx := db.WithContext(ctx).Where("id IN ?", id).Find(students).Order("id ASC")
	return tx.Error
}

func getStudentByEmail(ctx *gin.Context, student *Student, email string) error {
	tx := db.WithContext(ctx).Where("iitk_email =?", email).First(student)
	return tx.Error
}

func FetchStudents(ctx *gin.Context, students *[]Student, ids []string) error {
	tx := db.WithContext(ctx).Where("(iitk_email IN ? OR roll_no IN ?) AND is_verified = ? AND is_editable = ?", ids, ids, true,false).Find(students)
	return tx.Error
}

func getAllStudents(ctx *gin.Context, students *[]Student) error {
	tx := db.WithContext(ctx).Find(students)
	return tx.Error
}

func getLimitedStudents(ctx *gin.Context, students *[]Student, lastFetchedId uint, pageSize uint, batch uint) error {
	tx := db.WithContext(ctx).Order("id asc").Where("id >= ? AND roll_no LIKE ?", lastFetchedId, strconv.Itoa(int(batch))+"%").Limit(int(pageSize)).Find(students)
	return tx.Error
}

func updateStudentByID(ctx *gin.Context, student *Student) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", student.ID).Select("*").Omit("id", "created_at", "updated_at", "deleted_at", "iitk_email", "roll_no", "is_verified", "is_editable", "program_department_id", "secondary_program_department_id").Updates(Student{Name: student.Name, Specialization: student.Specialization, Preference: student.Preference, Gender: student.Gender, Disablity: student.Disablity, DOB: student.DOB, ExpectedGraduationYear: student.ExpectedGraduationYear, PersonalEmail: student.PersonalEmail, Phone: student.Phone, AlternatePhone: student.AlternatePhone, WhatsappNumber: student.WhatsappNumber, CurrentCPI: student.CurrentCPI, UGCPI: student.UGCPI, TenthBoard: student.TenthBoard, TenthYear: student.TenthYear, TenthMarks: student.TenthMarks, TwelfthBoard: student.TwelfthBoard, TwelfthYear: student.TwelfthYear, TwelfthMarks: student.TwelfthMarks, EntranceExam: student.EntranceExam, EntranceExamRank: student.EntranceExamRank, Category: student.Category, CategoryRank: student.CategoryRank, CurrentAddress: student.CurrentAddress, PermanentAddress: student.PermanentAddress, FriendName: student.FriendName, FriendPhone: student.FriendPhone, IsVerified: student.IsVerified, IsEditable: student.IsEditable})
	return tx.RowsAffected > 0, tx.Error
}

func verifyStudent(ctx *gin.Context, student *Student) (bool, error) {
	tx := db.WithContext(ctx).Model(&student).
		Clauses(clause.Returning{}).
		Where("id = ?", student.ID).
		Updates(map[string]interface{}{"is_verified": student.IsVerified, "is_editable": !student.IsVerified})
	return tx.RowsAffected > 0, tx.Error
}

func updateStudentByEmail(ctx *gin.Context, student *Student, email string) (bool, error) {
	var studentDetails Student
	err := getStudentByEmail(ctx, &studentDetails, student.IITKEmail)
	if err != nil {
		return false, err
	}
	verified := studentDetails.IsVerified
	var tx *gorm.DB
	if !verified {
		tx = db.WithContext(ctx).Model(&Student{}).
			Where("iitk_email = ? AND is_editable = ?", email, true).
			Updates(student)
	} else {
		tx = db.WithContext(ctx).Model(&Student{}).
			Where("iitk_email = ? AND is_editable = ?", email, true).
			Updates(&Student{Specialization: student.Specialization,Preference: student.Preference, ExpectedGraduationYear: student.ExpectedGraduationYear,PersonalEmail: student.PersonalEmail,Phone: student.Phone,AlternatePhone: student.AlternatePhone,WhatsappNumber: student.WhatsappNumber,CurrentAddress: student.CurrentAddress,PermanentAddress: student.PermanentAddress,FriendName: student.FriendName,FriendPhone: student.FriendPhone})
	}
	if tx.Error != nil {
		return false, tx.Error
	}

	// tx = db.WithContext(ctx).Model(&Student{}).
	// 	Where("iitk_email = ? AND is_editable = ?", email, true).Update("is_verified", false)
	return tx.RowsAffected > 0, tx.Error
}

func deleteStudent(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&Student{})
	return tx.Error
}

// func UpdateIsEditableWithIDs(ctx *gin.Context, ids []uint, editable bool) error {
// 	tx := db.WithContext(ctx).Model(&Student{}).Where("id in ?", ids).Update("is_editable", editable)
// 	return tx.Error
// }

func updateIsEditableWithID(ctx *gin.Context, id uint, editable bool) error {
	tx := db.WithContext(ctx).Model(&Student{}).Where("id = ?", id).Update("is_editable", editable)
	return tx.Error
}
