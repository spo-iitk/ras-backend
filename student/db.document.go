package student

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)


func saveDocument(ctx *gin.Context, document *StudentDocument) error {
	tx := db.WithContext(ctx).Save(document)
	return tx.Error
}

func getDocumentsByStudentID(ctx *gin.Context, documents *[]StudentDocument, studentID uint) error {
	tx := db.WithContext(ctx).Where("student_id = ?", studentID).Find(documents)
	return tx.Error
}

func getDocumentByID(ctx *gin.Context, document *StudentDocument, docID uint) error {
	tx := db.WithContext(ctx).First(document, docID)
	return tx.Error
}

func getAllDocuments(ctx *gin.Context, documents *[]StudentDocument) error {
	tx := db.WithContext(ctx).Find(documents)
	return tx.Error
}

func getDocumentsByType(ctx *gin.Context, documents *[]StudentDocument, docType string) error {
	tx := db.WithContext(ctx).Where("type = ?", docType).Find(documents)
	return tx.Error
}

func updateDocumentVerify(ctx *gin.Context, docid uint, verified bool, user string) (bool, error){
	var document StudentDocument
	tx := db.WithContext(ctx).Model(&document).Clauses(clause.Returning{}).Where("id = ?", docid).Updates(map[string]interface{}{"verified": verified, "action_taken_by": user})
	return tx.RowsAffected == 1, tx.Error
}