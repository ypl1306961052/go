package mp

import (
	"fmt"
	"time"
)

func (m *Mp3Player) Player(source string) {
	fmt.Println("mp3音乐的位置为:", source)
	//播放
	m.Stat = 1
	defer func() {
		fmt.Println("播放完毕")
		m.Stat = 0
		m.Progress = 0
	}()
	var index = 0
	for {
		if index >= 100 {
			break
		}
		fmt.Print("播放", index/100, "%")
		fmt.Print(">>>")
		time.Sleep(time.Second)
		index = index + 10
		m.Progress = index

	}
}
