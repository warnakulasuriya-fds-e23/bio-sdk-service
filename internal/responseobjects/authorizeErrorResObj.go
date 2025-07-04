package responseobjects

type AuthorizeErrorResObj400 struct {
	ActionStatus     string `json:"actionStatus"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}

type AuthorizeErrorResObj500 struct {
	ActionStatus     string `json:"actionStatus"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}
