package rc

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

type AllDocumentResponse struct {
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	Sid           uint         `json:"sid"`
	Did           uint         `json:"did"`
	Document      string       `json:"document"`
	Verified      sql.NullBool `json:"verified"`
	ActionTakenBy string       `json:"action_taken_by"`
	Type          DocumentType `json:"type"`
}

func getAllDocumentHandler(ctx *gin.Context) {
	// Extract recruitment cycle ID from URL parameters
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch all documents for the given recruitment cycle
	var documents []AllDocumentResponse
	err = fetchAllDocuments(ctx, rid, &documents)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of documents as JSON response
	ctx.JSON(http.StatusOK, documents)
}

func getAllDocumentHandlerByType(ctx *gin.Context) {
	// Extract recruitment cycle ID from URL parameters
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract document type from URL query parameters
	docTypeStr := ctx.Query("type")
	docType := DocumentType(docTypeStr)

	// Fetch documents by document type and recruitment cycle ID
	var documents []StudentDocument
	err = fetchAllDocumentsByType(ctx, rid, docType, &documents)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of documents as JSON response
	ctx.JSON(http.StatusOK, documents)
}

func getDocumentHandler(ctx *gin.Context) {
	sid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return   
	}

	var documents []StudentDocument
	err = fetchStudentDocuments(ctx, sid, &documents)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, documents)
}

type putDocumentVerifyRequest struct {
	Verified bool `json:"verified"`
}

func putDocumentVerifyHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		did, err := util.ParseUint(ctx.Param("docid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req putDocumentVerifyRequest

		err = ctx.BindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := middleware.GetUserID(ctx)


		ok, studentRCID, err := updateDocumentVerify(ctx, did, req.Verified, user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "document not found"})
			return
		}

		var student StudentRecruitmentCycle
		err = FetchStudent(ctx, studentRCID, &student)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("%v verified document with id %d, changed state to %v", user, did, req.Verified)

		msg := "Dear " + student.Name + "\n\n"
		msg += "Your document with id " + ctx.Param("did") + " has been "
		if req.Verified {
			msg += "ACCEPTED"
		} else {
			msg += "REJECTED"
		}
		mail_channel <- mail.GenerateMail(student.Email, "Action taken on document", msg)

		ctx.JSON(http.StatusOK, gin.H{"status": true})
	}
}
