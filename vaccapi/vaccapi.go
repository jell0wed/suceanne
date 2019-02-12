package vaccapi

import (
	"github.com/spf13/viper"
	"strings"
	"crypto/rsa"
	"crypto/rand"
	b64 "encoding/base64"
)

const clientKey string = "eJUWrzRv34qFSaYk"
const secret string = "Cyu5jcR4zyK6QEPn1hdIGXB5QIDAQABMA0GC"
const publicKey []byte = b64.StdEncoding.DecodeString("MIIB/TCCAWYCCQDJ7TMYJFzqYDANBgkqhkiG9w0BAQUFADBCMQswCQYDVQQGEwJjbjEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZhdWx0IENvbXBhbnkgTHRkMCAXDTE3MDUwOTA1MTkxMFoYDzIxMTcwNDE1MDUxOTEwWjBCMQswCQYDVQQGEwJjbjEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZhdWx0IENvbXBhbnkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDb8V0OYUGP3Fs63E1gJzJh+7iqeymjFUKJUqSD60nhWReZ+Fg3tZvKKqgNcgl7EGXp1yNifJKUNC/SedFG1IJRh5hBeDMGq0m0RQYDpf9l0umqYURpJ5fmfvH/gjfHe3Eg/NTLm7QEa0a0Il2t3Cyu5jcR4zyK6QEPn1hdIGXB5QIDAQABMA0GCSqGSIb3DQEBBQUAA4GBANhIMT0+IyJa9SU8AEyaWZZmT2KEYrjakuadOvlkn3vFdhpvNpnnXiL+cyWy2oU1Q9MAdCTiOPfXmAQt8zIvP2JC8j6yRTcxJCvBwORDyv/uBtXFxBPEC6MDfzU2gKAaHeeJUWrzRv34qFSaYkYta8canK+PSInylQTjJK9VqmjQ")
const mainUrlFormat  string = "https://eco-{country}-api.ecovacs.com/v1/private/{country}/{lang}/{deviceId}/{appCode}/{appVersion}/{channel}/{deviceType}"

type Vaccapi struct {
	deviceId string
	countryCode string
	continentCode string
}

func encrypt(payload string) {
	payload_b = []byte(payload)
	ciphertext, err := rsa.SignPKCS1v15(rand.Reader, &publicKey, payload_b)

	if err != nil {

	}
}

func (api Vaccapi) login() {

}

// Factory method to create a new api instance
func NewVaccapi() {
	var api Vaccapi
	api.baseUrl = "http://" + viper.Get("baseUrl")
	api.deviceId = viper.Get("deviceId")

	return api
}