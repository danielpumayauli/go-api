package main

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (t *Message) setStatus(data string) {
	t.Status = data
}

func (t *Message) setMessage(data string) {
	t.Message = data
}
