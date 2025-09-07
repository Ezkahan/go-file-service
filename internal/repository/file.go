package repository

import (
	"context"
	"time"

	"github.com/ezkahan/go-file-service/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FileRepo struct {
	Pool *pgxpool.Pool
}

type FileRepository interface {
	Create(file *domain.File) error
	GetByID(id string) (*domain.File, error)
	List() ([]domain.File, error)
	Update(file *domain.File) error
	Delete(id string) error
}

func NewFileRepo(pool *pgxpool.Pool) *FileRepo {
	return &FileRepo{Pool: pool}
}

// Create inserts a new file
func (r *FileRepo) Create(f *domain.File) error {
	_, err := r.Pool.Exec(context.Background(),
		`INSERT INTO files (id, name, icon_path, file_path, category_id, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		f.ID, f.Name, f.IconPath, f.FilePath, f.CategoryID, f.CreatedAt, f.UpdatedAt,
	)
	return err
}

// GetByID fetches a file by id
func (r *FileRepo) GetByID(id string) (*domain.File, error) {
	row := r.Pool.QueryRow(context.Background(),
		`SELECT id, name, icon_path, file_path, category_id, created_at, updated_at
		 FROM files WHERE id=$1`, id,
	)

	var f domain.File
	if err := row.Scan(&f.ID, &f.Name, &f.IconPath, &f.FilePath, &f.CategoryID, &f.CreatedAt, &f.UpdatedAt); err != nil {
		return nil, err
	}
	return &f, nil
}

// List returns all files
func (r *FileRepo) List() ([]domain.File, error) {
	rows, err := r.Pool.Query(context.Background(),
		`SELECT id, name, icon_path, file_path, category_id, created_at, updated_at FROM files`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []domain.File
	for rows.Next() {
		var f domain.File
		if err := rows.Scan(&f.ID, &f.Name, &f.IconPath, &f.FilePath, &f.CategoryID, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, nil
}

// Update modifies file fields
func (r *FileRepo) Update(f *domain.File) error {
	_, err := r.Pool.Exec(context.Background(),
		`UPDATE files SET name=$1, icon_path=$2, file_path=$3, category_id=$4, updated_at=$5 WHERE id=$6`,
		f.Name, f.IconPath, f.FilePath, f.CategoryID, time.Now(), f.ID,
	)
	return err
}

// Delete removes a file by id
func (r *FileRepo) Delete(id string) error {
	_, err := r.Pool.Exec(context.Background(),
		`DELETE FROM files WHERE id=$1`, id,
	)
	return err
}
