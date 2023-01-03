package mocking

import (
	"testing"

	"github.com/golang/mock/gomock"
)

var any = gomock.Any()

func TestGoMockAssertions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockFoo(ctrl)

	mock.EXPECT().Do(any).
		Times(2)

	mock.Do(10)
	mock.Do(10)
}
