package repository

import (
	"github.com/moura95/base-project-golang/internal/domain/entity"
)

type StorageRepository interface {
	Get() (response *entity.Storage, err error)
	Set(value entity.Storage) error
}
