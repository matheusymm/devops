package models

import (
	"time"

	"github.com/google/uuid"
)

type Mood struct {
	Id uuid.UUID `json:"id,omitempty"`
	UserId uuid.UUID `json:"user_id,omitempty"`
	Mood int `json:"mood,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}