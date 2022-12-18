package encrypt

import (
	`bytes`
	`crypto/rand`
	`crypto/rsa`
	`crypto/x509`
	`encoding/base64`
	`encoding/pem`
	`github.com/gookit/goutil/fsutil`
)

// RSAEncrypt RSA加密
type RSAEncrypt struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func publicKey(public string) (*rsa.PublicKey, error) {
	buf := fsutil.GetContents(public)
	block, _ := pem.Decode(buf)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}
func privateKey(private string) (*rsa.PrivateKey, error) {
	buf := fsutil.GetContents(private)
	block, _ := pem.Decode(buf)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
func RSA(public, private string) (*RSAEncrypt, error) {
	publicKey, err := publicKey(public)
	if err != nil {
		return nil, err
	}
	privateKey, err := privateKey(private)
	if err != nil {
		return nil, err
	}
	return &RSAEncrypt{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// PublicEncryption 公钥加密
func (r *RSAEncrypt) PublicEncryption(data string) (encrypted string, err error) {
	var encrypt []byte
	encrypt, err = rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, []byte(data))
	if err != nil {
		return
	}
	encrypted = base64.StdEncoding.EncodeToString(encrypt)
	return
}

// PrivateDecryption 私钥解密
func (r *RSAEncrypt) PrivateDecryption(data string) (decrypted string, err error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}
	var decrypt []byte
	decrypt, err = rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, encryptedBytes)
	decrypted = string(decrypt)
	return
}

func GenerateKey(bits int) (public *bytes.Buffer, private *bytes.Buffer, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	private = new(bytes.Buffer)
	err = pem.Encode(private, block)
	if err != nil {
		return nil, nil, err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	public = new(bytes.Buffer)
	err = pem.Encode(public, block)
	if err != nil {
		return nil, nil, err
	}
	return public, private, nil
}
