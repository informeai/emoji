# :tada: Emoji
Simple emojis package write with golang.
### :fire: Install
Get package
```
go get github.com/informeai/emoji
```
### :pencil2: Usage
```
package main

import(
	"fmt"
	"github.com/informeai/emoji"
)

func main(){
	fmt.Println("Hello Emoji.")
	emoji.Println(":fire: is fire this package.")
	heart := emoji.Sprint("I love :heart:")
	fmt.Println(heart)
}
```
by wellington gadelha :heart:.

