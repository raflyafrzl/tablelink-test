package user

type User struct {
	Id       int    `db:"id"`
	Xid      string `db:"xid"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
