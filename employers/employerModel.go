package employers

import "time"

// Employer representa a estrutura do empregador no banco de dados
type Employer struct {
	EmployerID        string
	Name              string
	Email             string
	MobileCountryCode int
	MobileAreaCode    int
	MobileNumber      string
	Password          string
	HasPicture        bool
	RegisterDate      time.Time
	LastUpdateDate    time.Time
}

// SignUpRequest entrada do método SignUp
type SignUpRequest struct {
	Name              string
	Email             string
	MobileCountryCode int
	MobileAreaCode    int
	MobileNumber      string
	Password          string
	PasswordCheck     string
}

// SignUpResponse retorno do método SignUp
type SignUpResponse struct {
	EmployerID        string
	Name              string
	Email             string
	MobileCountryCode int
	MobileAreaCode    int
	MobileNumber      string
	Password          string
	HasPicture        bool
	RegisterDate      time.Time
	LastUpdateDate    time.Time
}
