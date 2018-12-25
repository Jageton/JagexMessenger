package chat

import (
	"sync"
)

type UserManager struct {
	users map[int64]*User
	mutex *sync.RWMutex
}


func NewUserManager() *UserManager {
	return &UserManager{
		users : make(map[int64]*User),
		mutex : &sync.RWMutex{},
	}
}


func (u *UserManager) Add(id int64, user *User){
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.users[id] = user
}

func (u *UserManager) Remove(id int64) bool {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	var ok bool
	if _, ok = u.users[id]; ok {
		delete(u.users, id)
	}
	return ok
}

func (u *UserManager) Get(id int64) (*User, bool)  {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	user, ok := u.users[id]
	return user, ok
}

func (u *UserManager) Contains(id int64) bool {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	_, b := u.users[id]
	return b
}

func (u *UserManager) Ides() (ides []int64) {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	for id := range u.users {
		ides = append(ides, id)
	}
	return
}

func (u *UserManager) Users() (users []*User) {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	for _, user := range u.users {
		users = append(users, user)
	}
	return
}
