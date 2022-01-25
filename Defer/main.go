package main

import (
"fmt"
)

func main() {
    // The defer act consider the tmp by the at calling time
    a := 7
    defer fmt.Printf("The value of a in first defer call %v\n",a)
    a = 45
    

    var b int = 12
    defer fmt.Printf("The value of b in second defer call %v\n",b)

}

