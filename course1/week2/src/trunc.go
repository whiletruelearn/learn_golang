package main
import (
	"fmt"
)

func main(){
	var number float32
	fmt.Printf("Enter the floating point number => \n")
	fmt.Scan(&number)
	new_num := int32(number)
	fmt.Printf("Entered number is  => %f \n",number)
	fmt.Printf("Truncated number is => %d \n",new_num)
	
	
}