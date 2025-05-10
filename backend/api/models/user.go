package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Password []byte `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}