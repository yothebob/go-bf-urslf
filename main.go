package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
)

func interpret_bf(){
	
  file, err := os.Open("data.txt")
  if err != nil{
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    fmt.Println(scanner.Text())
  }
}


func main(){
  read_file()
}


how 
