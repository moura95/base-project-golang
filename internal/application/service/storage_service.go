package service

import (
	"fmt"
	"time"

	"github.com/moura95/base-project-golang/config"
	"github.com/moura95/base-project-golang/internal/domain/entity"
	"github.com/moura95/base-project-golang/internal/infrastructure/repository"
	"go.uber.org/zap"
)

type StorageService struct {
	repository repository.StorageRepository
	config     config.Config
	logger     *zap.SugaredLogger
}

func NewStorageService(repo repository.StorageRepository, cfg config.Config, log *zap.SugaredLogger) *StorageService {
	return &StorageService{
		repository: repo,
		config:     cfg,
		logger:     log,
	}
}

func (s *StorageService) Set(value int32) (string, error) {
	storage := entity.Storage{
		Value: value,
	}
	err := storage.Validate()
	if err != nil {
		s.logger.Error("Failed to validate storage value: ", err)
		return "", fmt.Errorf("failed to validate storage value: %w", err)
	}

	if err != nil {
		s.logger.Error("Failed to create interact: ", err)
		return "", fmt.Errorf("failed to create interact: %w", err)
	}

	return "", nil
}

func (s *StorageService) Get() (*entity.Storage, error) {

	intValue := 1

	storage := entity.ToEntity(int32(intValue), time.Now())
	return storage, nil
}

func (s *StorageService) Check() (bool, error) {

	return true, nil
}

func (s *StorageService) Sync() (bool, error) {
	return true, nil
}
