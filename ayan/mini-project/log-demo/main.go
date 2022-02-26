package main

import (
	logger "logger/logger"
)

func main() {

	l := logger.Logger{Filepath: "."}

	go test1(l)

	go test2(l)

	l.Log("main() : Log 1")

	l.Log("main() : Log 2")

	l.Log("main() : Log 3")

}

func test1(l logger.Logger) {

	l.Log("test1() : Log 1")

	l.Log("test1() : Log 2")

	l.Log("test1() : Log 3")
}

func test2(l logger.Logger) {

	l.Log("test2() : Log 1")

	l.Log("test2() : Log 2")

	l.Log("test2() : Log 3")
}
