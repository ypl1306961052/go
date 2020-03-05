package musicManager

import (
	"bufio"
	"com/ypl/music/model"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type FileManager interface {
	readFileContent(file *os.File)
}

//没有使用缓存
func ReadFileContent(file *os.File, m *[]model.Music) {
	var tmp = make([]byte, 1024)
	var tmpStr string
	for {
		count, err := file.Read(tmp)
		if err != nil {
			if err != io.EOF {

				fmt.Println("读取music文件失败,失败的原因:", err)
				break
			}
		}
		if count <= 0 {
			//退出
			break
		}
		tmpStr = tmpStr + string(tmp)
	}
	if tmpStr != "" {
		lines := strings.Split(tmpStr, "\n")
		for index, line := range lines {
			if line != "" {
				cells := strings.Split(line, " ")
				if len(cells) == 3 {
					*m = append(*m, model.Music{
						Id:         uint64(index),
						Name:       cells[0],
						MusicType:  cells[1],
						MusicAdder: cells[2],
					})
				}
			}

		}
	}
}

//使用缓存
func ReadFileBufContent(file *os.File, m *[]model.Music) {
	var reader = bufio.NewReader(file)
	var index = 0
	for {
		index++
		lineByte, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取文件失败,失败的原因:", err)
			break
		}
		line := string(lineByte)
		lineArr := strings.Split(line, " ")
		if len(lineArr) != 3 {
			continue
		}
		*m = append(*m, model.Music{
			Id:         uint64(index),
			Name:       lineArr[0],
			MusicType:  lineArr[1],
			MusicAdder: lineArr[2],
		})
	}

}

//使用IO
func ReadFileIoContent(file *os.File, m *[]model.Music) {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件失败")
	}
	fileContent := string(b)
	lineArr := strings.Split(fileContent, "\n")
	for i, line := range lineArr {
		cells := strings.Split(line, " ")
		*m = append(*m, model.Music{
			Id:         uint64(i),
			Name:       cells[0],
			MusicType:  cells[1],
			MusicAdder: cells[2],
		})
	}
}
