package services

import (
	"log/slog"
	"time"

	"yosbomb.com/bucketbudget/internal/models"
)

type BucketsService struct {
	logger *slog.Logger
}

type BucketsServiceI interface {
	NewBucket(input newBucketInput) (models.Bucket, error)
	GetBucket(id uint64) (models.Bucket, error)
	UpdateBucket(id uint64, input newBucketInput) (models.Bucket, error)
	DeleteBucket(id uint64) error
	GetMyBuckets() ([]models.Bucket, error)
}

func NewBucketsService(logger *slog.Logger) BucketsService {
	return BucketsService{
		logger: logger,
	}
}

type newBucketInput struct {
	name    string
	balance string
}

func (b *BucketsService) NewBucket(input newBucketInput) (models.Bucket, error) {
	b.logger.Info("New Bucket")
	// TODO authz
	// TODO validate input
	// TODO database insert
	return models.Bucket{
		ID:      1,
		Name:    input.name,
		Balance: input.balance,
		Created: time.Now(),
	}, nil
}

func (b *BucketsService) GetBucket(id uint64) (models.Bucket, error) {
	b.logger.Info("Get Bucket", id)
	// TODO authz
	// TODO get from db
	return models.Bucket{
		ID:      1,
		Name:    "Test Buket",
		Balance: "100.00",
		Created: time.Now(),
	}, nil
}

func (b *BucketsService) UpdateBucket(id uint64, input newBucketInput) (models.Bucket, error) {
	b.logger.Info("Update Bucket", "id", id)
	// TODO validate input
	// TODO authz
	// TODO update database
	return models.Bucket{
		ID:      1,
		Name:    input.name,
		Balance: input.balance,
		Created: time.Now(),
	}, nil
}

func (b *BucketsService) DeleteBucket(id uint64) error {
	b.logger.Info("Delete Bucket", "id", id)
	// TODO authz
	// TODO make sure bucket has a balance of zero?
	// TODO delete from db
	return nil
}

func (b *BucketsService) GetMyBuckets() ([]models.Bucket, error) {
	return []models.Bucket{
		{
			ID:      1,
			Name:    "Test",
			Balance: "100.00",
			Created: time.Time{},
		},
	}, nil
}
