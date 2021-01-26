package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	filePath := "./test01.txt"
	file,err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND,0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str :="hello,Garden08\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	/* 
		因為writer是帶緩存，因此調用WriterString方法時，
		內容事先寫入到緩存忠，所以需要調用Flush()，將緩
		存的數據真正寫入到文件中，否則文建會沒有數據!
		*/
	writer.Flush()

}