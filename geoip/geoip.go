
package geoip

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type GeoIPReq struct{}

type GeoIPResp struct {
	Country, Region, City, CityLatLong, IP string
}

type GeoIPService struct{}

func (*GeoIPService) Get(
	r *http.Request, req *GeoIPReq, resp *GeoIPResp) error {

	resp.Country = r.Header.Get("X-AppEngine-Country")
	resp.Region = r.Header.Get("X-AppEngine-Region")
	resp.City = r.Header.Get("X-AppEngine-City")
	resp.CityLatLong = r.Header.Get("X-AppEngine-CityLatLong")
	resp.IP = r.RemoteAddr

	return nil
}

func init() {

	geoip := &GeoIPService{}
	api, err := endpoints.RegisterService(geoip,
		"geoip", "v1", "GeoIP Service", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("Get").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc =
		"get", "GET", "geoip", "Get GeoIP information."

}
