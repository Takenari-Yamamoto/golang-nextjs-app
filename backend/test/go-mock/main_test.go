package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock "golang-nextjs-app/test/go-mock/mock"
)

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGreeter := mock.NewMockGreeter(ctrl)
	mockGreeter.EXPECT().Greet("John").Return("Hello, John!")

	result := SayHello(mockGreeter, "John")
	assert.Equal(t, "Hello, John!", result)
}
