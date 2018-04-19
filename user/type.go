package user

var New = func() (*User) {
	return &User{}
}
// ---- ---- ---- ----
type User struct {
	Id          int64      `ez:"name(id);type(int)"`
	FirstName   string     `ez:"name(first_name);type(varchar)"`
	LastName    string     `ez:"name(last_name);type(varchar)"`
	Group       string	   `ez:"name(group_name);type(varchar)"`
}

func(u *User) TableName() string {
	return "u_user"
}
