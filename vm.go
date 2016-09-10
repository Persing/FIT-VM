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
    "time"
)

const MOD6 uint32 = 63
const MOD9 uint32 = 511
const MOD25 uint32 = 33554431
const MAXUINT uint32 = 4294967295 //2^n-1

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

	platterCollection[0] = zeroArray
	var counter uint32 = 1 //used to allocate addresses in the platterCollection
	//channel := make(chan uint32, 255) //channel for deallocated addresses
	
	queue := make([]uint32, 0, 1000)
	
	for {
	    platter := platterCollection[0][ef]
	    ef++
	    
	    operation := getOp(platter)
	    //obtains the last 9 digits of platter, where register indices are
	    A := (platter >> 6) & 7
	    B := (platter >> 3) & 7
	    C := platter & 7
	    
	    reader := bufio.NewReader(os.Stdin)
	    //var input byte
	    start := time.Now()
	    
	    switch operation {
	        case 0:
	            //fmt.Println("0")
	            if (registers[C] != 0) {
	                registers[A] = registers[B]
	            } 
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 0: ", time.Since(start))
	                }
	        case 1:
//		        fmt.Println(A)
//		        fmt.Println(B)
                //fmt.Println("1")
                registers[A] = platterCollection[registers[B]][registers[C]]
                if time.Since(start) > 10000 {
                    fmt.Println("Duration 1: ", time.Since(start))
                    }
	        case 2:
		        //fmt.Println("2")
	            //fmt.Println(registers[A])
	            //fmt.Println(platterCollection[registers[A]])
	            platterCollection[registers[A]][registers[B]] = registers[C]
	            //fmt.Println(platterCollection[registers[A]])
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 2: ", time.Since(start))
	                }
	        case 3:
		        //fmt.Println("3")
	            registers[A] = (registers[B] + registers[C]) //% MAXUINT
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 3: ", time.Since(start))
	                }
	        case 4:
		        //fmt.Println("4")
	            registers[A] = (registers[B] * registers[C]) //% MAXUINT
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 4: ", time.Since(start))
	                }
	        case 5:
		        //fmt.Println("5")
	            /*if registers[C] == 0 {
	                fmt.Println("Divide by zero. System fail")
	                System.exit(3)
	            }*/
	            registers[A] = (registers[B] / registers[C])
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 5: ", time.Since(start))
	                }
	        case 6:
	        //fmt.Println("6")
	            registers[A] = ^(registers[B] & registers[C]) //bitwise not and
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 6: ", time.Since(start))
	                
	                }
	        case 7:
	        //fmt.Println("7")
	            os.Exit(3)
	        case 8:
	        //fmt.Println("8")
		        //for{
		        	address := counter
		        	
		        	if len(queue) > 0 {
		        	    
		        	    address = queue[0]
		        	    queue = queue[1:]
		        	} else {
		        	    counter++
		        	}
			         //_, ok := platterCollection[counter]
		            //if !ok {
		            	platterCollection[address] = make([] uint32, registers[C])
		            	registers[B] = address
		            	//break
		            //}
		        //}
		        
//	            address := counter
//	            select {
//	                case address = <- channel:
//	                    //fmt.Println("channeled")
//	                default:
//	                    //fmt.Println("default")
//	                    counter++
//	            }
//                //fmt.Println(address)
//	            platterCollection[address] = make([] uint32, registers[C])
//	            registers[B] = address
                if time.Since(start) > 10000 {
                    fmt.Println("Duration 8: ", time.Since(start))
                    }

	        case 9:
	            //check if key value exists
	           // _, ok := platterCollection[registers[C]]
	           //if ok {
	                delete(platterCollection, registers[C])
                    queue = append(queue, registers[C])
//	                
//	                if (registers[C] == 0) {
//	                    fmt.Println("Abandoned 0 array. System fail")
//	                    os.Exit(3) 
//	                }
	            //}
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 9: ", time.Since(start))
	                }
	        case 10:
		        //fmt.Println("10")
	            fmt.Print(string(registers[C]))
	        case 11:
		        //fmt.Println("11")
	            input, err := reader.ReadByte()
	            if err != nil && err != io.EOF {
	                panic(err)
	            }
	            
	            registers[C] = uint32(input)
	            
	            if (err == io.EOF) {
	                registers[C] = MAXUINT //make it pregnant with bits
	            }
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 11: ", time.Since(start))
	                }
	        case 12:
		        //fmt.Println("12")
//		        copy(platterCollection[0], platterCollection[registers[B]])
//		        if len(platterCollection[0]) == len(platterCollection[registers[B]]){
//		        	fmt.Println("sheeet son")
//		        }
//		        ef = registers[C] 
	            //if _, ok := platterCollection[registers[B]]; ok {
	            if registers[B] != 0 {    
	                originArray := platterCollection[registers[B]]
	                copiedArray := make([]uint32, len(originArray))
	                copy(copiedArray, originArray)
	                platterCollection[0] = copiedArray
	            }
	            
	            ef = registers[C]
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 12: ", time.Since(start))
	            }
	        case 13:
		        //fmt.Println("13")
	            registers[(platter >> 25) & 7] = platter & MOD25
	            if time.Since(start) > 10000 {
	                fmt.Println("Duration 13: ", time.Since(start))
	                }
	            
	        default:	            
	            fmt.Println("Not a valid operator. System fail")
	            os.Exit(3)
	    }
	}	
}
