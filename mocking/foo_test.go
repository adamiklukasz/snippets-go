package mocking

import (
	"github.com/golang/mock/gomock"
	"testing"
)

var any = gomock.Any()

func TestGoMockAssertions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockFoo(ctrl)

	mock.EXPECT().Do(any).
		Times(3)

	mock.Do(10)
	mock.Do(10)
}
