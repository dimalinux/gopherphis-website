//go:build js && wasm

package main

import (
	"encoding/hex"
	"errors"
	"strings"
	"syscall/js"

	"github.com/dimalinux/gopherphis/mnemonic"
	"github.com/dimalinux/gopherphis/polyseed"
)

type mnemonicType int

const (
	mnemonicUnknown mnemonicType = iota
	mnemonicCryptonote
	mnemonicPolyseed

	selectedValueVarName  = "selectedMnemonic"
	mnemonicCryptonoteStr = "cryptonote_mnemonic"
	mnemonicPolyseedStr   = "polyseed_mnemonic"
)

func getMnemonicType() mnemonicType {
	mTypeStr := js.Global().Get(selectedValueVarName).String()
	switch mTypeStr {
	case mnemonicCryptonoteStr:
		return mnemonicCryptonote
	case mnemonicPolyseedStr:
		return mnemonicPolyseed
	default:
		return mnemonicUnknown
	}
}

//go:wasm-module gopherphis
//export updateKeyFromSeeds
func updateKeyFromSeeds() {
	phrase := js.Global().Get("document").Call("getElementById", "phrase").Get("value").String()
	seedsSplit := strings.Split(phrase, " ")

	var key []byte
	err := errors.New("unknown mnemonic type")
	mType := getMnemonicType()

	switch mType {
	case mnemonicCryptonote:
		key, err = mnemonic.CreateKeyFromSeeds(seedsSplit)
	case mnemonicPolyseed:
		var seedData *polyseed.SeedData
		seedData, err = polyseed.CreateSeedData(seedsSplit)
		if err == nil {
			key = seedData.KeyGen()
		}
	}

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
