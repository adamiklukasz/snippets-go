package mocking

//go:generate mockgen -source $GOFILE -destination foo_mock.go -package mocking
type Foo interface {
	Do(int) int
}
