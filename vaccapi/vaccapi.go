package vaccapi

import (
	"github.com/spf13/viper"
	//"strings"
	"crypto/rsa"
	"crypto/rand"
	log "github.com/sirupsen/logrus"
)

const clientKey string = "eJUWrzRv34qFSaYk"
const secret string = "Cyu5jcR4zyK6QEPn1hdIGXB5QIDAQABMA0GC"
const mainUrlFormat  string = "https://eco-{country}-api.ecovacs.com/v1/private/{country}/{lang}/{deviceId}/{appCode}/{appVersion}/{channel}/{deviceType}"
var pemPublicKey []byte 

type Vaccapi struct {
	publicKey []byte
	deviceId string
	countryCode string
	continentCode string
}

func (api Vaccapi) login() {

}

func Encrypt(payload string) {
	payload_b := []byte(payload)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &pemPublicKey, payload_b)

	if err != nil {
		log.WithFields(log.Fields{"payload_str": payload}).Error("Unable to encrypt payload string")
	}

	return ciphertext
}

// Factory method to create a new api instance
func NewVaccapi() Vaccapi {
	pemPublicKey := []byte(viper.Get("publicKey").(string))

	var api Vaccapi
	api.deviceId = viper.Get("deviceId").(string)
	return api
}