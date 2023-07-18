package token

type Data struct {
	UserId   string
	Username string
	Roles    []string
}

func (d Data) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["user_id"] = d.UserId
	m["username"] = d.Username
	m["roles"] = d.Roles
	return m
}
