package data

import (
	"database/sql"

	"yosbomb.com/bucketbudget/internal/models"
)

type BucketsRepo struct {
	DB *sql.DB
}

func (r *BucketsRepo) Insert(name string, balance models.Amount) (*models.Bucket, error) {
	query := `INSERT INTO buckets (name, balance)
		VALUES($1, $2)
		RETURNING id, name, balance, created_at`

	var bucket models.Bucket
	err := r.DB.QueryRow(query, name, balance.Cents).
		Scan(&bucket.ID, &bucket.Name, &bucket.Created)
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (r *BucketsRepo) Get(id uint64) (*models.Bucket, error) {
	query := `SELECT id, name, balance, created_at FROM buckets
		WHERE id=$1`

	var bucket models.Bucket
	var balanceCents int64
	err := r.DB.QueryRow(query, id).
		Scan(&bucket.ID, &bucket.Name, &balanceCents, &bucket.Created)
	if err != nil {
		return nil, err
	}
	bucket.Balance = models.Amount{Cents: balanceCents}
	return &bucket, nil
}
