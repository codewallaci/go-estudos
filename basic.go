package main

import "fmt"

//the main function
func main() {
    //print Hello World! to the console
    fmt.Println("Hello World!")

    //data types function
    // data_types()

    //go outputs
    // outputs()

    //array
    array()
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

    const PI float32 = 3.14
    fmt.Println("const PI", PI)
}

func outputs() {
    var a, b string = "hello", "world"


    fmt.Print(a)
    fmt.Print(b)

    fmt.Print("\n")

    //break line
    fmt.Print(a, "\n", b, "\n")

    //blank space
    fmt.Print(a, " ", b, "\n")

    //printf function
    fmt.Printf("a has a value: %v and type: %T\n", a, a)
    fmt.Printf("b has a value: %v and type %T\n", b, b)

    //print the value in go-syntax format
    fmt.Printf("a has a value: %#v and type: %T\n", a, a)
    
}

func array () {
    

    //array with length 2
    var array_1 = [2]string{"ford", "fiat"}
    fmt.Println(array_1)

    //array witn inferred length
    var array_2 = [...]string{"ford", "fiat", "bmw"}
    fmt.Println(array_2)

    //sign array
    array_3 := [4]string{"ford", "fiat", "bmw", "chevrolet"}

    //acess array element
    fmt.Println(array_3[1])

    //change array element
    array_3[1] = "toyota"
    fmt.Println(array_3[1])

    //array len
    fmt.Println(len(array_3))


    //arra Initialization

    //not initialized
    arr1 := [5]int{}
    
    //partially initialized
    arr2 := [5]int{1, 2}

    //fully initialized
    arr3 := [5]int{1, 2, 3, 4, 5}

    fmt.Println("Array 1", arr1)
    fmt.Println("Array 2", arr2)
    fmt.Println("Array 3", arr3)

    //initialize specific index
    arr4 := [5]int{1:2, 4:10}
    fmt.Println("Array 4", arr4)
    
}