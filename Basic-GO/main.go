package main
import "fmt"

func main() {
	// Variable Define and Common Data Type in golang

	var userName string  //String Data Type
	var userAge int // Int data Type use for Store number
	const conferenceTickets int = 50 // Const Data Type it can't be changed 
	var isAlive = true
	var accountBalance float64



	fmt.Println("Give you Name: ")
    fmt.Scanln(&userName)
	fmt.Printf("Hello %v..! Your age is %d You not dead %t, you have %v balance \n", userName,userAge,isAlive,accountBalance)

	// Array 

	// Array String type and store values there 
	var arrayStr[4]string
	arrayStr[0] = "sonu"
	fmt.Println(len(arrayStr))
	fmt.Println(arrayStr)
    
	// Array with integer 
	var arrayInt[4]int
	arrayInt[0] = 2
	fmt.Println(arrayInt)

    // Array with Bool
	var arrayBool[4]bool
	arrayBool[1]=true
	fmt.Println(arrayBool)

	// 2D array
	 array2D := [2][2]int{{2,3},{4,5}}
	fmt.Println(array2D)

	// Slices Dynamic Array

	var emptySli[] string
	fmt.Println(len(emptySli))
}

