package util

import "os/user"

// User provides a more flexible way to deal with current user
// information allowing applications to change the effective user ID
// without using the original user and enabling mocks for tests that
// involve user information.
var User *user.User

func init() {
	User, _ = user.Current()
	if User == nil {
		User = new(user.User)
	}
}
