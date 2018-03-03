package testify

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 12, 12)
}

func TestAdd1(t *testing.T) {
	assert.Equal(t, Add(2, 10), 12)
}

/*
  Test objects
*/

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MyMockedObject struct {
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MyMockedObject) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestSomething(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// setup expectations
	testObj.On("DoSomething", 123).Return(true, nil)

	// call the code we are testing
	//targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)

}
