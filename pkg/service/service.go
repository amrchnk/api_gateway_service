package service

type Implementation struct {
	AuthService
}

func NewApiGWService(as AuthService) *Implementation {
	return &Implementation{
		AuthService:     as,
	}
}