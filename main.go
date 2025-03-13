package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main (){
	fmt.Println("Open Calc wrapper. Uses cmd to open calc waits 30 seconds and ends this program.")
	cmd := exec.Command("calc.exe")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error opening calc.exe")
	}
	time.Sleep(30 * time.Second)
}