package define

type Model struct {
	Table  string  `json:"table"`
	Define *Define `json:"define"`
}
