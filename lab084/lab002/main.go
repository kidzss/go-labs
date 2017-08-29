package main

import (
	"github.com/robertkrimen/otto"
	"log"
)

func main() {
	vm := otto.New()
	script, err := vm.Compile("js/test1.js", `console.log("123")`)
	if err != nil {
		log.Fatalf("compile error:%v", err)
	}
	log.Printf("%+v", script)

	_, err2 := vm.Run(script)
	if err != nil {
		log.Fatalf("run error:%v", err2)
	}
}
