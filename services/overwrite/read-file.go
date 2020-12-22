package overwrite

import (
	"bufio"
	"fmt"
	"os"
)

func OverwriteData() string{
	//прочитать данные и записать все в строку
	file, err :=os.Open("hello.txt")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	//data:= make([]byte,64)
	//a:=""
	//for{
	//	n,err:=file.Read(data)
	//	if err!=nil{
	//		break
	//	}
	//	a+=string(data[:n])
	//}
	//return a
	r := bufio.NewReader(file)
	s, e := Readln(r)
	var a string
	for e == nil {
		a=a+s
		s,e = Readln(r)
	}
	return a
}

func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
}

