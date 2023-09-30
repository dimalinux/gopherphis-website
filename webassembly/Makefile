# -no-debug -scheduler=none

.PHONY: build
build:
	tinygo build -o _web/gopherphis.wasm -target wasm ./main.go
	cp "$$(tinygo env TINYGOROOT)/targets/wasm_exec.js" _web/
