package main

import (
	"fmt"
	"os"
)


func writeFile(fileName string,fileContent string) int {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil{
		fmt.Println("Error opening the file: ",err)
		return 0
	}

	defer file.Close()

	writeData := []byte(fileContent)
	_,err = file.Write(writeData)

	if err != nil{
		fmt.Println("Error opening the file: ",err)
		return 0
	}

	return 1
}

func readFile(fileName string, bufferSize int) (int,string){
	file,err := os.OpenFile(fileName, os.O_RDONLY,0444)

	if err != nil{
		fmt.Println("Error reading the file: ", err)
		return -1,""
	}

	defer file.Close()


	buffer := make([]byte,bufferSize)
	dataSize, readErr := file.Read(buffer)

	if readErr != nil{
		fmt.Println("Error reading the file: ", readErr)
		return -1, ""
	}

	dataString := string(buffer[:dataSize])

	return 1,dataString


}


func main(){

	isWritten := writeFile("test.txt","Hello, how are you?")

	if isWritten != 1{
		return
	}

	isRead, readData := readFile("test.txt",10)

	if isRead != 1{
		return
	}	

	fmt.Println(readData)
}