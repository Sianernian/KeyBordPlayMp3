package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

import (
	_ "github.com/codyguo/godaemon"
	termbox "github.com/nsf/termbox-go"
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
				Play1()
				fmt.Println("You press F1")
			case termbox.KeyArrowUp:
				fmt.Println("You press 👆")
			default:
				break Loop

			}

		}

	}
}

func Play1() {
	// 1. 打开mp3文件
	audioFile, err := os.Open("./keybord/sound/20210512-00185381 00_00_01-00_00_02.mp3")
	if err != nil {
		log.Fatal(err)
	}
	// 使用defer防止文件描述服忘记关闭导致资源泄露
	defer audioFile.Close()
	// 对文件进行解码
	audioStreamer, format, err := mp3.Decode(audioFile)
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
