package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	if _, ok := users[name]; !ok {
		deleted, err = false, errors.New("not found")
		return deleted, err
	}

	userInfo := users[name]
	if userInfo.scheduledForDeletion {
		deleted, err = true, nil
		delete(users, name)
	} else {
		deleted, err = false, nil
	}
	return deleted, err
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
