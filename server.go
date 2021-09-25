package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"bytes"
	"strings"
	shell "github.com/ipfs/go-ipfs-api"
)
var sh *shell.Shell
type (
	ipfsItem struct {
		Key  string  `json:"key" form:"key" query:"key"`
		Data string `json:"data" form:"data" query:"data"`
	}
)
//----------
// Handlers
//----------

func addData(c echo.Context) error {
	item := new(ipfsItem)
	if err := c.Bind(item); err != nil {
		return err
	}
	//AES only supports key sizes of 16, 24 or 32 bytes. For an example, suggested key:abcdabcdabcdabcd 
	log.Printf(item.Key)
	log.Printf(item.Data)	
	//encrypt value to base64
	log.Printf("password:\t %v \n", []byte(item.Key))
	crytoText := encrypt([]byte(item.Key), item.Data)
	log.Printf(crytoText)
	cid, _ := sh.Add(strings.NewReader(crytoText))
	log.Printf(cid)
	return c.JSON(http.StatusCreated, cid)
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		// panic(err)
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func getData(c echo.Context) error {
	cid := c.Param("cid")
	log.Printf(cid)
	resp, err := sh.Cat(cid)
	if err != nil {
		log.Printf("err: ", err)
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp)
	//The return of Cat is io.ReaderCloser, convert to string for output
	encodedData := buf.String()
	return c.JSON(http.StatusOK, encodedData)
}


func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	sh = shell.NewShell("localhost:5001")
	// Routes
	e.POST("/add", addData)
	e.GET("/get/:cid", getData)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
