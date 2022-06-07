package student

import "github.com/gin-gonic/gin"

func createStudent(ctx *gin.Context, student *Student) error {
	tx := db.WithContext(ctx).Create(student)
	return tx.Error
}

func getStudent(ctx *gin.Context, student *Student, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(student)
	return tx.Error
}

func FetchStudents(ctx *gin.Context, students *[]Student, emails []string) error {
	tx := db.WithContext(ctx).Where("email IN ?", emails).Find(students)
	return tx.Error
}

func updateStudent(ctx *gin.Context, student *Student, id uint) (bool, error) {
	tx := db.WithContext(ctx).Model(&Student{}).Where("id = ?", id).Updates(student)
	return tx.RowsAffected > 0, tx.Error
}

func getAllStudents(ctx *gin.Context, students *[]Student) error {
	tx := db.WithContext(ctx).Find(students)
	return tx.Error
}
