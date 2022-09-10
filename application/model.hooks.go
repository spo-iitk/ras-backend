package application

import (
	"strings"

	"gorm.io/gorm"
)

func (jp *Proforma) AfterUpdate(tx *gorm.DB) (err error) {
	if jp.IsApproved.Valid && jp.IsApproved.Bool {
		event := ProformaEvent{
			ProformaID:       jp.ID,
			Name:             string(Recruited),
			Duration:         "-",
			StartTime:        0,
			EndTime:          0,
			Sequence:         1000,
			RecordAttendance: false,
		}

		err = tx.Where("proforma_id = ? AND name = ?", event.ProformaID, event.Name).FirstOrCreate(&event).Error
		if err != nil {
			return
		}

		event = ProformaEvent{
			ProformaID:       jp.ID,
			Name:             string(ApplicationSubmitted),
			Duration:         "-",
			StartTime:        0,
			EndTime:          0,
			Sequence:         0,
			RecordAttendance: false,
		}

		err = tx.Where("proforma_id = ? AND name = ?", event.ProformaID, event.Name).FirstOrCreate(&event).Error
		if err != nil {
			return
		}

		if jp.Deadline > 0 {
			go insertCalenderApplicationDeadline(jp, &event)
		}
	}
	return
}

// Set default eligibility to none
func (p *Proforma) BeforeCreate(tx *gorm.DB) (err error) {
	p.Eligibility = strings.Repeat("0", 110)
	return
}

// Set default options of boolean to true,false
func (ques *ApplicationQuestion) BeforeCreate(tx *gorm.DB) (err error) {
	if ques.Type == BOOLEAN {
		ques.Options = "True,False"
	}
	return
}
func (ques *ApplicationQuestion) BeforeUpdate(tx *gorm.DB) (err error) {
	if ques.Type == BOOLEAN {
		ques.Options = "True,False"
	}
	return
}
