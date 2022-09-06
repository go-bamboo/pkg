package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func ErrBadPublicKey(format string, a ...interface{}) error {
	return errors.InternalServer("ErrBadPublicKey", fmt.Sprintf(format, a...))
}

func IsErrBadPublicKey(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrBadPublicKey" && se.Code == 500
}

func RSAGen() (puk string, priv string, err error) {
	pkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	privb := x509.MarshalPKCS1PrivateKey(pkey)
	pukb := x509.MarshalPKCS1PublicKey(&pkey.PublicKey)
	puk = base64.StdEncoding.EncodeToString(pukb)
	priv = base64.StdEncoding.EncodeToString(privb)
	return
}

func RSAEncrypt(orgidata []byte, publickey []byte) (pb []byte, err error) {
	block, _ := pem.Decode(publickey)
	if block == nil {
		err = ErrBadPublicKey("publickey(%v)", publickey)
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := pubInterface.(*rsa.PublicKey)
	pb, err = rsa.EncryptPKCS1v15(rand.Reader, pub, orgidata) //加密
	return
}

func RSAEncryptWithPub(orgidata []byte, pub *rsa.PublicKey) (pb []byte, err error) {
	pb, err = rsa.EncryptPKCS1v15(rand.Reader, pub, orgidata) //加密
	return
}

func RSADecrypt(cipertext, privatekey []byte) (pb []byte, err error) {
	block, _ := pem.Decode(privatekey)
	if block == nil {
		err = ErrBadPublicKey("publickey(%v)", privatekey)
		return
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	pb, err = rsa.DecryptPKCS1v15(rand.Reader, priv, cipertext)
	return
}

func RSADecryptWitPriv(cipertext []byte, priv *rsa.PrivateKey) (pb []byte, err error) {
	pb, err = rsa.DecryptPKCS1v15(rand.Reader, priv, cipertext)
	return
}
