package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	CurrentMouse()
	robotgo.KeyTap(`command`)

	robotgo.KeyTap(`e`, `command`)

	robotgo.KeyTap(`esc`, `shift`, `control`)

}

func CurrentMouse() {
	fmt.Println(robotgo.GetMousePos())
}
