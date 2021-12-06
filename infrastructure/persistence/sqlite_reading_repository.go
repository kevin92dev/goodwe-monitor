package persistence

import (
	"github.com/jinzhu/gorm"
	"goodwe_monitor/domain/reading"
	"goodwe_monitor/domain/repository"
)

type ReadingRepositoryImpl struct {
	Conn *gorm.DB
}

func NewReadingRepositoryWithRDB(conn *gorm.DB) repository.ReadingRepository {
	return &ReadingRepositoryImpl{Conn: conn}
}

func (r *ReadingRepositoryImpl) Save(reading *reading.Reading) error {
	if err := r.Conn.Save(&reading).Error; err != nil {
		return err
	}

	return nil
}
