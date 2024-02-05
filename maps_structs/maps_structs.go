package main

import (
	"fmt"
);

func main(){
	//maps 
	fmt.Println("Maps in golang");

	newMap := make(map[string]int)

	newMap["first"] = 1;
	newMap["second"] = 2;
	newMap["third"] = 3;

	fmt.Println("The map is : ",newMap)
	fmt.Println("The value corresponding to first key is : ",newMap["first"]);

	//deleting from map 

	delete(newMap,"first");

	fmt.Println("After deleting first : ",newMap);

	//iterating 
	for i,j := range newMap{
		fmt.Println(i," ",j);
	}

	//structs 
	//no inheritence

	type User struct{
		Name string
		Email string
		Status string
		Age int
	}

	user1 := User{"user1","user1@gmail.com","online",20};

	fmt.Println(user1)
	fmt.Printf("The name of user is %v",user1.Name);
}