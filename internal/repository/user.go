package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ezkahan/go-file-service/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	VerifyCredential(username string, password string) (*domain.User, error)
	SaveUserData(userID uint, ip, device string) error
	List(page, limit int) ([]domain.User, int64, error)
	GetByID(id uint) (*domain.User, error)
	Delete(id uint) error
}

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) Save(user domain.User) (domain.User, error) {
	query := `
		INSERT INTO users (id, username, firstname, lastname, password, ip, device)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (id) DO UPDATE
		SET username=$2, firstname=$3, lastname=$4, password=$5, ip=$6, device=$7
		RETURNING id
	`
	_, err := r.pool.Exec(context.Background(), query,
		user.ID, user.Username, user.Firstname, user.Lastname,
		user.Password, user.Ip, user.Device,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) List(page, limit int) ([]domain.User, int64, error) {
	offset := (page - 1) * limit

	// Count total
	var total int64
	err := r.pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Fetch paginated users
	rows, err := r.pool.Query(context.Background(),
		"SELECT id, username, firstname, lastname, ip, device, created_at, updated_at FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2",
		limit, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(
			&u.ID, &u.Username, &u.Firstname, &u.Lastname,
			&u.Ip, &u.Device, &u.CreatedAt, &u.UpdatedAt,
		); err != nil {
			return nil, total, err
		}
		users = append(users, u)
	}
	return users, total, nil
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var u domain.User
	err := r.pool.QueryRow(context.Background(),
		"SELECT id, username, firstname, lastname, ip, device, created_at, updated_at FROM users WHERE id=$1",
		id,
	).Scan(
		&u.ID, &u.Username, &u.Firstname, &u.Lastname,
		&u.Ip, &u.Device, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) VerifyCredential(username string, password string) (*domain.User, error) {
	var u domain.User
	err := r.pool.QueryRow(context.Background(),
		"SELECT id, username, password FROM users WHERE username=$1",
		username,
	).Scan(&u.ID, &u.Username, &u.Password)

	if err != nil {
		return nil, err
	}

	// ⚠️ Compare password hashes in service layer
	if u.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return &u, nil
}

func (r *userRepository) SaveUserData(userID uint, ip, device string) error {
	cmd, err := r.pool.Exec(context.Background(),
		"UPDATE users SET ip=$1, device=$2, updated_at=NOW() WHERE id=$3",
		ip, device, userID,
	)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("no user found with id: %d", userID)
	}
	return nil
}

func (r *userRepository) Delete(id uint) error {
	cmd, err := r.pool.Exec(context.Background(), "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("no user found with id: %d", id)
	}
	return nil
}
