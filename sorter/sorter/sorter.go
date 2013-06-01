package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"

var infile *string = flag.String("i","infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string)(values []int, err error){
    file, err := os.Open(infile)
    if err != nil{
        fmt.Println("Failed to open the input file ", infile)
        return
    }

    defer file.Close()


}


func main(){
    flag.Parse()

    if infile != nil{
        fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
    }

}

