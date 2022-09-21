package repository

type DBConfig struct {
	Database      string
	User          string
	Password      string
	Address       string
	MaxConnection int
}
