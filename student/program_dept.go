package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Program struct {
	ID   ProgramID
	Name string
}

type Department struct {
	ID   DepartmentID
	Name string
}

func getPrograms(ctx *gin.Context) {
	var programs [10]Program
	var i ProgramID
	for i = 0; i < 10; i++ {
		programs[i].ID = i
		programs[i].Name = GetProgram(i)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": programs})
}

func getDepartments(ctx *gin.Context) {
	var departments [26]Department
	var i DepartmentID
	for i = 0; i < 26; i++ {
		departments[i].ID = i
		departments[i].Name = GetDepartment(i)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": departments})
}

func getProgramsDepartments(ctx *gin.Context) {
	var programDepartments [10]ProgramDepartment
	var i uint
	for i = 0; i < 10; i++ {
		programDepartments[i] = GetProgramDepartment(i)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": programDepartments})
}
