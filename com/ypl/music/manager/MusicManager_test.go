package musicManager

import (
	"com/ypl/music/model"
	"fmt"
	"testing"
	"time"
)

func TestMusicManager(t *testing.T) {
	lib := NewMusicManager()
	_, err := lib.Add(model.Music{Name: "a", MusicType: "mp3", MusicAdder: "/d", Id: 1})
	if err != nil {
		fmt.Println("添加失败")
	}
	_, erra := lib.Add(model.Music{Name: "a", MusicType: "mp3", MusicAdder: "/d", Id: 1})
	if erra != nil {
		fmt.Println("添加失败")
	}
	lib.Remove(0)
	fmt.Println(lib.List())
	time.Sleep(time.Second * 2)

}
