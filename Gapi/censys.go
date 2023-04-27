package Gapi

import (
	"Goku/Ghttp"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

type Censys struct {
	UserName string
	PassWord string
	result   Result
	http     Ghttp.Http
}

type CensysResult struct {
	Results []struct {
		Data  []string `json:"parsed.names"`
		Data1 []string `json:"parsed.extensions.subject_alt_name.dns_names"`
	} `json:"results"`
}

func (s *Censys) GetInfo() string {
	return "CensysImpl ver 0.1 with  "
}

func (s *Censys) SetUsername(username string) {
	s.UserName = username
}
func (s *Censys) SetPassword(password string) {
	s.PassWord = password
}

func (s *Censys) GetResult(quser string) Result {
	var tmp []string
	s.result = Result{}
	resp := s.send(quser)
	result := CensysResult{}
	json.Unmarshal(resp, &result)
	fmt.Println(len(result.Results))
	for _, i2 := range result.Results {
		for _, i3 := range i2.Data {
			tmp = append(tmp, i3)
		}
		for _, i4 := range i2.Data1 {
			tmp = append(tmp, i4)
		}

	}
	tmp1 := RemoveRepeatedElement(tmp)
	for _, i4 := range tmp1 {
		s.result[i4] = quser
	}
	return s.result

}

func (s *Censys) send(domain string) []byte {
	var result []byte
	apiURL := "https://search.censys.io/api/v1/search/certificates"
	jsonData := make(map[string]interface{})
	jsonData["query"] = domain
	jsonData["fields"] = []string{"parsed.names", "parsed.extensions.subject_alt_name.dns_names"}
	jsonData["flatten"] = true
	s.http.Post(apiURL, jsonData)
	key := s.UserName + ":" + s.PassWord
	key = "Basic " + base64.StdEncoding.EncodeToString([]byte(key))
	s.http.SetHeader("Authorization", key)
	s.http.Execute()
	defer s.http.Close()
	resp, err := s.http.Byte()
	//fmt.Println((resp))
	if err != nil {
		log.Println(err)
		return result
	}
	return resp
	//jsonData, err := json.Marshal(data)
	//if err != nil {
	//	panic(err)
	//}
	//req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	//req.SetBasicAuth(s.UserName, s.PassWord)
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Accept", "application/json")

	//设置http客户端参数
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //忽略https验证
	//}

	//// Read the response body
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	err = errors.New(fmt.Sprintf("[Censys-err] Can't reading response body %s", domain))
	//	return err
	//}
	//return resp
}

// 去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}
