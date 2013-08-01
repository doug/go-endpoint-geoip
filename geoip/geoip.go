// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
