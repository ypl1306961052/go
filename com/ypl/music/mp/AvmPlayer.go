package mp

import (
	"fmt"
	"time"
)

func (a *AvmPlayer) Player(source string) {
	fmt.Println("avm播放位置为:", source)
	a.Stat = 1
	defer func() {
		fmt.Println("播放完毕")
		a.Stat = 0
	}()
	defer func() {
		fmt.Println("播放完毕")
		a.Stat = 0
		a.Progress = 0
	}()
	var index = 0
	for {
		if index >= 100 {
			break
		}
		fmt.Print("播放", a.Progress, "%")
		fmt.Print("...")
		time.Sleep(time.Second)
		index = index + 10
		a.Progress = index / 100
		fmt.Println()

	}
}
