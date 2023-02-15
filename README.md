# wasmc
A ui package for web assembly in Go

## Test

```
cd wasm/
GOOS=js GOARCH=wasm go build -o ../assets/main.wasm

cd ../server/
go run main.go
```

## Example

wasm/main.go
```go
div := NewElement(WithText("hello"))
h3 := NewElement(WithTag("h3"), WithText("Hello World"))
box := NewBox(
    300,
    200,
    WithText("box"),
    WithStyle(map[string]string{"background-color": "blue"}),
)
circle := NewCircle(
    100,
    WithStyle(map[string]string{"background-color": "green"}),
)
flex := NewFlex("row")
item1 := NewBox(
    200,
    100,
    WithStyle(map[string]string{"background-color": "red"}),
)
item2 := NewBox(
    300,
    100,
    WithStyle(map[string]string{"background-color": "yellow"}),
)
flex.Children().Append(item1, item2)

mainEle := NewElementFromJS(doc.Call("getElementById", "main"))
mainEle.Children().Append(div, h3, box, circle, flex)
```