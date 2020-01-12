package client_tldr_sh

type GetIndexOutput struct {
	Commands []Commands `json:"commands"`
}
type Targets struct {
	Os       string `json:"os"`
	Language string `json:"language"`
}
type Commands struct {
	Name     string    `json:"name"`
	Platform []string  `json:"platform"`
	Language []string  `json:"language"`
	Targets  []Targets `json:"targets"`
}

type GetIndexInput struct {
}
