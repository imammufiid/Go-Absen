package user

type Service interface {
	// FunctionName(param dataType)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

// Write code here, implementation interface