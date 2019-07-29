package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "runtime"
    "github.com/mitchellh/go-ps"
)

/* Print setting info */
func PrintExecInfo() {
    pid := os.Getpid()
    pidInfo, _ := ps.FindProcess(pid)
    //fmt.Printf("%+v\n", pidInfo)
    fmt.Println("==========================")

    fmt.Printf("Proc : %s\n", pidInfo.Executable())
    fmt.Printf("OS   : %s\n", runtime.GOOS)
    fmt.Printf("PID  : %d\n", pidInfo.Pid())
    fmt.Printf("PPID : %d\n", pidInfo.PPid())

    fmt.Println("==========================")
}

/* Load config file */
func LoadIniConfigFile(filepath string) {
    file, err := os.Open(filepath)
    if err != nil {
	fmt.Println("Open Error")
    }

    b, err := ioutil.ReadAll(file)
    fmt.Println(string(b))
}

/* main function */
func main() {
    PrintExecInfo()
    fmt.Println("Start")
    LoadIniConfigFile("./test.ini")
}
