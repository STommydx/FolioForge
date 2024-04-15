package profile

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Profile struct {
	ID          ulid.ULID `json:"id"`
	ProfileName string    `json:"profile_name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromModel(p *Profile) *Profile {
	return &Profile{
		ID:          p.ID,
		ProfileName: p.ProfileName,
		FirstName:   p.FirstName,
		LastName:    p.LastName,
		Email:       p.Email,
		Phone:       p.Phone,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
