package piscine

//import "github.com/01-edu/z01"
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var i int
	var j int
	var k int
	array := make([]string, 0, 3)

	for i = 0; i <= 9; i++ {
		for j = i + 1; j <= 9; j++ {
			for k = j + 1; k <= 9; k++ {
				s1 := strconv.Itoa(i)
				s2 := strconv.Itoa(j)
				s3 := strconv.Itoa(k)
				array = append(array, s1+s2+s3)
			}
		}
	}
	output := "" + strings.Join(array, `, `) + ``
	fmt.Println(output)
}
