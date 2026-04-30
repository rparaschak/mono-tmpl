package persistence

import (
	"time"

	"github.com/google/uuid"
)

type Sample struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string    `gorm:"not null;index:idx_sample_name"`
	CreatedAt time.Time `gorm:"default:now();index:idx_sample_created_at"`
	UpdatedAt time.Time `gorm:"default:now()"`
}
