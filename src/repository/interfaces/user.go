package interfaces

import "muhammedarifp/cleancode-example/src/domain"

type UserRepository interface {
	AddUser(usr *domain.User)
}
