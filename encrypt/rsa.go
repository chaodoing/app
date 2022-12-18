package encrypt

import (
	`bytes`
	`crypto`
	`crypto/rand`
	`crypto/rsa`
	`crypto/sha256`
	`crypto/x509`
	`encoding/base64`
	`encoding/pem`
	`github.com/gookit/goutil/fsutil`
)

// RSAEncrypt RSA加密
type RSAEncrypt struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// publicKey 公钥文件
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

// privateKey 私钥文件
func privateKey(private string) (*rsa.PrivateKey, error) {
	buf := fsutil.GetContents(private)
	block, _ := pem.Decode(buf)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// publicKeyString 公钥字符串
func publicKeyString(publicKey string) (*rsa.PublicKey, error) {
	publicKeyInterface, err := x509.ParsePKIXPublicKey([]byte(publicKey))
	if err != nil {
		return nil, err
	}
	return publicKeyInterface.(*rsa.PublicKey), nil
}

// privateKeyString 私钥字符串
func privateKeyString(privateKey string) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey([]byte(privateKey))
}

// RSAString RSA字符串
func RSAString(publicKey, privateKey string) (*RSAEncrypt, error) {
	public, err := publicKeyString(publicKey)
	if err != nil {
		return nil, err
	}
	private, err := privateKeyString(privateKey)
	if err != nil {
		return nil, err
	}
	return &RSAEncrypt{
		PrivateKey: private,
		PublicKey:  public,
	}, nil
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
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

// SignWithSha256 签名
func (r *RSAEncrypt) SignWithSha256(ciphertext string) (encrypted string, err error) {
	var (
		data = []byte(ciphertext)
		h    = sha256.New()
	)
	h.Write(data)
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), err
}

// VerySignWithSha256 验证
func (r *RSAEncrypt) VerySignWithSha256(ciphertext, signText string) (success bool, e error) {
	var (
		data      = []byte(ciphertext)
		sign, err = base64.StdEncoding.DecodeString(signText)
	)
	if err != nil {
		return false, err
	}
	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, hashed[:], sign)
	if err != nil {
		return false, err
	}
	return true, nil
}

// PublicEncryption 公钥加密
func (r *RSAEncrypt) PublicEncryption(data string) (encrypted string, err error) {
	var encrypt []byte
	encrypt, err = rsa.EncryptPKCS1v15(rand.Reader, r.PublicKey, []byte(data))
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
	decrypt, err = rsa.DecryptPKCS1v15(rand.Reader, r.PrivateKey, encryptedBytes)
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
