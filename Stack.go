package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Stack struct{
	n int
	top int
	arr[10]int
}

func push(s *Stack, val int){
	if s.top >= s.n-1{
		fmt.Println("Stack Overflow")
	}else{
		s.top++
		s.arr[s.top]=val
	}

}
 
func pop(s *Stack) int{
	var r int
	if s.top <= -1{
		fmt.Println("Stack is empty")
		return -1
	}else{
		r = s.arr[s.top]
		s.arr[s.top]=0
		s.top--
	}
	return r
}

func display(s *Stack){
	for i:=0; i<=s.top; i++{
		fmt.Println(s.arr[i], "\t")
	}
}

func main(){
	reader := bufio.NewScanner(os.Stdin)
	S :=Stack{}
	fmt.Println("Enter size of stack")
	reader.Scan()
	nval,_:=strconv.ParseInt(reader.Text(), 10,64)
	S.n = int(nval)
	S.top = -1
	fmt.Println("1 Push\n2 Pop\n3 Display\n4 Exit")
	for {
		fmt.Println("Enter Choice")
		reader.Scan()
		choice,_:=strconv.ParseInt(reader.Text(), 10,64)
		switch choice{
		
	    case 1:
			fmt.Println("Enter no want to add")
			reader.Scan()
			val,_:=strconv.ParseInt(reader.Text(), 10,64)
			push(&S, int(val))

		case 2:
			a:= pop(&S)
			if a > -1{
				fmt.Println("Element Poped", a)
			}
		
		case 3:
			fmt.Println("Stack Values")
			display(&S)
			fmt.Println()
		
		case 4:
			fmt.Println("Exit main")
			return
		
		default:
			fmt.Println("Invalid Choice")

		}	
	}
}


