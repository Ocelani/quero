package address

import "quero2pay/internal/database"

type DefaultRepository struct {
	conn database.Connector
}

func NewDefaultRepository(conn database.Connector) *DefaultRepository {
	if err := conn.GetDatabase().AutoMigrate(&Address{}); err != nil {
		panic(err)
	}
	return &DefaultRepository{conn}
}

func (r *DefaultRepository) Create(data *Address) error {
	db := r.conn.GetDatabase()
	db.FirstOrCreate(data)
	return db.Error
}

func (r *DefaultRepository) ReadOne(data *Address) error {
	db := r.conn.GetDatabase()
	db.First(data)
	return db.Error
}

func (r *DefaultRepository) ReadAll(data *[]Address) error {
	db := r.conn.GetDatabase()
	db.Find(data)
	return db.Error
}

func (r *DefaultRepository) Update(data *Address) error {
	db := r.conn.GetDatabase()
	db.Save(data)
	return db.Error
}

func (r *DefaultRepository) Delete(data *Address) error {
	db := r.conn.GetDatabase()
	db.Delete(data)
	return db.Error
}
