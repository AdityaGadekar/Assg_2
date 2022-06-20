package main
import "fmt"
type info struct{
	firstname string
	lastname string
}
func value(i info){}

func pointer(i *info){	
	i.firstname = " Pass by "
	i.lastname = "Pointer"
}

func main(){
	data:= info{
		firstname : "Pass by",
		lastname : "Value" ,
	}
	data1:= info{}

	//Pass by value
	fmt.Println(data)

	//Pass by reference
	pointer(&data1)
	fmt.Println(data1)
}