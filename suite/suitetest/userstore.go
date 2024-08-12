package suitetest

import (
	"github.com/pwinning1991/testingWithGo/suite"
	"testing"
)

func UserStore(t *testing.T, us suite.UserStore) {
	_, err := us.ByID(123)
	if err != suite.ErrNotFound {
		t.Errorf("ByID(123) err = nil, wanted ErrNotFound")
	}

	t.Run("create", func(t *testing.T) {
		user := suite.User{
			Email: "p@testing.io",
		}
		err = us.Create(&user)
		if err != nil {
			t.Errorf("Create() err = %v, wanted nil", err)
		}
		if user.ID <= 0 {
			t.Errorf("Create() user.ID = %d; want a positive integer", user.ID)
		}
	})

	t.Run("ByID", func(t *testing.T) {
		//setup
		user := suite.User{
			Email: "p@testing.io",
		}
		_ = us.Create(&user)

		// teardown
		defer func() {
			us.Delete(&user)
		}()

		got, err := us.ByID(user.ID)
		if err != nil {
			t.Errorf("ByID() err = %v, wanted nil", err)
		}
		if got != &user {
			t.Errorf("ByID() got = %v, wanted = %v", got, user)
		}
	})

}
