package domain

var connectionString string = "/Users/jstoddard/microservice-in-go.db"

type CustomerRepositoryDb struct {
}

func (reciever CustomerRepositoryDb) FindAll() ([]Customer, error) {
	// db, err :=
}
