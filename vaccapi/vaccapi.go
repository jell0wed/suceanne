package vaccapi

import (
	"github.com/spf13/viper"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	log "github.com/sirupsen/logrus"
)

const clientKey string = "eJUWrzRv34qFSaYk"
const secret string = "Cyu5jcR4zyK6QEPn1hdIGXB5QIDAQABMA0GC"
const mainUrlFormat  string = "https://eco-{country}-api.ecovacs.com/v1/private/{country}/{lang}/{deviceId}/{appCode}/{appVersion}/{channel}/{deviceType}"
var publicKey *rsa.PublicKey

type Vaccapi struct {
	publicKey []byte
	deviceId string
	countryCode string
	continentCode string
}

func (api Vaccapi) login() {

}

func initializePublicKey() *rsa.PublicKey {
	pemPublicKey := []byte(viper.Get("publicKey").(string))
	block, _ := pem.Decode(pemPublicKey) // ignore the remaining data
	if block == nil || block.Type != "PUBLIC KEY" {
		log.WithFields(log.Fields{"block": block}).Error("Unable to decode PEM public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Error("Unable to parse PKIX public key")
	}
	return pub.(*rsa.PublicKey)
}

func Encrypt(payload string) []byte {
	payload_b := []byte(payload)
	
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, payload_b)

	if err != nil {
		log.WithFields(log.Fields{"payload_str": payload}).Error("Unable to encrypt payload string")
	}

	return ciphertext
}

// Factory method to create a new api instance
func NewVaccapi() Vaccapi {
	publicKey = initializePublicKey()

	var api Vaccapi
	api.deviceId = viper.Get("deviceId").(string)
	return api
}