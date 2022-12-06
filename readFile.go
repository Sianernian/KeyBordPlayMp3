package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	//path := "./keybord/sound"
	//ReadMp3File(path)

}

func ReadMp3File(Path string) {
	fileInfo, err := ioutil.ReadDir(Path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	count := 0
	fmt.Println(fileInfo)
	for _, info := range fileInfo {
		count++
		a := strconv.Itoa(count)
		os.Rename(Path+"\\"+info.Name(), Path+"\\"+a+".mp3")
		fmt.Println(info.Name())
	}
	fmt.Println(count)

}
