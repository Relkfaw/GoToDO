package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID    uuid.UUID
	begin time.Time
	end   time.Time
}
