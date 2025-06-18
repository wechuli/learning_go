package main
import(
	"fmt"
	"time"
)

func MyPrint(start int,finish int){
	for i:=start;i <= finish;i++{
		fmt.Print(i," ")

	}
	fmt.Println()
	time.Sleep(100*time.Microsecond)
}