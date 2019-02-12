package vaccapi

import (
	"github.com/spf13/viper"
)

type Vaccapi struct {
	baseUrl string
	deviceId string
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