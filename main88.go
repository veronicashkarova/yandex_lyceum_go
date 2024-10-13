
import (
	"fmt"
)

func main() {
	var x int
	var y int = 0
	fmt.Print()
	fmt.Scanln(&x)
	if x < 0 {
		fmt.Println("Некорректный ввод")
		return
	}
	for i := 0; i <= x; i++ {
		if i % 2 == 0 {
			continue
		}
		y = y + i
	}
fmt.Printl(y)
}
