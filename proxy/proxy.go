package proxy

import (
	"errors"
)

type UserFinder interface {
	FindUser(id int) (*User, error)
}

type User struct {
	ID int
}

type UserList []*User

func (u *UserList) FindUser(id int) (*User, error) {
	for _, user := range *u {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("not found")
}

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) AddUserToStack(user *User) {
	if len(u.StackCache) == 0 {
		u.StackCache = []*User{user}
	} else {
		u.StackCache = append(u.StackCache, user)
		if len(u.StackCache) > u.StackCapacity {
			u.StackCache = u.StackCache[1:]
		}
	}
}

func (u *UserListProxy) FindUser(id int) (*User, error) {
	if u.SomeDatabase == nil {
		return nil, errors.New("empty database")
	}

	if user, err := u.StackCache.FindUser(id); err == nil {
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err := u.SomeDatabase.FindUser(id)
	if err != nil {
		return nil, err
	}

	u.AddUserToStack(user)
	u.DidLastSearchUsedCache = false

	return user, nil
}
