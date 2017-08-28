package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var n int

	// 遅い
	fmt.Scan(&n)
	fmt.Println(n)

	// 1行づつ読み込む
	fmt.Println("Line")
	s, t := nextLine(), nextLine()
	fmt.Println(s)
	fmt.Println(t)

	// スペース区切りで読み込む Text
	fmt.Println("Split Text")
	line := nextLine()
	lines := strings.Fields(line)
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

	// スペース区切りで読み込む Number
	fmt.Println("Split Number")
	num := nextLine()
	nums := strings.Fields(num)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	arr := make([]int, len(nums))
	var total int = 0
	for i:= 0; i < len(nums); i++ {
		arr[i], _ = strconv.Atoi(nums[i])
		total += arr[i]
	}
	fmt.Println("Out Put")
	for i := 0; i < len(nums); i++ {
		print(arr[i])
		if i == len(nums) - 1 {
			fmt.Println("")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println(total)
}
