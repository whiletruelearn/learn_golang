package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"log"
)
func main(){
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the input String=>\n")
	line, err := in.ReadString('\n')
	line = strings.TrimRight(line,"\r\n")
	if err != nil{
		log.Fatal(err)
	}
	lowerInpString := strings.ToLower(string(line))
	if  strings.HasPrefix(lowerInpString,"i") {
		if strings.ContainsAny(lowerInpString,"a") && strings.HasSuffix(lowerInpString, "n"){
			fmt.Printf("Found!")
		} else {
			fmt.Printf("Not found!")
		}
		
	} else {
		fmt.Printf("Not found!")
	}
}