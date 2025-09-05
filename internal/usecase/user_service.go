package usecase

import (
	"strings"

	"github.com/ezkahan/meditation-backend/internal/delivery/http/validators"
	"github.com/ezkahan/meditation-backend/internal/domain"
	"github.com/ezkahan/meditation-backend/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	VerifyCredential(ctx *gin.Context, username string, password string) (*domain.User, error)
	Save(ctx *gin.Context, req validators.SaveUserRequest) (domain.User, error)
	List(page int, limit int) (domain.UsersWithPaginate, error)
	GetByID(id uint) (*domain.User, error)
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{repo: userRepo}
}

func (s *userService) Save(ctx *gin.Context, req validators.SaveUserRequest) (domain.User, error) {
	user := domain.User{
		ID:        req.ID,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Role:      req.Role,
		Ip:        ctx.ClientIP(),
		Device:    ctx.Request.UserAgent(),
	}

	user, err := s.repo.Save(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *userService) VerifyCredential(ctx *gin.Context, username string, password string) (*domain.User, error) {
	ip := ctx.ClientIP()
	device := ctx.Request.UserAgent()
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	// fetch user by username
	u, err := s.repo.VerifyCredential(username, password)
	if err != nil {
		return nil, err
	}

	// compare hashed password
	if !s.CompareHashAndPassword(u.Password, []byte(password)) {
		return nil, err
	}

	// update user IP + device
	_ = s.repo.SaveUserData(u.ID, ip, device)

	return u, nil
}

func (s *userService) CompareHashAndPassword(hashed string, plain []byte) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), plain) == nil
}

func (s *userService) List(page int, limit int) (domain.UsersWithPaginate, error) {
	var usersWithPaginate domain.UsersWithPaginate

	userList, total, err := s.repo.List(page, limit)
	if err != nil {
		return usersWithPaginate, err
	}

	lastPage := int((total + int64(limit) - 1) / int64(limit))
	if page > lastPage {
		return usersWithPaginate, nil
	}

	usersWithPaginate.Users = userList
	usersWithPaginate.LastPage = lastPage
	usersWithPaginate.Total = total
	return usersWithPaginate, nil
}

func (s *userService) GetByID(id uint) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
