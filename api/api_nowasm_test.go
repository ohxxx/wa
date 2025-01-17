// 版权 @2022 凹语言 作者。保留所有权利。

package api_test

import (
	"fmt"
	"log"

	"github.com/wa-lang/wa/api"
)

func ExampleRunCode() {
	const code = `
		fn main() {
			println(add(40, 2))
		}

		fn add(a: i32, b: i32) => i32 {
			return a+b
		}
	`

	output, err := api.RunCode("hello.wa", code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(output))

	// Output:
	// 42
}
