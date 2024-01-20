package data

type Repository interface {
	GetAll() ([]*User, error)
	Insert(user User) (int, error)
}
