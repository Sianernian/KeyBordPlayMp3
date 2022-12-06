package main

import (
	"fmt"
	_ "github.com/codyguo/godaemon"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	termbox "github.com/nsf/termbox-go"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				fmt.Println("You press Esc")
			case termbox.KeyCtrlA:
				a := Mp3{}
				a.OpenMp3()
				a.Play()
				fmt.Println("You press F1")
			case termbox.KeyArrowUp:
				fmt.Println("You press 👆")
			default:
				break Loop

			}

		}

	}
}

type Mp3 struct {
	Filestream *os.File // 文件流
}

func (m *Mp3) OpenMp3() {
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

func (m *Mp3) Play() {
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
