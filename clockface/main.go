package main

import (
	"clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.WriteSVG(os.Stdout, t)
}
