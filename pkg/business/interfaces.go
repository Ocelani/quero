package business

// Repository interface allows us to access the CRUD operations of the database.
type Repository interface {
	Create(data *Business) error
	ReadOne(data *Business) error
	ReadAll(data *[]Business) error
	Update(data *Business) error
	Delete(data *Business) error
}

// Service interface methods to access our repository operations.
type Service interface {
	Repository
}
