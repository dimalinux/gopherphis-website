//go:build js && wasm

package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/dimalinux/gopherphis/mnemonic"
)

//go:wasm-module gopherphis
//export updateKeyFromSeeds
func updateKeyFromSeeds() {
	phrase := js.Global().Get("document").Call("getElementById", "phrase").Get("value").String()
	fmt.Printf("Input is %q\n", phrase)
	seedsSplit := strings.Split(phrase, " ")
	fmt.Printf("input seeds are %d: %v\n", len(seedsSplit), seedsSplit)
	key, err := mnemonic.CreateKeyFromSeeds(seedsSplit)
	var keyStr string
	if err != nil {
		keyStr = err.Error()
	} else {
		keyStr = hex.EncodeToString(key)
	}
	js.Global().Get("document").Call("getElementById", "key").Set("value", keyStr)
}

func main() {
	wait := make(chan struct{}, 0)
	<-wait
	println("main exiting")
}
