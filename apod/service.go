package apod

import (
	"NasaEnjoyer/domain"
	"NasaEnjoyer/pkg/nasa"
	"context"
	"errors"
)

type ApodRepository interface {
	AddAPOD(ctx context.Context, apod *domain.APOD) error
	GetAPOD(ctx context.Context, date string) (*domain.APOD, error)
	GetAPODs(ctx context.Context) ([]domain.APOD, error)
}

type ImageStore interface {
	SaveImage(url, saveDir string) (string, error)
}

type Service struct {
	apodRepo     ApodRepository
	nasaCli      *nasa.NASAClient
	imageStore   ImageStore
	imageSaveDir string
}

func NewService(apodRepo ApodRepository, nasaCli *nasa.NASAClient, imageStore ImageStore, imageSaveDir string) *Service {
	return &Service{
		apodRepo:     apodRepo,
		nasaCli:      nasaCli,
		imageStore:   imageStore,
		imageSaveDir: imageSaveDir,
	}
}

func (s *Service) FetchAndStoreAPOD(ctx context.Context) error {
	apod, err := s.nasaCli.GetAPOD(ctx)
	if err != nil {
		return err
	}

	imagePath, err := s.imageStore.SaveImage(apod.URL, s.imageSaveDir)
	if err != nil {
		if errors.Is(err, domain.ErrNotAnImage) {
			// Если контент не является изображением
			apod.ImagePath = "NOT_AN_IMAGE"
		} else {
			return err
		}
	} else {
		apod.ImagePath = imagePath
	}

	err = s.apodRepo.AddAPOD(ctx, apod)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAPOD(ctx context.Context, date string) (*domain.APOD, error) {
	apod, err := s.apodRepo.GetAPOD(ctx, date)
	if err != nil {
		return nil, err
	}

	return apod, nil
}

func (s *Service) GetAPODs(ctx context.Context) ([]domain.APOD, error) {
	apods, err := s.apodRepo.GetAPODs(ctx)
	if err != nil {
		return nil, err
	}

	return apods, nil
}
