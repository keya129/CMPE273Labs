package main
import ("fmt";"time")
func Sleep(interval int){

<-time.After(time.Second*time.Duration(interval))
}
func main(){

for i:=0;i<10;i++{
fmt.Println(i)

Sleep(1)

}

}
