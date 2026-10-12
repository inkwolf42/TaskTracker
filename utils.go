package main

import "os"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfArgsNumNotEqual(n int) {
	if len(os.Args) != n {
		panic("Invalide Parameters!!!")
	}
}
