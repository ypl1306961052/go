package main

import (
	"bufio"
	musicManager "com/ypl/music/manager"
	"com/ypl/music/model"
	"com/ypl/music/mp"
	"errors"
	"fmt"
	"os"
	"strings"
)

var lib *musicManager.MusicManager
var i uint64 = 0

func main() {
	//初始化文件
	lib = musicManager.NewMusicManager()
	defer func() {
		musicManager.CloseFile()
	}()
	fmt.Println("Music Command:")
	showCommand()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入")
		lineType, _, _ := r.ReadLine()

		line := string(lineType)
		if line == "q" || line == "e" {
			fmt.Println("退出程序")
			break
		}
		if line != "" {
			tokens := strings.Split(line, " ")
			handleLineCommand(tokens)
		}
		fmt.Println("----")

	}
}

func showCommand() {
	fmt.Println("1:list\n" +
		"2:add {音乐的名字}　{音乐的类型}　{音乐的位置}\n" +
		"3:remove {音乐的名字}\n" +
		"4:play {音乐的名字}")
}

func handleLineCommand(tokens []string) {
	switch tokens[0] {
	case "comm":
		showCommand()
	case "list":
		if len(tokens) == 1 {
			fmt.Println("music list :")
			for _, music := range lib.Musics {
				if music.Name != "" {
					fmt.Println(music.Name, " ", music.MusicType, " ", music.MusicAdder)
				}
			}
		} else {
			fmt.Println("输入有误,请输入　list")
		}

	case "add":
		if len(tokens) == 4 {
			i++
			_, e := lib.Add(model.Music{
				Id:         i,
				Name:       tokens[1],
				MusicType:  tokens[2],
				MusicAdder: tokens[3],
			})
			if e != nil {
				fmt.Println(e)
				break
			}
			fmt.Println("添加音乐成功　", tokens[1], " ", tokens[2], " ", tokens[3])
		} else {
			fmt.Println("输入有误,请输入　add　{音乐名字}　{音乐的类型}　{音乐的位置}")
		}
	case "remove":
		if len(tokens) == 2 {
			music, index := lib.Find(tokens[1])
			if music == nil || index == -1 {
				fmt.Println("删除失败,库中没有:", tokens[1])
			}
			m, e := lib.Remove(index)
			if e != nil {
				fmt.Println("删除失败")
			} else {
				fmt.Println("删除成功", m.Name)
			}

		}
	case "play":
		if len(tokens) == 2 {

			music, _ := lib.Find(tokens[1])
			if music == nil {
				fmt.Println("没有找到", tokens[1], "歌曲")
				break
			}
			e := play(music.MusicAdder, music.MusicType)
			if e != nil {
				fmt.Println("播放失败,失败的原因:", e)
			}
		} else {
			fmt.Println("play {音乐的名字}")
		}
	default:
		fmt.Println("输入的指令有误")
	}

}
func play(source, musicType string) error {
	if source == "" || musicType == "" {
		return errors.New("音乐的地址或则类型为空")
	}

	if musicType == "mp3" {
		p := &mp.Mp3Player{
			Play: &mp.Play{Stat: 0, Progress: 0},
		}
		p.Player(source)

	} else if musicType == "avm" {
		p := &mp.AvmPlayer{
			Play: &mp.Play{Stat: 0, Progress: 0},
		}
		p.Player(source)
	}

	return nil
}
