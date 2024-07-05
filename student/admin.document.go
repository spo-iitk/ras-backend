package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getDocumentHandler(ctx *gin.Context) {
	sid, err := strconv.ParseUint(ctx.Param("sid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var documents []StudentDocument
	err = getDocumentsByStudentID(ctx, &documents, uint(sid))
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

        var document StudentDocument
        err = getDocumentByID(ctx, &document, uint(did))
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        ok, err := updateDocumentVerify(ctx, did, req.Verified, user)
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
        }

        if !ok {
            ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not verify document"})
			return
        }

        var student Student
        err = getStudentByID(ctx, &student, document.StudentID)
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        logrus.Infof("%v verified document with id %d, changed state to %v", user, did, req.Verified)

        // Constructing the email message
        msg := "Dear " + student.Name + "\n\n"
        msg += "Your document (" + document.Type + ") with id " + strconv.Itoa(int(did)) + " has been "
        if req.Verified {
            msg += "ACCEPTED."
        } else {
            msg += "REJECTED."
        }
        msg += "\n\nBest regards,\nYour Verification Team"

        mail_channel <- mail.GenerateMail(student.PersonalEmail, "Action taken on document", msg)

        ctx.JSON(http.StatusOK, gin.H{"status": true})
    }
}


func getAllDocumentHandler(ctx *gin.Context) {
	var documents []StudentDocument
	err := getAllDocuments(ctx, &documents)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, documents)
}

func getAllDocumentHandlerByType(ctx *gin.Context) {
	docType := ctx.Param("type")
	var documents []StudentDocument
	err := getDocumentsByType(ctx, &documents, docType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, documents)
}