package main
import (
	"fmt"
)
func main(){
	var num int=5647
	var sum int=0
	for num>0{
		sum=sum+num%10
		num=num/10

	}
	var total int=0
	if sum>9{
		for sum>0{
			total=total+sum%10
			sum=sum/10
		}
		fmt.Println(total)
	}else{
		fmt.Println(sum)
	}
	
	
}