package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main(){
	fmt.Println("hello world");
	
	var a int = 100;

	fmt.Println("hello world ",a);

	// Reading input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter String : ")
	
	// comma ok || err ok 
	input , _ := reader.ReadString('\n') //keep reading till the new line 
	// input or error 

	fmt.Println("The input is ",input);
	fmt.Printf("The type of the input is %T \n",input);
	//string conversion to float64
	num,_ := strconv.ParseFloat(input , 64)
	fmt.Printf("The type after parsing is %T \n",num);

	//time in go 
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006"))

	//Pointer 
	//pointer storing the integer 
	var ptr *int;
	fmt.Println("The default value of the pointer is ",ptr);
	myNum := 56;
	ptr = &myNum; //& will accessing the memory address (reference)
	
	fmt.Println("The value of ",ptr);
	fmt.Println("The actual value is ",*ptr);

	*ptr = *ptr + 2;
	fmt.Println("The new value is ",myNum);

	//Arrays

	var fruitList [5]string;
	fruitList[2] = "Apple";

	fmt.Println("The fruit at index 2 is ",fruitList[2]);
	fmt.Println("The fruit at index 1 is ",fruitList[1]);

	var newList = [3]string{"item1","item2","item3"};

	fmt.Println("The length of array is ",len(newList));
	fmt.Println("The fruits are ",fruitList);


}