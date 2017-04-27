package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	resultPositions, err := gocompany.GetPositions("2191023")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultPositions)
	os.Exit(0)
}
