package address

type DefaultService struct {
	repository Repository
}

func NewDefaultService(repo Repository) *DefaultService {
	return &DefaultService{
		repository: repo,
	}
}

func (b *DefaultService) Create(data *Address) error {
	return b.repository.Create(data)
}

func (b *DefaultService) ReadOne(data *Address) error {
	return b.repository.ReadOne(data)
}

func (b *DefaultService) ReadAll(data *[]Address) error {
	return b.repository.ReadAll(data)
}

func (b *DefaultService) Update(data *Address) error {
	return b.repository.Update(data)
}

func (b *DefaultService) Delete(data *Address) error {
	return b.repository.Delete(data)
}
