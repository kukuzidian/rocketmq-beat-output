package rocketmq

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFilebeat(t *testing.T) {

	var strbuf strings.Builder
	var constr string

	for k:=1;k<10;k++ {
		strbuf.WriteString("xxxxx")
	}
	constr = strbuf.String()

	file , err := os.Create("/Users/zhangjungang/test/c1.txt")
	if err != nil {
		fmt.Println("file create err")
		return
	}

	bw := bufio.NewWriter(file)
	for i := 1 ; i< 300;i++ {
		bw.WriteString(strconv.Itoa(i)+"|"+constr+"\n")
		bw.Flush()
		time.Sleep(time.Duration(1)*time.Second)
		fmt.Println("curtime=",time.Now())
	}
	defer file.Close()

}

func TestTime(t *testing.T)  {

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(strconv.FormatInt(time.Now().UnixNano()/1000,10))

}