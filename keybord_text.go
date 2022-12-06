package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//使用通道获取一系列击键的示例：
	KeysEvents, err := keyboard.GetKeys(5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	for {
		event := <-KeysEvents
		if event.Err != nil {
			fmt.Println(event.Err)
			return
		}
		a := Mp3File{}
		a.OpenMp3File()
		a.PlayMp3()
		fmt.Printf("You pressed:string %s, key %X\r\n", string(event.Rune), event.Key)

		if event.Key == keyboard.KeyEsc {
			break
		}
	}

}

type Mp3File struct {
	Filestream *os.File // 文件流
}

func (m *Mp3File) OpenMp3File() {
	// 1. 打开mp3文件
	var err error
	rand.Seed(time.Now().UnixNano()) // 设置随机种子数

	a := strconv.Itoa(rand.Intn(31))
	fmt.Println(a)
	m.Filestream, err = os.Open("./keybord/sound/" + a + ".mp3")
	if err != nil {
		log.Fatal(err)
	}
	// 使用defer防止文件描述服忘记关闭导致资源泄露
	//defer m.Filestream.Close()

}

func (m *Mp3File) PlayMp3() {
	// 对文件进行解码
	audioStreamer, format, err := mp3.Decode(m.Filestream)
	if err != nil {
		log.Fatal(err)
	}

	defer audioStreamer.Close()

	// SampleRate is the number of samples per second. 采样率
	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// 用于数据同步，当播放完毕的时候，回调函数中通过chan通知主goroutine
	done := make(chan bool)
	// 这里播放音乐
	speaker.Play(beep.Seq(audioStreamer, beep.Callback(func() {
		// 播放完成调用回调函数
		done <- true
	})))
	// 等待播放完成
	<-done
}

func DangeKey() {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("You pressed :%s\n", string(char))
}
