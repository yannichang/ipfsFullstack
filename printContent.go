package main

import (
	"fmt"
	"bytes"
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewShell("localhost:5001")
	exampleCIDStr := "QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9"
	resp, err := sh.Cat(exampleCIDStr)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp)
	//the return of Cat is io.ReaderCloser, convert to string for output
	newStr := buf.String()
	fmt.Println("Content of the given CID is \n",newStr)
}

