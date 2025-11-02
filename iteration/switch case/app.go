package main

import(
	"fmt"
	"runtime"
)

func main(){
	switch os:= runtime.GOOS; os{
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	case "darwin":
		fmt.Println("macOS.")
	default:
		fmt.Printf("%s", os)
	}
}