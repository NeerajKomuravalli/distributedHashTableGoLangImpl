package model

type UserSentData struct {
	Id    string `json:id`
	Value int    `json:value`
}

type ResponseData struct {
	Success      bool         `json:success`
	ErrorMessage string       `json:error,omitempty`
	Data         UserSentData `json:data,omitempty`
}
