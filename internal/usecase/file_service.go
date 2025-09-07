package usecase

import (
	"time"

	"github.com/ezkahan/go-file-service/internal/domain"
	"github.com/ezkahan/go-file-service/internal/repository"
	"github.com/google/uuid"
)

type FileService interface {
	CreateFile(name, iconPath, filePath string, categoryID *string) (*domain.File, error)
	GetFile(id string) (*domain.File, error)
	ListFiles() ([]domain.File, error)
	UpdateFile(id, name, iconPath, filePath string, categoryID *string) error
	DeleteFile(id string) error
}

type fileService struct {
	repo repository.FileRepository
}

func NewFileService(r repository.FileRepository) FileService {
	return &fileService{repo: r}
}

// CreateFile makes a new file record
func (s *fileService) CreateFile(name, iconPath, filePath string, categoryID *string) (*domain.File, error) {
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
func (s *fileService) GetFile(id string) (*domain.File, error) {
	return s.repo.GetByID(id)
}

// ListFiles returns all files
func (s *fileService) ListFiles() ([]domain.File, error) {
	return s.repo.List()
}

// UpdateFile modifies an existing file
func (s *fileService) UpdateFile(id, name, iconPath, filePath string, categoryID *string) error {
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
func (s *fileService) DeleteFile(id string) error {
	return s.repo.Delete(id)
}
