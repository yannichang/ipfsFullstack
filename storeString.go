package main

import (
	"fmt"
	"bytes"
	"strings"
	shell "github.com/ipfs/go-ipfs-api"

)
func main() {
	sh := shell.NewShell("localhost:5001")
	//Store a string in IPFS
	cid, err := sh.Add(strings.NewReader("This is a test."))
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("added the given string `This is a test`, the CID is", cid)
	//Retrieve the string using its CID
	resp, err := sh.Cat(cid)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp)
	//the return of Cat is io.ReaderCloser, convert to string for output
	newStr := buf.String()
	fmt.Println("Retrieving the string: \n",newStr)
}
