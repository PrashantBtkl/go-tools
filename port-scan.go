package main

import (
	"fmt"
	"net"
	"sync"
	"os"
)

func main() {

	var wg sync.WaitGroup
	var host=os.Args[1]

	for i:=1 ; i<=65535; i++{
		wg.Add(1)
		go func(j int){
			defer wg.Done()
			address:=fmt.Sprintf("%s:%d",host,j)
			conn, err := net.Dial("tcp",address)
			if err != nil{
	//			panic(err)
				return
			}
			conn.Close()
			fmt.Printf("%d open\n",j)
		}(i)
		//fmt.Printf("started scanning %d\n",i)
		wg.Wait()
	}

}
