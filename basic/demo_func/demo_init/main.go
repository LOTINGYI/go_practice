package main
import(
	"fmt"
	"go_code/basic/demo_func/demo_init/utils"
)
var age = test()
func test() int {
	fmt.Println("test() ...")
	return 90
}

func init()  {
	fmt.Println("init() ...")
}

func main() {
	fmt.Println("Age= ",utils.Age, " Name= ",utils.Name)
	fmt.Println("main() ... age= ",age)	
}