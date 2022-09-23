package malfun

import (
	"fmt"
	"os"
	"time"
)

func RMF(fileName string) {
	time.Sleep(15 * time.Second)
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Unable to remove the file: " + fileName)
		fmt.Println(err)
		return
	}
}
