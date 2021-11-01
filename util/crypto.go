/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 15:29
 */

package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

type AesCBC struct {
	key      string
	isBase64 bool
}

func NewAesCBC(key string, base bool) *AesCBC {
	return &AesCBC{
		key:      key,
		isBase64: base,
	}
}

func (ac *AesCBC) Encrypt(plaintext string) (chipText string, err error) {
	var (
		block cipher.Block
	)

	if block, err = aes.NewCipher([]byte(ac.key)); err != nil {
		return
	}
	blockSize := block.BlockSize()
	padding := PKCS7Padding([]byte(plaintext), blockSize)
	cipherText := make([]byte, blockSize+len(padding))
	iv := cipherText[:block.BlockSize()]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	cip := cipher.NewCBCEncrypter(block, iv)
	cip.CryptBlocks(cipherText[blockSize:], padding)

	if ac.isBase64 {
		chipText = base64.StdEncoding.EncodeToString(cipherText)
	} else {
		chipText = string(cipherText)
	}
	return

}

func (ac *AesCBC) Decrypt(chipText string) (plaintext string, err error) {
	var (
		block       cipher.Block
		chipperText []byte
	)

	if block, err = aes.NewCipher([]byte(ac.key)); err != nil {
		return
	}

	blockSize := block.BlockSize()

	if ac.isBase64 {
		if chipperText, err = base64.StdEncoding.DecodeString(chipText); err != nil {
			return "", fmt.Errorf("wrong param :%s", err)
		}
	} else {
		chipperText = []byte(chipText)
	}

	if len(chipperText) < blockSize {
		return "", errors.New("encrypt string is too short")
	}

	iv := chipperText[:blockSize]
	chipperText = chipperText[blockSize:]
	if len(chipperText)%blockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	cip := cipher.NewCBCDecrypter(block, iv)
	cip.CryptBlocks(chipperText, chipperText)
	chipperText = PKCS7UnPadding(chipperText)

	return string(chipperText), nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	return plantText[:(length - unPadding)]
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}
