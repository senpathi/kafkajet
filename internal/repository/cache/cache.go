package cache

type Cache interface {
	Read(id string) (value interface{}, exists bool)
	Write(id string, value interface{}) error
	Delete(id string)
}
