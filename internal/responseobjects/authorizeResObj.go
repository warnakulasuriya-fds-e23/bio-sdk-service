package responseobjects

type AuthorizeResObj struct {
	ActionStatus string               `json:"actionStatus"`
	Data         AuthorizeResObj_Data `json:"data"`
}

type AuthorizeResObj_Data struct {
	User AthorizeResObj_User `json:"user"`
}

type AthorizeResObj_User struct {
	Id        string                    `json:"id"`
	Claims    []AuthorizeResObj_Claims  `json:"claims"`
	UserStore AuthorizeResObj_UserStore `json:"userStore"`
}

type AuthorizeResObj_Claims struct {
	Uri   string `json:"uri"`
	Value string `json:"value"`
}

type AuthorizeResObj_UserStore struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
