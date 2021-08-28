package model

import (
	"github.com/google/uuid"
)

type Job struct {
	ID            uuid.UUID `json:"id"`
	ApplicationID uuid.UUID `json:"application_id"`
	Status        string    `json:"status"`
}
