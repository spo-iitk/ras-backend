package ras

import (
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/spo-iitk/ras-backend/constants"
)

type Program struct {
	ID   c.ProgramID
	Name string
}

type Department struct {
	ID   c.DepartmentID
	Name string
}

func getPrograms(ctx *gin.Context) {
	var programs [10]Program
	var i c.ProgramID
	for i = 0; i < 10; i++ {
		programs[i].ID = i
		programs[i].Name = c.GetProgram(i)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": programs})
}

func getDepartments(ctx *gin.Context) {
	var departments [26]Department
	var i c.DepartmentID
	for i = 0; i < 26; i++ {
		departments[i].ID = i
		departments[i].Name = c.GetDepartment(i)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": departments})
}

func getProgramsDepartments(ctx *gin.Context) {
	var programDepartments [10]c.ProgramDepartment
	var i uint
	for i = 0; i < 10; i++ {
		programDepartments[i] = c.GetProgramDepartment(i)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": programDepartments})
}
