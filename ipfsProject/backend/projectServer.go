package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"bytes"
	"strings"
	"encoding/json"
	shell "github.com/ipfs/go-ipfs-api"
)
var sh *shell.Shell
type (
	ipfsItem struct {
		Key  string  `json:"key" form:"key" query:"key"`
		Data string `json:"data" form:"data" query:"data"`
	}
)
type Response struct {
	Status  int `json:"status" xml:"status"`
	CID string `json:"cid" xml:"cid"`
}
type KeyStruct struct {
	Key string `json:"key" form:"key" query:"key"`
}
//----------
// Handlers
//----------

func addData(c echo.Context) error {
	item := new(ipfsItem)
	if err := c.Bind(item); err != nil {
		return err
	}
	log.Printf(item.Key)
	log.Printf(item.Data)	
	//encrypt value to base64
	log.Printf("password:\t %v \n", []byte(item.Key))
	crytoText := encrypt([]byte(item.Key), item.Data)
	log.Printf(crytoText)
	cid, _ := sh.Add(strings.NewReader(crytoText))
	log.Printf(cid)

	resp := &Response{
		Status:  http.StatusOK,
		CID: cid,
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(resp)
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key needs tp be 16, 24 or 32 bytes.
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func getData(c echo.Context) error {
	keyParam := c.QueryParam("key")
	log.Printf(keyParam)
	key := []byte(keyParam)
	log.Printf("key:\t %v \n", key)
	cid := c.Param("cid")
	log.Printf(cid)
	resp, err := sh.Cat(cid)
	if err != nil {
		log.Printf("err: ", err)
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp)
	encodedData := buf.String()
	text := decrypt(key, encodedData)
	log.Printf(text)
	return c.JSON(http.StatusOK, text)
}

func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	sh = shell.NewShell("localhost:5001")
	// Routes
	e.POST("/add", addData)
	e.GET("/get/:cid", getData)

	// Start server
	fmt.Printf("Starting server at port 1323\n")
	e.Logger.Fatal(e.Start(":1323"))
}
