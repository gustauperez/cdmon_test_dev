package main

import (
    "fmt"
    "strconv"
)

func main(){
    for i := 1; i <= 100; i++ {
        output := ""
        if(i%3==0) {output="CD"}
        if(i%5==0) {output=output+"mon"}
        if(output == "") {output=strconv.Itoa(i)}
        fmt.Println(output);
	}
}
