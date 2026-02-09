package report

import "time"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetDailyReport() (*DailyReport, error) {
	// For production, pass time location properly. using Local for now.
	return s.repo.GetDailyReport(time.Now())
}

func (s *Service) GetReport(startDateStr, endDateStr string) (*DailyReport, error) {
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, err
	}

	// Adjust endDate to include the whole day
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

	return s.repo.GetReport(startDate, endDate)
}
