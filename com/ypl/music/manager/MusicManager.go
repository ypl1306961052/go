package musicManager

import (
	"com/ypl/music/model"
	"errors"
	"fmt"
	"os"
)

type MusicManager struct {
	//切片音乐
	Musics []model.Music
}

var file *os.File
var fileName = "music.txt"

// 1 直接读取　2:使用buf 3:Io
var fileReadType = 1

//获取文件
func getFile() *os.File {
	if file == nil {
		//不存在　打开文件
		file = openFile()
	}
	return file
}
func CloseFile() {
	if file != nil {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件" + fileName + "失败")
		}
	}
}

//打开文件
func openFile() *os.File {
	mFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 077)
	if err != nil {

		fileNotExist := os.IsNotExist(err)
		if fileNotExist {
			newFile, e := os.Create(fileName)
			if e != nil {
				fmt.Println("创建文件失败,失败的原因:", e)
				panic("创建" + fileName + "失败")
			}
			mFile = newFile
		} else {
			fmt.Println("打开"+fileName+"失败,", err)
			panic("打开" + fileName + "失败")
		}
	}
	return mFile
}
func initMusicList(m *[]model.Music) {
	mFile := getFile()
	switch fileReadType {
	case 1:
		ReadFileContent(mFile, m)
	case 2:
		ReadFileBufContent(mFile, m)
	case 3:
		ReadFileIoContent(mFile, m)
	}
}

//保存音乐文件成功
func saveMusicToFile(ms *[]model.Music) {
	var line = ""
	for _, musicEntry := range *ms {
		line = line + musicEntry.Name + " " + musicEntry.MusicType + " " + musicEntry.MusicAdder + "\n"
	}
	mFile := getFile()
	//清空文件
	tError := os.Truncate(mFile.Name(), 0)
	if tError != nil {
		fmt.Println("清空文件失败,", tError)
	}
	_, err := mFile.Write([]byte(line))
	//if index == len(line) {
	//	fmt.Println("保存音乐到文件成功")
	//}
	if err != nil {
		fmt.Println("保存音乐到文件失败,失败的原因:", err)
	}
}

//创建音乐管理
func NewMusicManager() *MusicManager {
	manager := MusicManager{Musics: make([]model.Music, 0)}
	initMusicList(&manager.Musics)
	if len(manager.Musics) == 0 {
		fmt.Println("==========================================================")
		fmt.Println("音乐列表为空")
		fmt.Println("==========================================================")
	} else {
		fmt.Println("==========================================================")
		for _, music := range manager.Musics {

			fmt.Println(music.Id, " ", music.Name, " ", music.MusicType, " ", music.MusicAdder)
		}
		fmt.Println("==========================================================")

	}
	return &manager
}

//列表
func (m *MusicManager) List() *[]model.Music {
	return &m.Musics
}

//添加
func (m *MusicManager) Add(musicEntry model.Music) (*[]model.Music, error) {
	b := m.Musics
	if !checkExist(&b) { //不存在
		return nil, errors.New("列表没有初始化")
	}

	m.Musics = append(m.Musics, musicEntry)
	//开启线程保存
	go saveMusicToFile(&m.Musics)
	return &m.Musics, nil

}

//删除
func (m *MusicManager) Remove(index int) (*model.Music, error) {
	if index < 0 || index > len(m.Musics) {
		return nil, errors.New("该index" + string(index) + "不存在")
	}
	if len(m.Musics) == 0 {
		return nil, errors.New("数据为空")
	}

	oldMusic := m.Musics[index]
	if index > 0 && index < len(m.Musics)-1 { //中间
		m.Musics = append(m.Musics[:index-1], m.Musics[index+1:]...)
	} else if index == 0 { //第一个
		m.Musics = m.Musics[1:]
	} else {
		m.Musics = m.Musics[:index-1]
	}
	//开启线程保存
	go saveMusicToFile(&m.Musics)
	return &oldMusic, nil

}

//查询
func (m *MusicManager) Get(index int) *model.Music {
	b := m.Musics
	if !checkExist(&b) { //不存在
		return nil
	}
	if index < 0 || index > len(b) {
		return nil
	}
	return &m.Musics[index]
}

//听过名字　查询名字
func (m *MusicManager) Find(name string) (*model.Music, int) {
	if !checkExist(&m.Musics) {
		return nil, -1
	}
	for index, music := range m.Musics {
		if name == music.Name {
			return &music, index
		}
	}
	//没有查询到数据
	return nil, -1
}
func checkExist(ms *[]model.Music) bool {
	if ms == nil {
		return false
	} else {
		return true
	}

}
