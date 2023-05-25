package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/echoframe/conf"

	"github.com/echoframe/pkg/uber"

	"golang.org/x/crypto/md4"
)

type AesEncrypt struct {
	key   []byte
	iv    []byte
	block cipher.Block
}

var AesEcp AesEncrypt

func init() {
	AesEcp.key = []byte("bGgGfWb3Kg2s4gcG")
	AesEcp.iv = []byte("aebksHkG4jAEk2Ag")
	var err error
	AesEcp.block, err = aes.NewCipher(AesEcp.key)
	if err != nil {
		panic(err)
	}
}

// AesBase64Encrypt
func (a *AesEncrypt) AesBase64Encrypt(in string) (string, error) {
	origData := []byte(in)
	origData = PKCS5Adding(origData, a.block.BlockSize())
	crypted := make([]byte, len(origData))
	bm := cipher.NewCBCEncrypter(a.block, a.iv)
	bm.CryptBlocks(crypted, origData)
	var b = base64.StdEncoding.EncodeToString(crypted)
	return b, nil
}

// AesBase64Decrypt
func (a *AesEncrypt) AesBase64Decrypt(b string) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return "", err
	}
	origData := make([]byte, len(crypted))
	bm := cipher.NewCBCDecrypter(a.block, a.iv)
	bm.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	var out = string(origData)
	return out, nil
}

// PKCS5Adding
func PKCS5Adding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS5UnPadding
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// Md4
func Md4(str string) string {
	srcByte := []byte(str)
	hash := md4.New()
	hash.Write(srcByte)
	c := hash.Sum(nil)
	return hex.EncodeToString(c)
}

// Md5
func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}

// Sha256
func Sha256(str string) string {
	srcByte := []byte(str)
	hash := sha256.New()
	hash.Write(srcByte)
	hashBytes := hash.Sum(nil)
	sha256String := hex.EncodeToString(hashBytes)
	return sha256String
}

func Sha512(str string) string {
	srcByte := []byte(str)
	hash := sha512.New()
	hash.Write(srcByte)
	hashBytes := hash.Sum(nil)
	sha256String := hex.EncodeToString(hashBytes)
	return sha256String
}

// FileMd5
func FileMd5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		uber.EchoScaLog.Error("os open error")
		return "", err
	}
	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		uber.EchoScaLog.Error("io copy error")
		return "", err
	}
	md5Str := hex.EncodeToString(hash.Sum(nil))
	return md5Str, nil
}

// GeneratePasswd
func GeneratePasswd(str string) string {
	return Sha256(fmt.Sprintf("%s%s", Md5(str), conf.Config.Slat))
}
