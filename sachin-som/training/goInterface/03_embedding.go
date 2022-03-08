package main

type ReadWriter interface {
	Read(b buffer) bool
	Write(b buffer) bool
}

type Locker interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWriter
	Locker
	Close()
}
