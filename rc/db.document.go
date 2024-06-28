package rc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func addStudentDocument(ctx *gin.Context, document string, sid uint, rid uint, docType DocumentType) error {
	tx := db.WithContext(ctx).Model(&StudentDocument{}).Create(&StudentDocument{
		StudentRecruitmentCycleID: sid,
		Document:                  document,
		RecruitmentCycleID:        rid,
		DocumentType:              docType,
	})
	return tx.Error
}

func fetchStudentDocuments(ctx *gin.Context, sid uint, documents *[]StudentDocument) error {
	tx := db.WithContext(ctx).Model(&StudentDocument{}).Where("student_recruitment_cycle_id = ?", sid).Find(documents)
	return tx.Error
}

func fetchAllDocuments(ctx *gin.Context, rid uint, documents *[]AllDocumentResponse) error {
	tx := db.WithContext(ctx).Model(&StudentDocument{}).
		Joins("JOIN student_recruitment_cycles ON student_recruitment_cycles.id = student_documents.student_recruitment_cycle_id AND student_documents.recruitment_cycle_id = ?", rid).
		Select("student_recruitment_cycles.name as name, student_recruitment_cycles.email as email, student_recruitment_cycles.id as sid, student_documents.id as dsid, student_recruitment_cycles.roll_no as roll_no, student_documents.document as document, student_documents.verified as verified, student_documents.action_taken_by as action_taken_by, student_documents.document_type as document_type").
		Scan(documents)
	return tx.Error
}

// fetchAllDocumentsByType fetches all documents of a specific type for a given recruitment cycle.
func fetchAllDocumentsByType(ctx *gin.Context, rid uint, docType DocumentType, documents *[]StudentDocument) error {
	// Ensure the type is valid and matches your constants or enums

	// Query documents from the database based on recruitment cycle ID and document type
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND document_type = ?", rid, docType).Find(documents)

	// Handle any errors that occur during the query
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func updateDocumentVerify(ctx *gin.Context, did uint, verified bool, user string) (bool, uint, error) {
	var document StudentDocument
	tx := db.WithContext(ctx).Model(&document).Clauses(clause.Returning{}).
		Where("id = ? ", did).
		Update("verified", verified).
		Update("action_taken_by", user)
	return tx.RowsAffected == 1, document.StudentRecruitmentCycleID, tx.Error
}
