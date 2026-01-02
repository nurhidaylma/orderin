package user

type Role string

const (
	RoleUser     Role = "USER"
	RoleMerchant Role = "MERCHANT"
	RoleAdmin    Role = "ADMIN"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      Role
	CreatedAt string
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsMerchant() bool {
	return u.Role == RoleMerchant
}
