package usecase

import (
	"time"

	"github.com/ezkahan/meditation-backend/internal/domain"
	"github.com/ezkahan/meditation-backend/internal/repository"
	"github.com/google/uuid"
)

type FileService struct {
	repo repository.FileRepository
}

func NewFileService(r repository.FileRepository) *FileService {
	return &FileService{repo: r}
}

// CreateFile makes a new file record
func (s *FileService) CreateFile(name, iconPath, filePath string, categoryID *string) (*domain.File, error) {
	f := &domain.File{
		ID:         uuid.New().String(),
		Name:       name,
		IconPath:   iconPath,
		FilePath:   filePath,
		CategoryID: categoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := s.repo.Create(f); err != nil {
		return nil, err
	}
	return f, nil
}

// GetFile fetches a file by id
func (s *FileService) GetFile(id string) (*domain.File, error) {
	return s.repo.GetByID(id)
}

// ListFiles returns all files
func (s *FileService) ListFiles() ([]domain.File, error) {
	return s.repo.List()
}

// UpdateFile modifies an existing file
func (s *FileService) UpdateFile(id, name, iconPath, filePath string, categoryID *string) error {
	f := &domain.File{
		ID:         id,
		Name:       name,
		IconPath:   iconPath,
		FilePath:   filePath,
		CategoryID: categoryID,
		UpdatedAt:  time.Now(),
	}
	return s.repo.Update(f)
}

// DeleteFile removes a file by id
func (s *FileService) DeleteFile(id string) error {
	return s.repo.Delete(id)
}
