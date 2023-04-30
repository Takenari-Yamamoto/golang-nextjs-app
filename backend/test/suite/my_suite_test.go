package suite

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MySuite struct {
	suite.Suite
	SharedVariable int
}

func (suite *MySuite) SetupTest() {
	suite.SharedVariable = 42
}

func (suite *MySuite) TearDownTest() {
	suite.SharedVariable = 0
}

func (suite *MySuite) TestExample() {
	suite.Equal(42, suite.SharedVariable)
}

// my_suite_test.go
func TestMySuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}
