package db

import (
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/oklog/ulid/v2"
)

type Profile struct {
	ID          ulid.ULID `gorm:"primaryKey"`
	ProfileName string    `gorm:"uniqueIndex"`
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
