package thread

import (
	"fmt"
	"sync"
)

func SayHello(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()

}
