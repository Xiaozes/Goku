package Gapi

import (
	"Goku/Ghttp"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Quake struct {
	PassWord string
	result   Result
	http     Ghttp.Http
}

type InfoJson struct {
	Code int `json:"code"`
	Data struct {
		AvatarID       string `json:"avatar_id"`
		BanStatus      string `json:"ban_status"`
		Baned          bool   `json:"baned"`
		ConstantCredit int    `json:"constant_credit"`
		Credit         int    `json:"credit"`
		Disable        struct {
			DisableTime interface{} `json:"disable_time"`
			StartTime   interface{} `json:"start_time"`
		} `json:"disable"`
		EnterpriseInformation struct {
			Email  interface{} `json:"email"`
			Name   interface{} `json:"name"`
			Status string      `json:"status"`
		} `json:"enterprise_information"`
		FreeQueryAPICount  int    `json:"free_query_api_count"`
		ID                 string `json:"id"`
		InvitationCodeInfo struct {
			Code                string `json:"code"`
			InviteAcquireCredit int    `json:"invite_acquire_credit"`
			InviteNumber        int    `json:"invite_number"`
		} `json:"invitation_code_info"`
		IsCashedInvitationCode    bool   `json:"is_cashed_invitation_code"`
		MobilePhone               string `json:"mobile_phone"`
		MonthRemainingCredit      int    `json:"month_remaining_credit"`
		PersistentCredit          int    `json:"persistent_credit"`
		PersonalInformationStatus bool   `json:"personal_information_status"`
		PrivacyLog                struct {
			AnonymousModel bool   `json:"anonymous_model"`
			QuakeLogStatus bool   `json:"quake_log_status"`
			QuakeLogTime   string `json:"quake_log_time"`
			Status         bool   `json:"status"`
			Time           string `json:"time"`
		} `json:"privacy_log"`
		Role []struct {
			Credit   int    `json:"credit"`
			Fullname string `json:"fullname"`
			Priority int    `json:"priority"`
		} `json:"role"`
		RoleValidity struct {
		} `json:"role_validity"`
		Source string `json:"source"`
		Time   string `json:"time"`
		Token  string `json:"token"`
		User   struct {
			Email    string   `json:"email"`
			Fullname string   `json:"fullname"`
			Group    []string `json:"group"`
			ID       string   `json:"id"`
			Username string   `json:"username"`
		} `json:"user"`
	} `json:"data"`
	Message string   `json:"message"`
	Meta    struct{} `json:"meta"`
}

type ServiceInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Time      time.Time `json:"time"`
		Transport string    `json:"transport"`
		Service   struct {
			HTTP struct {
				HTMLHash string `json:"html_hash"`
				Favicon  struct {
					Hash     string `json:"hash"`
					Location string `json:"location"`
					Data     string `json:"data"`
				} `json:"favicon"`
				Robots          string `json:"robots"`
				SitemapHash     string `json:"sitemap_hash"`
				Server          string `json:"server"`
				Body            string `json:"body"`
				XPoweredBy      string `json:"x_powered_by"`
				MetaKeywords    string `json:"meta_keywords"`
				RobotsHash      string `json:"robots_hash"`
				Sitemap         string `json:"sitemap"`
				Path            string `json:"path"`
				Title           string `json:"title"`
				Host            string `json:"host"`
				SecurityText    string `json:"security_text"`
				StatusCode      int    `json:"status_code"`
				ResponseHeaders string `json:"response_headers"`
			} `json:"http"`
			Version  string `json:"version"`
			Name     string `json:"name"`
			Product  string `json:"product"`
			Banner   string `json:"banner"`
			Response string `json:"response"`
		} `json:"service"`
		Images     []interface{} `json:"images"`
		OsName     string        `json:"os_name"`
		Components []interface{} `json:"components"`
		Location   struct {
			DistrictCn  string    `json:"district_cn"`
			ProvinceCn  string    `json:"province_cn"`
			Gps         []float64 `json:"gps"`
			ProvinceEn  string    `json:"province_en"`
			CityEn      string    `json:"city_en"`
			CountryCode string    `json:"country_code"`
			CountryEn   string    `json:"country_en"`
			Radius      float64   `json:"radius"`
			DistrictEn  string    `json:"district_en"`
			Isp         string    `json:"isp"`
			StreetEn    string    `json:"street_en"`
			Owner       string    `json:"owner"`
			CityCn      string    `json:"city_cn"`
			CountryCn   string    `json:"country_cn"`
			StreetCn    string    `json:"street_cn"`
		} `json:"location"`
		Asn       int    `json:"asn"`
		Hostname  string `json:"hostname"`
		Org       string `json:"org"`
		OsVersion string `json:"os_version"`
		IsIpv6    bool   `json:"is_ipv6"`
		IP        string `json:"ip"`
		Port      int    `json:"port"`
	} `json:"data"`
	Meta struct {
		Total        int    `json:"total"`
		PaginationID string `json:"pagination_id"`
	} `json:"meta"`
}

func (s *Quake) GetInfo() string {
	return "QuakeImpl ver 0.1 with  "
}

func (s *Quake) SetPassword(password string) {
	s.PassWord = password
}

func (s *Quake) GetResult(quser, key, size string) Result {
	s.result = Result{}
	resp := s.send(quser, key, size)
	var Quakers ServiceInfo
	json.Unmarshal([]byte(string(resp)), &Quakers)
	fmt.Println(Quakers)
	return s.result
}

func (s *Quake) Login(key string) bool {
	urls := fmt.Sprintf("https://quake.360.net/api/v3/user/info")
	s.http.New("GET", urls)
	s.http.SetHeader("X-QuakeToken", key)
	s.http.Execute()
	defer s.http.Close()
	resp, err := s.http.Byte()
	result := InfoJson{}
	json.Unmarshal(resp, &result)
	if err != nil {
		log.Println(err)
		return false
	} else {
		if result.Code != 0 {
			log.Println("请检查你session的有效性")
			return false
		} else {
			return true
		}
	}
}

func (s *Quake) send(query, key, size string) []byte {
	var result []byte
	//query = base64.URLEncoding.EncodeToString([]by
	//	jsonData["ignore_cache"] = "true"te(query))
	urls := fmt.Sprintf("https://quake.360.cn/api/v3/scroll/quake_service")
	//s.http.SetHeader("X-QuakeToken", key)
	jsonData := make(map[string]interface{})
	jsonData["query"] = query
	jsonData["start"] = 0
	jsonData["size"] = size
	s.http.Post(urls, jsonData)
	s.http.SetHeader("X-QuakeToken", key)
	s.http.SetProxy("http://127.0.0.1:8080")
	s.http.Execute()
	defer s.http.Close()
	resp, err := s.http.Byte()
	if err != nil {
		log.Println(err)
		return result
	}
	return resp
}
