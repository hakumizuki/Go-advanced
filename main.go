package main
import(
	"fmt"

	// アルファベット順にしておくといい
	"./mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
}