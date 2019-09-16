package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"os"
)

func sesSmtpPasswordFromSecretKey(secret []byte) string {
	version := byte(0x02)
	message := []byte("SendRawEmail")
	hmacKey := []byte(secret)
	h := hmac.New(sha256.New, hmacKey)
	h.Write(message)
	rawSig := h.Sum(nil)
	versionedSig := make([]byte, 0, len(rawSig)+1)
	versionedSig = append(versionedSig, version)
	versionedSig = append(versionedSig, rawSig...)
	return base64.StdEncoding.EncodeToString(versionedSig)
}

func main() {
	secret, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	smtpPassword := sesSmtpPasswordFromSecretKey(secret)
	_, err = os.Stdout.WriteString(smtpPassword)
	if err != nil {
		panic(err)
	}
}
