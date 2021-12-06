package daily_stats

import (
	"github.com/google/uuid"
)

type DailyStats struct {
	ID           uuid.UUID
	PVGeneration float32
	Load         float32
	Grid         float32
	Date         string
}
