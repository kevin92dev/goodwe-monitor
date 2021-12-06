package repository

import "goodwe_monitor/domain/reading"

type ReadingRepository interface {
	Save(*reading.Reading) error
	/*Get(id int) (*domain.News, error)
	GetAll() ([]domain.News, error)
	GetBySlug(slug string) ([]*domain.News, error)
	GetAllByStatus(status string) ([]domain.News, error)
	Remove(id int) error
	Update(*domain.News) error*/
}
