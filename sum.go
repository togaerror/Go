package main
import (
	"fmt"
)
func main() {
	var n, m int = 1, 10
	fmt.Println(sumOfInt(n, m))
	fmt.Println(sumOfSquare(n, m))
	fmt.Println(sumOfCube(n, m))
}
func sumOfInt(n, m int) int {
	sum := 0
	for ; n <= m; n++ {
		sum += n
	}
	return sum
}
func sumOfSquare(n, m int) int {
	sum := 0
	for ; n <= m; n++ {
		sum += n * n
	}
	return sum
}
func sumOfCube(n, m int) int {
	sum := 0
	for ; n <= m; n++ {
		sum += n * n * n
	}
	return sum
}
