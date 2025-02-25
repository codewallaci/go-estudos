package main

import "fmt"

//the main function
func main() {
    //print Hello World! to the console
    fmt.Println("Hello World!")

    //data types function
    data_types()
}

func data_types() {
    var a int = 1 //negative and positive numbers
    var b string = "test" //string
    var c float32 = 1.1 //float
    var d bool = true //boolean

    fmt.Println("int", a)
    fmt.Println("string", b)
    fmt.Println("float", c)
    fmt.Println("bool", d)

    //change var a value
    a = 2
    fmt.Println("int", a)

    //inferred type
    x := 10
    fmt.Println("inferred type", x)

    //variable without initial value
    var y int
    fmt.Println("var without init value", y)
    
    //set value to y
    y = 20
    fmt.Println("var y with value", y)

    //sum y + x
    sum := y + x
    fmt.Println("sum y + x", sum)

    //multiple vars in line
    var e, f, g int = 1, 2, 3

    fmt.Println("multiple vars", e, f, g)

    //multiple vars in a block
    var (
        h int
        i int = 1
        j string = "test"
    )

    fmt.Println("multiple vars in block", h, i, j)
}


