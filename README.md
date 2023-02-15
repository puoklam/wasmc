# wasmc
A ui package for web assembly in Go

## Test
```
cd wasm/
GOOS=js GOARCH=wasm go build -o ../assets/main.wasm

cd ../server/
go run main.go
```