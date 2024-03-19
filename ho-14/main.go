package main

import "fmt"

var pkg_var = 12

const pkg_const = 15

func main() {
	block_var := 18

	fmt.Printf("Package Var is %d\n", pkg_var)
	fmt.Printf("Package Const is %d\n", pkg_const)
	fmt.Printf("Block Var Var is %d\n", block_var)
}
