package main

import (
	"fmt"
	"hw3/sys_lib"
)

func main() {
	new_engineer := sys_lib.Init_input()
	fmt.Printf("Initiliasation is finished. You are %s, with %d years experience", new_engineer.Name, new_engineer.Experience)

}
