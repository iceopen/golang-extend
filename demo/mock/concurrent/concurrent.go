//go:generate mockgen -destination mock/concurrent_mock.go github.com/iceopen/golang-extend/demo/concurrent Math

// Package concurrent demonstrates how to use gomock with goroutines.
package concurrent

type Math interface {
	Sum(a, b int) int
}
