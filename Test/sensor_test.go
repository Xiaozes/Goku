package test

import (
	"Goku/Gapi"
	"fmt"
	"log"
	"testing"
)

func TestFofa(t *testing.T) {
	ff := Gapi.Fofa{}
	ff.SetAccount("574286399@qq.com")
	ff.SetPassword("c57b4617dcf834d5cd051359b13a644f")
	r := ff.GetResult("ip=\"222.161.231.2\"")
	for k, v := range r {
		fmt.Println(k, v)
	}
}

func TestHunter(t *testing.T) {
	hh := Gapi.Hunter{}
	hh.SetPassword("c35f5636120c216cedb124a8c1a5176887056af98aeba4d8cc09489f19181f50")
	hh.SetPage("1")
	hh.SetNubmer("100")
	r := hh.GetResult("ip=\"114.113.80.203\"")
	for k, v := range r {
		fmt.Println(k, v)
	}
}

func TestQuake(t *testing.T) {
	qq := Gapi.Quake{}
	qq.SetPassword("xxxxx")
	qq.Login(qq.PassWord)
	qq.GetResult("ip=\"114.113.80.203\"", qq.PassWord, "5")

}

func TestZoom(t *testing.T) {
	z := Gapi.ZoomEye{}
	z.SetType("subdomain")
	z.SetDomain("weibo.com")
	z.SetPassword("xxx-xxx-xxx-xxx-xxx")
	r := z.GetResult()
	log.Println(r)
}

func TestShodan(t *testing.T) {
	sd := Gapi.Shodan{}
	// shodan用的chrome插件的api，只能用来获取端口
	sd.SetDomain("172.67.168.89")
	sd.SetPassword("EAPkjRn3SB4ishKUUSi4R0grSSI1GXbB")
	sd.SetType("ports")
	rr := sd.GetResult()
	for k, v := range rr {
		fmt.Println(k, v)
	}
	//log.Println(rr)
}
func TestCensys(t *testing.T) {
	ce := Gapi.Censys{}
	ce.SetUsername("xxxxxx")
	ce.SetPassword("xxxxxxxxxxx")
	cresult := ce.GetResult("xiaomi.com")
	for k, v := range cresult {
		fmt.Println(k, v)
	}

}

//func TestBeian(t *testing.T) {
//	beian := Gsensor.Beian{}
//	beian.SetDomain("baidu.com")
//	beian.SetAccount("http://127.0.0.1:65511/")
//	beian.SetType("beian")
//	result := beian.GetResult()
//	log.Printf("%#v", result)
//}
//
//func TestSecTrail(t *testing.T) {
//	log.SetFlags(log.Lshortfile | log.LstdFlags)
//
//	ss := Gsensor.SecurityTrails{}
//	//可以不登录，只能查询第一页
//	ss.SetAccount("xxxxx@gmail.com")
//	ss.SetPassword("xxxxxxx")
//	ss.Login(true)
//	ss.MaxPage = 10 // 自定义最大翻页，登录后默认100页
//	ss.SetDomain("360.cn")
//	ss.SetType("subdomain")
//	//ss.SetType("ahistory")
//
//	//ss.SetDomain("172.67.168.89")
//	//ss.SetType("sameserver")
//	r := ss.GetResult()
//	log.Println(len(r))
//	log.Println(r)
//
//}
//func TestSecTrailApi(t *testing.T) {
//	log.SetFlags(log.Lshortfile | log.LstdFlags)
//
//	ss := Gsensor.SecurityTrailsApi{}
//	//可以不登录，只能查询第一页
//	ss.SetPassword("NQIcGiQBA53myDkCS8wXj2d4MzauIdkH")
//	ss.SetDomain("360.cn")
//	ss.SetType("subdomain")
//	r := ss.GetResult()
//	log.Println(len(r))
//	log.Println(r)
//
//}
//func TestZoom(t *testing.T) {
//	z := Gsensor.ZoomEye{}
//	z.SetType("subdomain")
//	z.SetDomain("weibo.com")
//	z.SetPassword("xxx-xxx-xxx-xxx-xxx")
//	r := z.GetResult()
//	log.Println(r)
//}
//func TestShodan(t *testing.T) {
//	sd := Gsensor.Shodan{}
//	// shodan用的chrome插件的api，只能用来获取端口
//	sd.SetDomain("172.67.168.89")
//	sd.SetType("ports")
//	rr := sd.GetResult()
//	log.Println(rr)
//}
//func TestBuff(t *testing.T) {
//	bf := Gsensor.Bufferover{}
//	bf.SetType("subdomain")
//	bf.SetDomain("baidu.com")
//	r := bf.GetResult()
//	log.Println(r)
//}
//
//func TestRapid(t *testing.T) {
//	bf := Gsensor.RapidDns{}
//	bf.SetType("subdomain")
//	bf.SetDomain("freebuf.com")
//	r := bf.GetResult()
//	log.Println(r)
//}
//func TestCrt(t *testing.T) {
//	bf := Gsensor.CrtSh{}
//	bf.SetType("subdomain")
//	bf.SetDomain("baidu.com")
//	r := bf.GetResult()
//	log.Println(r)
//}
//
//func TestDomainBoom(t *testing.T) {
//	bf := Gsensor.DomainBoom{}
//	for {
//		bf.SetPassword("xxxxxx")
//		bf.SetDomain("freebuf.com")
//		bf.SetType("subdomain")
//		r := bf.GetResult()
//		log.Println(r)
//	}
//}
//
//func TestWebapp(t *testing.T) {
//	log.SetFlags(log.Lshortfile | log.LstdFlags)
//
//	//bf := Gsensor.WappalyzerGo{}
//	//bf.SetType("finger")
//	//bf.SetDomain("http://baidu.com")
//	//r := bf.GetResult()
//	//log.Println(r)
//	//return
//	//var h http.Transport
//	//h = http.Transport{}
//	bf := Gsensor.WappalyzerGo{}
//	for i := 0; i < 10000; i++ {
//		bf.SetType("finger")
//		bf.SetDomain("https://www.freebuf.com")
//		//bf.Http.HttpTransport = &h
//		r := bf.GetResult()
//		log.Println(r)
//	}
//	select {}
//
//}
//func TestMas(t *testing.T) {
//	mas := Gsensor.MassScan{}
//	mas.SetType("ports")
//	mas.SetDomain("1.1.1.1")
//	mas.GetResult()
//}
//func TestQiYe(t *testing.T) {
//	qiye := Gsensor.AiQiCha{}
//	qiye.SetType("qiye")
//	qiye.SetDomain("中国电信集团有限公司")
//	qiye.SetAccount("AiQiCha cookie")
//	ret := qiye.GetResult()
//	log.Println(ret)
//}
//
//func TestQiYeHold(t *testing.T) {
//	qiye := Gsensor.AiQiCha{}
//	qiye.SetType("qiye_hold")
//	qiye.SetDomain("28696963178919")
//	qiye.SetAccount("AiQiCha cookie")
//	ret := qiye.GetResult()
//	log.Println(ret)
//}
//
//func TestSub(t *testing.T) {
//	log.SetFlags(log.Lshortfile | log.LstdFlags)
//	sub := Gsensor.NewKSubDomainSensor()
//	sub.SetType("baidu.com")
//	sub.SetType("subdomain")
//	log.Println(sub.GetResult())
//}
