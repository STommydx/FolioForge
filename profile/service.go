package profile

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{
		db: db,
	}
}

func (ps *ProfileService) GetProfileByName(ctx context.Context, profileName string) (*Profile, error) {
	var profile Profile
	result := ps.db.WithContext(ctx).First(&profile, "profile_name = ?", profileName)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &profile, nil
}
