package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	fmt.Printf("ユーザID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v\n", groups)

	wd, _ := os.Getwd()
	fmt.Println(wd)
}
