package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

const publicKeyPem = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDvekdPMHN3AYhm/vktJT+YJr7
cI5DcsNKqdsx5DZX0gDuWFuIjzdwButrIYPNmRJ1G8ybDIF7oDW2eEpm5sMbL9zs
9ExXCdvqrn51qELbqj0XxtMTIpaCHFSI50PfPpTFV9Xt/hmyVwokoOXFlAEgCn+Q
CgGs52bFoYMtyi+xEQIDAQAB
-----END PUBLIC KEY-----
`

var publicKey *rsa.PublicKey

func init() {
	block, _ := pem.Decode([]byte(publicKeyPem))
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	osutil.PanicErr(err)
	publicKey = key.(*rsa.PublicKey)
}

func RSAEncrypt(content string) string {
	bs, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(content))
	osutil.PanicErr(err)
	return base64.StdEncoding.EncodeToString(bs)
}
