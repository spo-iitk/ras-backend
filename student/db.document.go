package student

// db.go - package student

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type StudentDocument struct {
    gorm.Model
    StudentID     uint   `json:"student_id"`
    DocumentType  string `json:"document_type"`
    DocumentPath  string `json:"document_path"`
    Verified      bool   `json:"verified"`
    ActionTakenBy string `json:"action_taken_by"`
}

func saveDocument(ctx *gin.Context, document *StudentDocument) error {
    tx := db.WithContext(ctx).Save(document)
    return tx.Error
}

func getDocumentsByStudentID(ctx *gin.Context, documents *[]StudentDocument, studentID uint) error {
    tx := db.WithContext(ctx).Where("student_id = ?", studentID).Find(documents)
    return tx.Error
}

func getDocumentByID(ctx *gin.Context, document *StudentDocument, docID uint) error {
    tx := db.WithContext(ctx).Where("id = ?", docID).First(document)
    return tx.Error
}

func getAllDocuments(ctx *gin.Context, documents *[]StudentDocument) error {
    tx := db.WithContext(ctx).Find(documents)
    return tx.Error
}

func getDocumentsByType(ctx *gin.Context, documents *[]StudentDocument, docType string) error {
    tx := db.WithContext(ctx).Where("document_type = ?", docType).Find(documents)
    return tx.Error
}
