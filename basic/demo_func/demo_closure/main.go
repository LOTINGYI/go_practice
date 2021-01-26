package main
import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if !strings.HasSuffix(name,suffix) {
			return name+suffix
		}
		return name
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println("文件名處理後= ",f("winter.jpg"))
}