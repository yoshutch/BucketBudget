package services

import (
	"log/slog"
	"time"

	"github.com/yoshutch/BucketBudget/internal/data"
	"github.com/yoshutch/BucketBudget/internal/models"
)

type BucketsService struct {
	Logger *slog.Logger
	Repo   *data.BucketsRepo
}

type BucketsServiceI interface {
	NewBucket(name string, balance models.Amount) (*models.Bucket, error)
	GetBucket(id uint64) (models.Bucket, error)
	UpdateBucket(id uint64, name string, balance models.Amount) (models.Bucket, error)
	DeleteBucket(id uint64) error
	GetMyBuckets() ([]models.Bucket, error)
}

func (b *BucketsService) NewBucket(name string, balance models.Amount) (*models.Bucket, error) {
	b.Logger.Info("New Bucket")
	// TODO authz
	// TODO validate input

	bucket, err := b.Repo.Insert(name, balance)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (b *BucketsService) GetBucket(id uint64) (models.Bucket, error) {
	b.Logger.Info("Get Bucket", "id", id)
	// TODO authz
	// TODO get from db
	return models.Bucket{
		ID:      1,
		Name:    "Test Buket",
		Balance: models.Amount{},
		Created: time.Now(),
	}, nil
}

func (b *BucketsService) UpdateBucket(id uint64, name string, balance models.Amount) (models.Bucket, error) {
	b.Logger.Info("Update Bucket", "id", id)
	// TODO validate input
	// TODO authz
	// TODO update database
	return models.Bucket{
		ID:      1,
		Name:    name,
		Balance: models.Amount{},
		Created: time.Now(),
	}, nil
}

func (b *BucketsService) DeleteBucket(id uint64) error {
	b.Logger.Info("Delete Bucket", "id", id)
	// TODO authz
	// TODO make sure bucket has a balance of zero?
	// TODO delete from db
	return nil
}

func (b *BucketsService) GetMyBuckets() ([]models.Bucket, error) {
	bucket, err := b.Repo.Get(1)
	if err != nil {
		return make([]models.Bucket, 0), err
	}
	return []models.Bucket{
		*bucket,
	}, nil
}
