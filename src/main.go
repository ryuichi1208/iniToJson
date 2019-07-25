package main

import (
    "fmt"
    "os"
    "log"
)

func load_ini_file(filepath string) {
    file, err := os.Open(filepath)
    if err != nil {
	log.Fatal
    }

}

/* main function */
func main() {
    fmt.Printf("Start")
    load_ini_file("./test.ini")
}
