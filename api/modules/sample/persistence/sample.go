package persistence

import (
	"time"

	"github.com/google/uuid"
	"github.com/rparaschak/mono-tmpl/api/pkg/database"
)

type Sample struct {
	Id          uuid.UUID            `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string               `gorm:"not null;index:idx_sample_name"`
	Geolocation database.Geolocation `gorm:"type:geography(Point,4326);not null"`
	CreatedAt   time.Time            `gorm:"default:now();index:idx_sample_created_at"`
	UpdatedAt   time.Time            `gorm:"default:now()"`
}

func (Sample) TableName() string {
	return "samples"
}
