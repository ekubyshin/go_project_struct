package manual

// --- Абстракция (интерфейс) ---

type userRepository interface {
	FindByID(id int) (string, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo userRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Greet(id int) (string, error) {
	name, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Привет, %s!", name), nil
}