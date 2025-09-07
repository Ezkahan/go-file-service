package validators

import "github.com/ezkahan/go-file-service/internal/domain"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type SaveUserRequest struct {
	ID        uint        `json:"id"`
	Username  string      `json:"username"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Firstname string      `json:"firstname"`
	Lastname  string      `json:"lastname"`
	Role      domain.Role `json:"role"`
	IP        string      `json:"ip" swaggerignore:"true"`
	Device    string      `json:"device" swaggerignore:"true"`
}

// type AddRequest struct {
// 	Username  string `json:"username" validate:"required"`
// 	Password  string `json:"password" validate:"required"`
// 	Phone     string `json:"phone" validate:"required"`
// 	Email     string `json:"email"`
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// 	IP        string `json:"ip" swaggerignore:"true"`
// 	Device    string `json:"device" swaggerignore:"true"`
// }
