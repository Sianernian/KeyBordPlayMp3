package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Press any key to exit...\n")
	b := make([]byte, 1)
	os.Stdin.Read(b)
	//path := "./keybord/sound"
	//ReadMp3File(path)

}

func ReadMp3File(Path string) {
	fileInfo, err := ioutil.ReadDir(Path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	count := -1
	fmt.Println(fileInfo)
	for _, info := range fileInfo {
		count++
		//os.Rename(Path+"\\"+info.Name(), Path+"\\"+info.Name()[:17]+".mp3")
		os.Rename(Path+"\\"+info.Name(), Path+"\\"+strconv.Itoa(count)+".mp3")
		fmt.Println(info.Name())
	}

}
