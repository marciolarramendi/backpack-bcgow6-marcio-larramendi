package users

type Service interface {
	GetAll() ([]Usuario, error)
	Store(nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Usuario, error) {
	usuarios, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (s *service) Store(nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Usuario{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fechacreacion)
	if err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}

func (s *service) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error) {

	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechacreacion)
}
