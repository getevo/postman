package postman

type BodyMode string

const (
	BodyModeForm       BodyMode = "formdata"
	BodyModeURLEncoded BodyMode = "urlencoded"
	BodyModeRaw        BodyMode = "raw"
)

type Body struct {
	Mode       BodyMode   `json:"mode"`
	Raw        string     `json:"raw,omitempty"`
	Urlencoded []KeyValue `json:"urlencoded,omitempty"`
	FormData   []KeyValue `json:"formdata,omitempty"`
	File       *File      `json:"file,omitempty"`
}

type KeyValue struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

type File struct {
	Src     string `json:"src,omitempty"`
	Content string `json:"content,omitempty"`
}
