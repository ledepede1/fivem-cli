package config

type Structures struct {
	StructData map[string]StructData `json:"structures"`
}

type StructData struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

type UserName struct {
	UseUserName bool   `json:"useName"`
	Username    string `json:"username"`
}
