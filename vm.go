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
    "encoding/binary"
    //"bufio"
)

const MOD25 uint32 = 33554432

func getOp (platter uint32) uint32 {
    return platter >> 28
}

func main() {
	
	platterCollection := make(map[uint32] []uint32)
	var registers [8]uint32 //8 general purpose registers
	//registers[0] = 1 //so that the compiler stops whining
	 
	var ef uint32 = 0 //execution finger
	fileName := os.Args[1] //gets the filename from the command line
	data, err := os.Open(fileName)
	defer data.Close() //close file at the end of execution
	
	if err != nil {
	    panic(err)
	}
	
	//buffer := make([]byte, 56364) //14091 words * 4 = 56364 bytes
	buffer := make([]byte, 4)
	zeroArray := make([] uint32, 0, 14091)
	
	for {
	    n, err := data.Read(buffer)
	    if err != nil && err != io.EOF {
	        panic(err)
	    }
	    
	    if n == 0 {
	        break
	    }

	    zeroArray = append(zeroArray, binary.BigEndian.Uint32(buffer))
	}
	
	platterCollection[0] = zeroArray
	fmt.Println(zeroArray)
	
	for {
	    platter := zeroArray[ef]
	    ef++
	    
	    operation := getOp(platter)
	    A := (platter >> 6) % 7
	    B := (platter >> 3) % 7
	    C := platter % 7
	    
	    switch operation {
	        case 0:
	            if (registers[C] != 0) {
	                registers[A] = registers[B]
	            }  
	        case 1:
                //registers[A] =
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	            return
	        case 8:
	        case 9:
	        case 10:
	        case 11:
	        case 12:
	        case 13:
	            registers[(platter >> 25) % 7] = platter % MOD25
	            
	        //case default:
	        //    fmt.Println("fuck you")
	    }
	}
	
	
}

