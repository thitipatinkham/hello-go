package controller

// User ...
type User struct {
	ID   string
	Name string
}

// GetUsers return list user
func GetUsers() []*User {
	users := []*User{
		{ID: "1", Name: "11"},
		{ID: "2", Name: "22"},
		{ID: "3", Name: "33"},
		{ID: "4", Name: "44"},
		{ID: "5", Name: "55"},
	}
	// user := User{ID: "1", Name: "AAAA"}
	// users = append(users, user)
	return users
}
