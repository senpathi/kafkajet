package domain

type Error struct {
	Err     error  `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Err.Error()
}
