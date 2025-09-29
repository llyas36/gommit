package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func main(){
	wd, _ := os.Getwd()
	fmt.Println(wd)
	b1 := filepath.Base(wd+"/file.txt")
	fmt.Println("......")
	fmt.Println(b1)
}

func handleDir(){
	wd, _ := os.Getwd()
	fmt.Printf("CURRENT PWD: %v\n", wd)
	current_filepath := filepath.Base(wd + "/file.txt")
	fmt.Println("....") 
	fmt.Printf("THE FILE: %v\n", current_filepath)

}
