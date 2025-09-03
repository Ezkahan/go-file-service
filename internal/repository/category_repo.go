package repository

import (
	"context"
	"time"

	"github.com/ezkahan/meditation-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepo struct {
	Pool *pgxpool.Pool
}

type CategoryRepository interface {
	Create(category *domain.Category) error
	GetByID(id string) (*domain.Category, error)
	List() ([]domain.Category, error)
	Update(category *domain.Category) error
	Delete(id string) error
}

func NewCategoryRepo(pool *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{Pool: pool}
}

// Create inserts a new category
func (r *CategoryRepo) Create(c *domain.Category) error {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	_, err := r.Pool.Exec(context.Background(),
		`INSERT INTO categories (id, name, icon_path, parent_id, created_at, updated_at) 
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		c.ID, c.Name, c.IconPath, c.ParentId, c.CreatedAt, c.UpdatedAt,
	)
	return err
}

// GetByID retrieves a category by UUID
func (r *CategoryRepo) GetByID(id string) (*domain.Category, error) {
	row := r.Pool.QueryRow(context.Background(),
		`SELECT id, name, icon_path, parent_id, created_at, updated_at 
		 FROM categories WHERE id=$1`, id,
	)

	var c domain.Category
	if err := row.Scan(&c.ID, &c.Name, &c.IconPath, &c.ParentId, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return nil, err
	}
	return &c, nil
}

// List retrieves all categories
func (r *CategoryRepo) List() ([]domain.Category, error) {
	rows, err := r.Pool.Query(context.Background(),
		`SELECT id, name, icon_path, parent_id, created_at, updated_at 
		 FROM categories ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.IconPath, &c.ParentId, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

// Update modifies an existing category
func (r *CategoryRepo) Update(c *domain.Category) error {
	c.UpdatedAt = time.Now()

	_, err := r.Pool.Exec(context.Background(),
		`UPDATE categories 
		 SET name=$1, icon_path=$2, parent_id=$3, updated_at=$4 
		 WHERE id=$5`,
		c.Name, c.IconPath, c.ParentId, c.UpdatedAt, c.ID,
	)
	return err
}

// Delete removes a category by UUID
func (r *CategoryRepo) Delete(id string) error {
	_, err := r.Pool.Exec(context.Background(),
		`DELETE FROM categories WHERE id=$1`, id,
	)
	return err
}
