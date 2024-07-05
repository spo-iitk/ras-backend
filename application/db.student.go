package application

import "github.com/gin-gonic/gin"

func createEventStudents(ctx *gin.Context, eventStudents *[]EventStudent) error {
	tx := db.WithContext(ctx).Create(eventStudents)
	return tx.Error
}

func getCompanyRecruitmentStats(ctx *gin.Context, cid uint, stats *[]statsResponse) error {
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.name IN ? AND proforma_events.id = event_students.proforma_event_id", []EventType{Recruited, PIOPPOACCEPTED}).
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id AND proformas.company_recruitment_cycle_id = ?", cid).
		Select("event_students.student_recruitment_cycle_id, proformas.company_name, proformas.profile ,proforma_events.name as type").
		Order("event_students.student_recruitment_cycle_id").
		Find(stats)
	return tx.Error
}
func fetchCompanyRecruitCount(ctx *gin.Context, cids []uint) (map[uint]int, error) {
	resultCounts := make(map[uint]int)

	var stats []companyRecruitResponce
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.name IN ? AND proforma_events.id = event_students.proforma_event_id", []EventType{Recruited, PIOPPOACCEPTED}).
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id").
		Where("proformas.company_recruitment_cycle_id IN ?", cids).
		Select("proformas.company_recruitment_cycle_id, COUNT(*) as count").
		Group("proformas.company_recruitment_cycle_id").
		Order("proformas.company_recruitment_cycle_id").
		Find(&stats)

	if tx.Error != nil {
		return nil, tx.Error
	}

	// Populate resultCounts map
	for _, stat := range stats {
		resultCounts[stat.CompanyRecruitmentCycleID] = stat.Count
	}

	return resultCounts, nil
}

func fetchCompanyPPOCount(ctx *gin.Context, cids []uint) (map[uint]int, error) {
	resultCounts := make(map[uint]int)

	var stats []companyRecruitResponce
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.name IN ? AND proforma_events.id = event_students.proforma_event_id", []EventType{PIOPPOACCEPTED}).
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id").
		Where("proformas.company_recruitment_cycle_id IN ?", cids).
		Select("proformas.company_recruitment_cycle_id, COUNT(*) as count").
		Group("proformas.company_recruitment_cycle_id").
		Order("proformas.company_recruitment_cycle_id").
		Find(&stats)

	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, stat := range stats {
		resultCounts[stat.CompanyRecruitmentCycleID] = stat.Count
	}

	return resultCounts, nil
}

func getRecruitmentStats(ctx *gin.Context, rid uint, stats *[]statsResponse) error {
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.name IN ? AND proforma_events.id = event_students.proforma_event_id", []EventType{Recruited, PIOPPOACCEPTED}).
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id AND proformas.recruitment_cycle_id = ?", rid).
		Select("event_students.student_recruitment_cycle_id, proformas.company_name, proformas.profile ,proforma_events.name as type").
		Order("event_students.student_recruitment_cycle_id").
		Find(stats)
	return tx.Error
}

func fetchStudentRCIDByEvents(ctx *gin.Context, eventID uint) ([]uint, error) {
	var ids []uint
	tx := db.WithContext(ctx).Model(&EventStudent{}).Where("proforma_event_id = ?", eventID).Pluck("student_recruitment_cycle_id", &ids)
	return ids, tx.Error
}

func fetchStudentsByEvent(ctx *gin.Context, eventID uint, students *[]EventStudent) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ?", eventID).Find(students)
	return tx.Error
}

func getCurrentApplicationCount(ctx *gin.Context, sid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Where("student_recruitment_cycle_id = ?", sid).
		Group("company_recruitment_cycle_id").Count(&count)
	return int(count), tx.Error
}

func deleteStudentFromEvent(ctx *gin.Context, eventID, studentID uint) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ? AND student_recruitment_cycle_id = ?", eventID, studentID).Delete(&EventStudent{})
	return tx.Error
}

func deleteAllStudentsFromEvent(ctx *gin.Context, eventID uint) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ?", eventID).Delete(&EventStudent{})
	return tx.Error
}

func getStudentIDByEventID(ctx *gin.Context, eid uint) ([]uint, error) {
	var studentsIDs []uint
	tx := db.WithContext(ctx).Where("proforma_event_id = ?", eid).Model(&EventStudent{}).
		Pluck("student_recruitment_cycle_id", &studentsIDs)
	return studentsIDs, tx.Error
}
