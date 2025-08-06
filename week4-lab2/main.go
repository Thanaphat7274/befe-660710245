package main
import (
	"fmt"
)
// var emil string ="Thanaphat@su.ac.th"
func main(){
	// var name string = "Thanphat" 
	var age int = 20

	email := "sukchuen_t@su.ac.th"
	gpa := 4.00

	firstname, lastname := "Thanaphat","sukchuen"

	fmt.Printf("Name %s %s,age %d,email %s,gpa %.2f\n",firstname,lastname,age,email,gpa)
}