package application

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/calendar/v3"
)

func deleteCalenderEvent(cID string, cevent *ProformaEvent) {
	err := cal_srv.Events.Delete(cID, cevent.CalID).Do()
	if err != nil {
		logrus.Errorf("Unable to create event. %v", err)
	}
}

func insertCalenderEvent(event *ProformaEvent, proforma *Proforma, loc *time.Location, time_zone string, cID string) {
	cevent := &calendar.Event{
		Summary:  fmt.Sprintf("%s - %s, %s", event.Name, proforma.Profile, proforma.CompanyName),
		Location: event.Venue,
		Description: fmt.Sprintf(
			"%s of profile %s - %s has been scheduled from %s to %s\nhttps://placement.iitk.ac.in/student/rc/%d/events/%d",
			event.Name, proforma.Profile, proforma.CompanyName,
			time.UnixMilli(int64(event.StartTime)).In(loc).Format("2006-01-02 15:04"),
			time.UnixMilli(int64(event.EndTime)).In(loc).Format("2006-01-02 15:04"),
			proforma.RecruitmentCycleID, event.ID),
		Start: &calendar.EventDateTime{
			DateTime: time.UnixMilli(int64(event.StartTime)).In(loc).Format(time.RFC3339),
			TimeZone: time_zone,
		},
		End: &calendar.EventDateTime{
			DateTime: time.UnixMilli(int64(event.EndTime)).In(loc).Format(time.RFC3339),
			TimeZone: time_zone,
		},
	}

	if event.CalID == "" {
		cevent, err := cal_srv.Events.Insert(cID, cevent).Do()
		if err != nil {
			logrus.Errorf("Unable to create event. %v", err)
		}

		event.CalID = cevent.Id
		updateEvent(nil, event)
	}

	_, err := cal_srv.Events.Update(cID, event.CalID, cevent).Do()
	if err != nil {
		logrus.Errorf("Unable to update event. %v", err)
	}
}
