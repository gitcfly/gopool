package example

import (
	"fmt"
	"testing"

	"github.com/gitcfly/gopool"
)

func TestGoPool(t *testing.T) {
	// the first params is max goroutine count,default value is runtime.NumCPU()
	// the second params is max task cap,default is 100
	// p := gopool.NewPool(0, 0) // use the default params value
	p := gopool.NewPool(2, 100)
	p.Execute(func() error {
		fmt.Println("run task 1")
		return nil
	})
	p.Execute(func() error {
		fmt.Println("run task 2")
		return nil
	})
	p.Execute(func() error {
		fmt.Println("run task 3")
		return nil
	})
	p.Execute(func() error {
		fmt.Println("run task 4")
		return nil
	})
	err := p.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
