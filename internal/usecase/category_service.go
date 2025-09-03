package usecase

import (
	"github.com/ezkahan/meditation-backend/internal/domain"
	"github.com/ezkahan/meditation-backend/internal/repository"
	"github.com/google/uuid"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: r}
}

// CreateCategory makes a new category
func (s *CategoryService) CreateCategory(name, icon string, parentId *uint) (*domain.Category, error) {
	c := &domain.Category{
		ID:       uuid.New().String(),
		Name:     name,
		IconPath: icon,
		ParentId: parentId,
	}
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

// GetCategory fetches by id
func (s *CategoryService) GetCategory(id string) (*domain.Category, error) {
	return s.repo.GetByID(id)
}

// ListCategories returns all
func (s *CategoryService) ListCategories() ([]domain.Category, error) {
	return s.repo.List()
}

// UpdateCategory changes fields
func (s *CategoryService) UpdateCategory(id, name, icon string, parentId *uint) error {
	c := &domain.Category{
		ID:       id,
		Name:     name,
		IconPath: icon,
		ParentId: parentId,
	}
	return s.repo.Update(c)
}

// DeleteCategory removes one
func (s *CategoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}
