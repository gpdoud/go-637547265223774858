package main

import "fmt"


// Test for renaming the function  name 

func main() {                           
                 
  	fun()
                     
	
}

func fun() {                                              // <<<<< rename,15,6,15,6,newfunc,pass

	var fun string = "this is a variable"            

		fmt.Println(fun)
                  
	fmt.Println("this is just for fun")                                                        

}

