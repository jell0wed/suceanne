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
const mainUrlFormat string = "https://eco-{country}-api.ecovacs.com/v1/private/{country}/{lang}/{deviceId}/{appCode}/{appVersion}/{channel}/{deviceType}"
const pemPublicKey_s string = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDb8V0OYUGP3Fs63E1gJzJh+7iq
eymjFUKJUqSD60nhWReZ+Fg3tZvKKqgNcgl7EGXp1yNifJKUNC/SedFG1IJRh5hB
eDMGq0m0RQYDpf9l0umqYURpJ5fmfvH/gjfHe3Eg/NTLm7QEa0a0Il2t3Cyu5jcR
4zyK6QEPn1hdIGXB5QIDAQAB
-----END PUBLIC KEY-----`
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
	pemPublicKey := []byte(pemPublicKey_s)
	block, _ := pem.Decode(pemPublicKey) // ignore the remaining data
	if block == nil || block.Type != "PUBLIC KEY" {
		log.WithFields(log.Fields{"block": block}).Fatal("Unable to decode PEM public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("Unable to parse PKIX public key")
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

func init() {
	publicKey = initializePublicKey()
}

// Factory method to create a new api instance
func NewVaccapi() Vaccapi {
	var api Vaccapi
	api.deviceId = viper.Get("deviceId").(string)
	return api
}