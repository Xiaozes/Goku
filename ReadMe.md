#  Go Red-Team 工具开发库

---
- #### Ghttp    Http客户端
  - [x] Get/Post/Head..
  - [x] SetHeader
  - [x] Set/GetCookie
  - [x] Save To File
  - [x] Use Http/Socks5 Proxy
- ##### 使用方法
  -  #### 发送Get请求
  ```go
    httpClient=Ghttp.http{}
    httpClient.Get("https://baidu.com")
    httpClient.Execute()
  ```
  -  #### 发送Post请求
  ```go
    httpClient=Ghttp.http{}
    httpClient.Post("https://baidu.com","a=1&b=1")
    httpClient.Execute()
  ```
  ```go
    httpClient=Ghttp.http{}
    param:=url.Values{
    		"name":{"values"},
    	}
    httpClient.Post("https://baidu.com",param)
    httpClient.Execute()
  ```
  ```go
    httpClient=Ghttp.http{}
    jsonData:=make(map[string]interface{})
    jsonData["key"]="value"
    jsonData["json2"]=make(map[string]interface{})
    jsonData["json2"]["data"]="1234"
    httpClient.Post("https://baidu.com",jsonData)
    httpClient.Execute()
  ```
  
  -  #### 开启Session 自动记录cookie
  ```go
    httpClient := Ghttp.New()
    httpClient.Session()
    httpClient.New("GET","https://www.baidu.com")
    httpClient.Execute()

  ```
  -  #### 获取字符串返回值
  ```go
    responseText:=httpclient.Text()
    log.Println(responseText)
  ```
  -  #### 取Byte返回值
  ```go
    responseByte:=httpclient.Byte()
    log.Println(responseByte)
  ```
  -  #### 获取StatusCode
  ```go
    code:=httpclient.StatusCode()
    log.Println(code)
  ```
  - #### 设置代理
  ```go
  httpClient.SetProxy("http://127.0.0.1:6152")
  httpClient.SetProxy("socks5://ss:ss@127.0.0.1:6153")
  ```

---
- #### Gconvert  类型转换
  - [x] int float转string
  - [x] str 转int Float
  - [x] base64，url，raw，md5，sha1，sha256。sha512编码解码
  - [x] time跟str互转
  - [x] byte跟str互转
- ##### 使用方法
  -  #### int float转string
  ```go
    Gconvert.Int2String(1234)
  ```
  -  #### str 转int Float
  ```go
    Gconvert.Str2Int("12345")
    Gconvert.Str2Float("12345")
    Gconvert.Str2Float64("12345")
  ```
  -  ####  编码解码
   ```go
    Gconvert.B64Encode("12312312")
    Gconvert.B64Decode("324 d")
    Gconvert.UrlEncode("324=1;sd;'123 d")
    Gconvert.UrlDecode("%25%27%22")
    Gconvert.RawDecode("%25%27%22")
    Gconvert.RawEncode("324=1;sd;'123 d")
    Gconvert.Md5("123456")
    Gconvert.Sha1("123456")
    Gconvert.Sha256("123456")
    Gconvert.Sha512("123456")
  ```
  -  ####  time互转str
   ```go
   Gconvert.Time2Str(time.Now())
   Gconvert.Unix2Time(1614168000)
   Gconvert.Str2Time("2020-112-11 22:33:11")
  ```
  -  #### byte互转str 
  ```go
  Gconvert.Bytes2String("xxxx")
  Gconvert.String2Bytes("12345")
  ```
---
- ####  文件操作
- [x] 读取txt文件 返回string
- [x] 文件写入 (覆盖所有)
- [x] 文件或者目录是否存在
- [x] 读取目录下的文件信息
- ##### 使用方法
  -  ####  读取txt文件 返回string
   ```go
   Gfile.ReadFileToString("test.txt")
  ```
  -  ####   文件写入 (覆盖所有)
   ```go
  Gfile.WriteString("test.txt", "这是测试")
  ```
  -  ####  文件或者目录是否存在
   ```go
   if Gfile.CheckExist("test.txt") { 
      fmt.Println("test.txt 存在")
   } else {
		    fmt.Println("文件不存在")
   }
  ```
  -  ####  读取目录下的文件信息
  ```go
  Gfile.GetFileList("c://xxx/xxx/xxxx/xxx"))
  ```
---
- ####  各种登录操作
- [x] mysql
- [x] postgres
- [x] mssql
- [x] ftp
- [x] pop3
- [x] smtp
- [x] ssh (代理功能未完成)
- ##### 使用方法
  -  ####  数据库操作
  ```go
	 r := Glogin.SqlQuery("127.0.0.1", "3306", "root", "123123", "mysql", "select * from users limit 10")
  ```
  -  ####  ftp
   ```go
	 r := Glogin.FtpLogin("127.0.0.1", "2121", "anonymous", "anonymous")
   ```
  -  ####  pop3
   ```go
	 r := Glogin.Pop3Login("pop.126.com", "995", "123123", "123123", true)
   ```
  -  ####  smtp
   ```go
	 r := Glogin.SmtpLogin("smtp.126.com", "465", "123123", "123123", true)
   ```
  -  ####  ssh
   ```go
  Glogin.Sshlogin("password", "root", "xxxx", "IP", "ls", "", "", 22) 
   Glogin.login.Sshlogin("key", "root", "xxxx", "IP", "ls", "/xxx/id_rsa", "", 22)
  ```
  -  ####  redis
  ```go
     Glogin.RedisLogin("IP","6379","xxxxxx")
  ```
---
- ####  各种网络api调用
- [x] Fofa
- [x] Hunter
-  [x] Quake(待测试)
- [x] ZoomEye(待测试)
- [x] Shodan
- [x] Censys
- ##### 使用方法
  -  ####  Fofa
  ```go
    ff := Gapi.Fofa{}
    ff.SetAccount("XXXX@qq.com")
    ff.SetPassword("XXXXXXX")
    r := ff.GetResult("domain=\"baidu.com\"")
    for k, v := range r {
		     fmt.Println(k, v)
	   }
  ```
  -  ####  hunter
  ```go
    hh := Gapi.Hunter{}
	   hh.SetPassword("xxxxx")
	   hh.SetPage("1")
	   hh.SetNubmer("100")
	   r := hh.GetResult("ip=\"xx4.1x3.xx0.xx3\"")
	   for k, v := range r {
		    fmt.Println(k, v)
	    }
  ```
  -  ####  Shodan
  ```go
    sd := Gapi.Shodan{}
	 // shodan用的chrome插件的api，只能用来获取端口
	   sd.SetDomain("1xx.168.72.89")
	   sd.SetPassword("xxxxxxxx")
	   sd.SetType("ports")
	   rr := sd.GetResult()
  ```
  -  ####  Censys
  ```go
   ce := Gapi.Censys{}
	  ce.SetUsername("xxxxxxx")
	  ce.SetPassword("xxxxxxxxxxxxxxxxxx")
	  cresult := ce.GetResult("xiaomi.com")
	  for k, v := range cresult {
		    fmt.Println(k, v)
	  }
  ```
  
