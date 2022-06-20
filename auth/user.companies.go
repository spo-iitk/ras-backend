package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func companiesAddedHandler(ctx *gin.Context) {
	companies, err := getAllCompaniesAdded(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var company_names []string
	for _, company := range companies {
		company_names = append(company_names, company.Name)
	}

	ctx.JSON(http.StatusOK, gin.H{"companies": company_names})
}
