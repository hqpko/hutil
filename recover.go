package hutils

import (
	"log"
	"runtime/debug"
)

func Recover() {
	if r := recover(); r != nil {
		log.Printf("recoverd: %v", r)
		debug.PrintStack()
	}
}

func RecoverF(f func(r interface{})) {
	if r := recover(); r != nil {
		f(r)
	}
}
