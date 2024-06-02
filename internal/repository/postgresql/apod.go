package postgresql

import (
	"context"
	"errors"

	"NasaEnjoyer/domain"

	"gorm.io/gorm"
)

type APODRepository struct {
	db *gorm.DB
}

func NewAPODRepository(db *gorm.DB) *APODRepository {
	return &APODRepository{db: db}
}
func (r *APODRepository) AddAPOD(ctx context.Context, apod *domain.APOD) error {
	var existing domain.APOD
	err := r.db.Set("gorm:context", ctx).Where("date = ?", apod.Date).Take(&existing).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Если запись не найдена, создаем новую
			return r.db.Set("gorm:context", ctx).Create(apod).Error
		}
		return err
	}

	existing = *apod
	return r.db.Set("gorm:context", ctx).Model(&existing).Where("date = ?", apod.Date).Updates(&existing).Error
}

func (r *APODRepository) GetAPOD(ctx context.Context, date string) (*domain.APOD, error) {
	var apod domain.APOD
	err := r.db.Set("gorm:context", ctx).Where("date = ?", date).First(&apod).Error
	if err != nil {
		return nil, err
	}

	return &apod, nil
}

func (r *APODRepository) GetAPODs(ctx context.Context) ([]domain.APOD, error) {
	var apods []domain.APOD
	err := r.db.Set("gorm:context", ctx).Find(&apods).Error
	if err != nil {
		return nil, err
	}
	return apods, nil
}
