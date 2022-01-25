package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "bufio"
)

func main() {
    var text string
    fmt.Println("Please enter your text here")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text = scanner.Text()
    if err := write("myfile.txt", text); err != nil {
        log.Fatal("failed to write file:", err)
    }
}

func write(fileName string, text string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = io.WriteString(file, text)
    if err != nil {
        return err
    }
    return nil
}

