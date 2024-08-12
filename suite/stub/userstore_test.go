package stub_test

import (
	"github.com/pwinning1991/testingWithGo/suite"
	"github.com/pwinning1991/testingWithGo/suite/stub"
	"github.com/pwinning1991/testingWithGo/suite/suitetest"
	"testing"
)

var _ suite.UserStore = &stub.UserStore{}

func TestUserStore(t *testing.T) {
	us := &stub.UserStore{}
	suitetest.UserStore(t, us)

}
