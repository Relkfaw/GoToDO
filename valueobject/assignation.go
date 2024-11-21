package valueobject

import "github.com/google/uuid"

type Assignation struct {
	taskId uuid.UUID
	userId uuid.UUID
}
