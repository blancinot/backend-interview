package user

import "context"

type User struct {
	ID   string
	Name string
}

type Filter struct {
	ID string
}

type Store interface {
	Fetch(context.Context, Filter) (User, error)
}

type App interface {
	Store
}
