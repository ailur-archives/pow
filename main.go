package main

import (
	"fmt"
	"github.com/catalinc/hashcash"
	"syscall/js"
)

func main() {
	fmt.Println("HashCash online! (no, curious console-dweller, this isn't a cryptocurrency miner)")
	fmt.Println("Beginning proof of work (this may take a while)...")
	pow := hashcash.New(20, 16, "I love burgernotes!")
	stamp, err := pow.Mint("signup")
	if err != nil {
		js.Global().Set("returnVar", js.ValueOf(err.Error()))
		js.Global().Set("returnCode", js.ValueOf(2))
		fmt.Println("An error occurred whilst working:", err)
		js.Global().Call("WASMComplete")
	} else {
		js.Global().Set("returnVar", js.ValueOf(stamp))
		js.Global().Set("returnCode", js.ValueOf(0))
		fmt.Println("Proof of work completed successfully:", stamp)
		fmt.Println("Again, no, this isn't a Crypto miner. It's an anti-spam measure. I promise.")
		js.Global().Call("WASMComplete")
	}
}
