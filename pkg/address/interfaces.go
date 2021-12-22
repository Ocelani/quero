package address

// Repository interface allows us to access the CRUD operations of the database.
type Repository interface {
	Create(data *Address) error
	ReadOne(data *Address) error
	ReadAll(data *[]Address) error
	Update(data *Address) error
	Delete(data *Address) error
}

// Service interface methods to access our repository operations.
type Service interface {
	Repository
}
