package utils

import (
	"strconv"
	"strings"
	"time"
)

func LoadBar() {
	var unloaded string = "░"
	var loaded string = "▓"
	var clear string = "\r\033[K"
	for i := 0; i < 8; i++ {
		print(clear)
		print("[")
		for j := 0; j < i; j++ {
			print(strings.Repeat(loaded, 7))
		}
		print(strings.Repeat(unloaded, 7-i))
		print("]")
		print(" " + strconv.Itoa(i) + "%")
		time.Sleep(time.Second * 1)
	}
}
