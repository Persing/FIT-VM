/*
 * Author: Dmitry Kulakov, dkulakov2014@my.fit.edu
 * Author: Nicholas Persing, 
 * Course: CSE 4250, Spring 2016
 * Project: project tag, short project name
 */

package main

import (
    "fmt"
    "os"
    "io"
    //"bufio"
)

func main() {
	
	var registers [8]uint32 //8 general purpose registers
	registers[0] = 1 //so that the compiler stops whining
	 
	//var ef *uint32 = 1 //execution finger
	fileName := os.Args[1] //gets the filename from the command line
	data, err := os.Open(fileName)
	defer data.Close() //close file at the end of execution
	
	if err != nil {
	    panic(err)
	}
	
	buffer := make([]byte, 56500) //14091
	
	for {
	    n, err := data.Read(buffer)
	    if err != nil && err != io.EOF {
	        panic(err)
	    }
	    
	    if n == 0 {
	        break
	    }
	}
	
	fmt.Println(buffer)
}

