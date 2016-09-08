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
    //"container/heap"
    "bufio"
)

const MOD25 uint32 = 33554432
const MAXUINT uint32 = 4294967295 //-1

/*type Address struct {
    index uint32
    priority int //deallocated addresses have higher priorities
}

type addressQueue []*Address

func (queue addressQueue) Less(i, j int) bool {
    //return the address with the higher priority
    return queue[i].priority > pq[j].priority
}*/

func getOp (platter uint32) uint32 {
    return platter >> 28 
    //get the last 4 bits of platter to determine the operation
}

func main() {
	
	platterCollection := make(map[uint32] []uint32)
	var registers [8]uint32 //8 general purpose registers
	 
	var ef uint32 = 0 //execution finger
	fileName := os.Args[1] //gets the filename from the command line
	data, err := os.Open(fileName)
	defer data.Close() //close file at the end of execution
	
	if err != nil {
	    panic(err)
	}
	
	//buffer := make([]byte, 56364) //14091 words * 4 = 56364 bytes
	buffer := make([]byte, 4) //buffer for reading in input
	zeroArray := make([] uint32, 0, 14091) //zero array to store the program
	
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
	
	//fmt.Println(zeroArray)
	platterCollection[0] = zeroArray
	var counter uint32 //used to allocate addresses in the platterCollection
	channel := make(chan uint32) //channel for deallocated addresses
	
	for {
	    platter := zeroArray[ef]
	    ef++
	    
	    operation := getOp(platter)
	    A := (platter >> 6) % 7
	    B := (platter >> 3) % 7
	    C := platter % 7
	    
	    reader := bufio.NewReader(os.Stdin)
	    //var input byte
	    
	    switch operation {
	        case 0:
	            if (registers[C] != 0) {
	                registers[A] = registers[B]
	            }  
	        case 1:
                registers[A] = platterCollection[registers[B]][registers[C]]
	        case 2:
	            platterCollection[registers[A]][registers[B]] = registers[C]
	        case 3:
	            registers[A] = (registers[B] + registers[C]) % MAXUINT
	        case 4:
	            registers[A] = (registers[B] * registers[C]) % MAXUINT
	        case 5:
	            registers[A] = (registers[B] / registers[C])
	        case 6:
	            registers[A] = ^(registers[B] & registers[C]) //bitwise not and
	        case 7:
	            os.Exit(3)
	        case 8:
	            address := counter
	            select {
	                case x, ok := <-channel:
	                    if ok {
	                        address = x
	                    }
	                default:
	                    address = counter
	                    counter++
	            }
	            
	            
	            platterCollection[address] = make([] uint32, registers[C])
	            registers[B] = address

	        case 9:
	            //check if key value exists
	            _, ok := platterCollection[registers[C]]
	            if ok {
	                delete(platterCollection, registers[C])
	                channel <- registers[C]
	            }
	        case 10:
	            fmt.Print(string(registers[C]))
	        case 11:
	            input, err := reader.ReadByte()
	            if err != nil && err != io.EOF {
	                panic(err)
	            }
	            
	            registers[C] = uint32(input)
	            
	            if (err == io.EOF) {
	                registers[C] = MAXUINT //make it pregnant with bits
	            }
	        case 12:
	            if _, ok := platterCollection[registers[C]]; ok {
	                ef = registers[C]
	                copy(platterCollection[0], platterCollection[registers[B]])
	            }
	        case 13:
	            registers[(platter >> 25) % 7] = platter % MOD25
	            
	        default:
	            fmt.Println("fuck you")
	    }
	}
	
	
}

