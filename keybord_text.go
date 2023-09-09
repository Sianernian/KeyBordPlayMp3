package main

import (
	"container/list"
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
	//使用通道获取一系列击键：
	KeysEvents, err := keyboard.GetKeys(1024)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer keyboard.Close()
	a := Mp3File{}
	command := list.New()
	fmt.Println("Press ESC to quit")
	for {
		event := <-KeysEvents
		if event.Err != nil {
			fmt.Println(event.Err)
			return
		}
		command.PushBack(string(event.Rune))
		first := command.Front()
		fmt.Printf("You pressed:string %s, key %X\r\n", string(event.Rune), event.Key)
		for i := 0; i <= command.Len(); i++ { // 类似java 的迭代器
			fmt.Println("first.Next():", first.Next())

			if first.Next() != nil {
				a.Stop()
				for e := command.Front(); e != nil; e = e.Next() {
					command.Remove(e) // 清除所有命令
				}
				fmt.Printf("command:%+v\n", command)
			}
		}

		a.OpenMp3File()
		go a.PlayMp3() // 启动协程  使用chan 传递信息

		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}

type Mp3File struct {
	Filestream    *os.File              // 文件流
	done          chan bool             // 结束信号
	paused        chan bool             // 暂停标志
	ctrl          *beep.Ctrl            // 控制器
	audioStreamer beep.StreamSeekCloser // 流信息
	format        beep.Format           // 文件信息
	progress      float64               // 进度值

}

func (m *Mp3File) OpenMp3File() {
	// 1. 打开mp3文件
	var err error
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano()) // 设置随机种子数

	a := strconv.Itoa(rand.Intn(30))
	fmt.Printf("播放第%s首歌\n", a)
	m.Filestream, err = os.Open("./keybord/sound/" + a + ".mp3")
	if err != nil {
		log.Fatal(err)
	}
	// 使用defer防止文件描述服忘记关闭导致资源泄露
	//defer m.Filestream.Close()

}

func (m *Mp3File) PlayMp3() {
	// 对文件进行解码
	var err error
	m.audioStreamer, m.format, err = mp3.Decode(m.Filestream)
	if err != nil {
		log.Fatal(err)
	}

	defer m.audioStreamer.Close()

	// SampleRate is the number of samples per second. 采样率
	_ = speaker.Init(m.format.SampleRate, m.format.SampleRate.N(time.Second/10))

	// 用于数据同步，当播放完毕的时候，回调函数中通过chan通知主goroutine
	m.done = make(chan bool)
	m.paused = make(chan bool)
	//startTime := time.Now()
	m.ctrl = &beep.Ctrl{Streamer: beep.Seq(m.audioStreamer, beep.Callback(func() {
		m.done <- true
	})), Paused: false}
	// 这里播放音乐
	//speaker.Play(beep.Seq(audioStreamer, beep.Callback(func() {
	//	// 播放完成调用回调函数
	//	m.done <- true
	//})))
	// 等待播放完成
	speaker.Play(m.ctrl)
	//<-m.done
	for { //延时处理
		select {
		case <-m.done:
			// 此处必须调用，否则下次Init会有死锁
			speaker.Clear()
			return
		case value := <-m.paused:
			speaker.Lock()
			m.ctrl.Paused = value
			speaker.Unlock()
		case <-time.After(time.Second):
			speaker.Lock()
			m.progress = float64(m.audioStreamer.Position())
			fmt.Println(m.progress)
			speaker.Unlock()
			//done <- true //写入
		}
	}
	//endTime := time.Since(startTime) / time.Millisecond // ms
	//fmt.Printf("music finished in %dms\n", endTime)

}

func (m *Mp3File) Stop() {
	select {
	case m.done <- true:
	default:
	}
}

func DangeKey() {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("You pressed :%s\n", string(char))

	//if err := keyboard.Open(); err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	_ = keyboard.Close()
	//}()
	//
	//fmt.Println("Press ESC to quit")
	//for {
	//	char, key, err := keyboard.GetKey()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
	//	if key == keyboard.KeyEsc {
	//		break
	//	}
	//}
}
