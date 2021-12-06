package reading

import (
	"github.com/google/uuid"
)

type Reading struct {
	ID           uuid.UUID
	PVGeneration int64
	Load         int64
	Grid         int64
	Timestamp    int64
}
