package healthz

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HealthzResult struct {
	Status  string               `json:"status" example:"ok" enums:"ok,error"`
	Uptime  string               `json:"uptime"`
	Details HealthzResultDetails `json:"details"`
}

type HealthzResultDetails struct {
	Db HealthzResultDetail `json:"db"`
}

type HealthzResultDetail struct {
	Status string `json:"status" example:"ok" enums:"ok,error"`
}

type HealthzService struct {
	db               *gorm.DB
	startupTimestamp time.Time
}

func NewHealthzService(db *gorm.DB) *HealthzService {
	return &HealthzService{
		db:               db,
		startupTimestamp: time.Now(),
	}
}

func (hs *HealthzService) HealthCheck(ctx context.Context) (*HealthzResult, error) {
	db, err := hs.db.DB()
	if err != nil {
		return nil, fmt.Errorf("fail to get underlying db: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return &HealthzResult{
			Status: "error",
			Details: HealthzResultDetails{
				Db: HealthzResultDetail{
					Status: "error",
				},
			},
		}, nil
	}
	return &HealthzResult{
		Status: "ok",
		Uptime: time.Since(hs.startupTimestamp).String(),
		Details: HealthzResultDetails{
			Db: HealthzResultDetail{
				Status: "ok",
			},
		},
	}, nil
}
