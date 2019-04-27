package wshub

type Message struct {
	ID   int
	Data interface{}
}

/*type Message struct {
	ID          int       `json:"id,omitempty"`
	Data        string    `json:"message"`
}*/