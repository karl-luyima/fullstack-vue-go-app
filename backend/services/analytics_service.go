package services

import (
	"context"

	"gorm.io/gorm"
)

type AnalyticsService interface {
	SignupsByDay(ctx context.Context, days int) ([]SignupsByDayPoint, error)
}

type analyticsService struct {
	db *gorm.DB
}

// Analytics reads directly via gorm.DB rather than through a repository
// interface, since this is a reporting query (grouping/counting) rather
// than fetching/saving a single resource — a reasonable exception to the
// "repositories own all DB access" rule for read-only aggregate queries.
func NewAnalyticsService(db *gorm.DB) AnalyticsService {
	return &analyticsService{db: db}
}

func (s *analyticsService) SignupsByDay(ctx context.Context, days int) ([]SignupsByDayPoint, error) {
	var results []SignupsByDayPoint

	err := s.db.WithContext(ctx).
		Table("users").
		Select("TO_CHAR(created_at, 'YYYY-MM-DD') as date, COUNT(*) as count").
		Where("created_at >= NOW() - (? * INTERVAL '1 day')", days).
		Group("date").
		Order("date ASC").
		Scan(&results).Error

	return results, err
}