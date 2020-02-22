package main


import (
"fmt"

)


const (
	host     = "localhost"
	port     = 5432
	user     = "rssuser"
	password = "password"
	dbname   = "rssfeeds"
) 

func main(){

fmt.Println("Starting program")	
app := App{}
app.Initialize(host,port,user,password,dbname)


}