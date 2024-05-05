package todo

type user struct {
	Id       int    `json:"*"`
	Name     string `json:"name"`
	UserName string `json:"user"`
	Password string `json:"password"`
}
