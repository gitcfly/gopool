# gopool
golang实现的协程池

```
import (
	"fmt"
	"testing"

	"github.com/gitcfly/gopool"
)

func TestGoPool(t *testing.T) {
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
```