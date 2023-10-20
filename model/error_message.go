package model

type ErrorMessage struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}
