package employee

// Repository interface allows us to access the CRUD operations of the database.
type Repository interface {
	Create(data *Employee) error
	ReadOne(data *Employee) error
	ReadAll(data *[]Employee) error
	Update(data *Employee) error
	Delete(data *Employee) error
}

// Service interface methods to access our repository operations.
type Service interface {
	Repository
}
