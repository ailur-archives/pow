package main

import (
	"git.ailur.dev/ailur/pow-argon2/library"

	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Proof of work module online")
	resource := js.Global().Get("resource").String()
	difficulty := js.Global().Get("difficulty").Int()
	fmt.Println("Beginning PoW with difficulty", difficulty, "and resource", resource)
	result, err := library.PoW(uint64(difficulty), resource)
	if err != nil {
		fmt.Println("Error:", err)
		js.Global().Set("return", js.ValueOf(err.Error()))
		js.Global().Set("returnCode", js.ValueOf(1))
		js.Global().Call("WASMComplete")
	} else {
		fmt.Println("Result:", result)
		js.Global().Set("return", js.ValueOf(result))
		js.Global().Set("returnCode", js.ValueOf(0))
		js.Global().Call("WASMComplete")
	}
}
