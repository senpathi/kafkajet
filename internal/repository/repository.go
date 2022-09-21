package repository

type DB interface {
	Table(name string) Repo
	Close()
}

type Repo interface {
	Read(filter map[string]interface{}) (value interface{}, err error)
	Write(value interface{}) error
	Delete(filter map[string]interface{})
}
