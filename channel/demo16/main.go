package main

import (
	"context"
	"fmt"
)

func main() {
	arr := []int{1,2,3,4}

	length := len(arr)
	fmt.Println(float64(arr[length/2 - 1] + arr[length/2]) /2)

}

func testContext(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
}
