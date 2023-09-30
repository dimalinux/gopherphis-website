# -no-debug -scheduler=none

wasm_dir=../static
wasm_exec_dir=../src/routes/keys

.PHONY: build
build:
	tinygo build -o "$(wasm_dir)"/gopherphis.wasm -target wasm ./main.go
	cp "$$(tinygo env TINYGOROOT)/targets/wasm_exec.js" "$(wasm_exec_dir)"
