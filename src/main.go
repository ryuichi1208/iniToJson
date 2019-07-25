package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func load_ini_file(filepath string) {
    file, err := os.Open(filepath)
    if err != nil {
	fmt.Println("Open Error")
    }

    b, err := ioutil.ReadAll(file)
    fmt.Println(string(b))
}

/* main function */
func main() {
    fmt.Println("Start")
    load_ini_file("./test.ini")
}
