package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}

	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: lexb file.lexb")
		os.Exit(1)
	}

	if !strings.HasSuffix(os.Args[1],".lexb") {
		fmt.Println("usage: lexb file.lexb")
		os.Exit(1)
	}

	fl,er := os.OpenFile(os.Args[1],os.O_RDONLY,0644)

	Check(er)

	scanner := bufio.NewScanner(fl)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "print:" {
			txt = ""
		} else if txt == "endl" {
			fmt.Println()
		} else {
			fmt.Print(txt+" ")
		}
	}

	defer fl.Close()
}
