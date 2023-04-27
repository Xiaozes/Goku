package Gapi

import (
	"Goku/Gconvert"
	"Goku/Ghttp"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

type Hunter struct {
	PassWord string
	Page     string
	Number   string
	result   Result
	http     Ghttp.Http
}

type HunterJsonResult struct {
	Code int64 `json:"code"`
	Data struct {
		AccountType string `json:"account_type"`
		Arr         []struct {
			AsOrg        string `json:"as_org"`
			Banner       string `json:"banner"`
			BaseProtocol string `json:"base_protocol"`
			City         string `json:"city"`
			Company      string `json:"company"`
			Component    []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"component"`
			Country        string `json:"country"`
			Domain         string `json:"domain"`
			IP             string `json:"ip"`
			IsRisk         string `json:"is_risk"`
			IsRiskProtocol string `json:"is_risk_protocol"`
			IsWeb          string `json:"is_web"`
			Isp            string `json:"isp"`
			Number         string `json:"number"`
			Os             string `json:"os"`
			Port           int64  `json:"port"`
			Protocol       string `json:"protocol"`
			Province       string `json:"province"`
			StatusCode     int64  `json:"status_code"`
			UpdatedAt      string `json:"updated_at"`
			URL            string `json:"url"`
			WebTitle       string `json:"web_title"`
		} `json:"arr"`
		ConsumeQuota string `json:"consume_quota"`
		RestQuota    string `json:"rest_quota"`
		SyntaxPrompt string `json:"syntax_prompt"`
		Time         int64  `json:"time"`
		Total        int64  `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}

func (s *Hunter) GetInfo() string {
	return "HunterPro ver 0.1 with  "
}

func (s *Hunter) SetPassword(password string) {
	s.PassWord = password
}
func (s *Hunter) SetPage(page string) {
	s.Page = page
}
func (s *Hunter) SetNubmer(number string) {
	s.Number = number
}

func (s *Hunter) GetResult(quser string) Result {
	s.result = Result{}
	resp := s.send(quser)
	var hunterJR HunterJsonResult
	json.Unmarshal([]byte(string(resp)), &hunterJR)
	//fmt.Println((hunterJR.Data.Arr))
	for _, v := range hunterJR.Data.Arr {
		if v.Domain == "" {
			s.result[v.IP+":"+Gconvert.Int2String(v.Port)] = v.IP
		} else {
			s.result[v.Domain] = v.IP
		}
	}

	return s.result

}

func (s *Hunter) send(query string) []byte {
	var result []byte
	query = base64.URLEncoding.EncodeToString([]byte(query))
	urls := fmt.Sprintf("https://hunter.qianxin.com/openApi/search?api-key=%s&search=%s&page=%s&page_size=%s&is_web=3&port_filter=true", s.PassWord, query, s.Page, s.Number)
	s.http.New("GET", urls)
	s.http.Execute()
	defer s.http.Close()
	resp, err := s.http.Byte()
	//res1, _ := s.http.Text()
	if err != nil {
		log.Println(err)
		return result
	}
	return resp
}
