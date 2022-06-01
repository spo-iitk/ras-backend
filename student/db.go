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

func updateStudent(ctx *gin.Context, student *Student, id uint) error {
	tx := db.WithContext(ctx).Model(&Student{}).Where("id = ?", id).Updates(student)
	return tx.Error
}

func getAllStudents(ctx *gin.Context, students *[]Student) error {
	tx := db.WithContext(ctx).Find(students)
	return tx.Error
}

