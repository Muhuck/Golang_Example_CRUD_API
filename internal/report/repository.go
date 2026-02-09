package report

import "time"

type Repository interface {
	GetDailyReport(date time.Time) (*DailyReport, error)
	GetReport(startDate, endDate time.Time) (*DailyReport, error)
}
