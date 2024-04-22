package main

import(
	"fmt"
	"github.com/onodaryo/practice/go/util"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <year>")
		os.Exit(1)
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("コマンドライン引数は整数である必要があります.")
		os.Exit(1)
	}
	if !(util.IsLeapYear(year)) {
		fmt.Println(fmt.Sprintf("%s 年は閏年ではありません.", strconv.Itoa(year)))
	} else {
		fmt.Println(fmt.Sprintf("%s 年は閏年です.", strconv.Itoa(year)))
	}
}
