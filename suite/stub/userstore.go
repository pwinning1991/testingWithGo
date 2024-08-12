package stub

import "github.com/pwinning1991/testingWithGo/suite"

type UserStore struct{}

func (us *UserStore) Create(user *suite.User) error {
	user.ID = 123
	return nil
}

func (us *UserStore) ByID(id int) (*suite.User, error) {
	return nil, suite.ErrNotFound
}

func (us *UserStore) ByEmail(email string) (*suite.User, error) {
	return nil, nil
}
func (us *UserStore) Delete(user *suite.User) error {
	return nil
}
