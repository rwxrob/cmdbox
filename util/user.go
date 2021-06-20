/*
Copyright 2021 Robert S. Muhlestein.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
