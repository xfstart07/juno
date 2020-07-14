package install

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/douyu/juno/pkg/model/db"
	"github.com/labstack/echo/v4"
)

func mockApp(url string, router *echo.Echo) {
	file, err := ioutil.ReadFile("./data/mockdata/app.json")
	if err != nil {
		panic(err)
	}
	param := string(file)
	// 发起post请求，以表单形式传递参数
	body := PostForm(url, param, router)
	fmt.Println(string(body))
}

func mockAppNode(url string, router *echo.Echo) {
	mockAppNodeByFile("./data/mockdata/appnode1.json", url, router)
	mockAppNodeByFile("./data/mockdata/appnode2.json", url, router)
	mockAppNodeByFile("./data/mockdata/appnode3.json", url, router)
	mockAppNodeByFile("./data/mockdata/appnode4.json", url, router)
}

func mockAppNodeByFile(fileName, url string, router *echo.Echo) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	param := string(file)
	// 发起post请求，以表单形式传递参数
	body := PostForm(url, param, router)
	fmt.Println(string(body))
}

func mockTplCreate(url string, router *echo.Echo) {
	mockTplByFile("mysql", "./data/mockdata/configtpl1.json", url, router)
	mockTplByFile("mysql", "./data/mockdata/configtpl2.json", url, router)
	mockTplByFile("redis", "./data/mockdata/configtpl3.json", url, router)
	mockTplByFile("grpc", "./data/mockdata/configtpl4.json", url, router)
}

func mockTplByFile(tplType, fileName string, url string, router *echo.Echo) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	req := map[string]interface{}{
		"tpl_type": tplType,
		"content":  string(file),
	}
	paramByte, _ := json.Marshal(req)

	// 发起post请求，以表单形式传递参数
	PostForm(url, string(paramByte), router)
}

func mockConfig(urlFile string, urlContent string, router *echo.Echo) {
	mockConfigByFile("./data/mockdata/appconfigfile1.json", "./data/mockdata/appconfigcontent1.toml", urlFile, urlContent, router)
	mockConfigByFile("./data/mockdata/appconfigfile2.json", "./data/mockdata/appconfigcontent2.toml", urlFile, urlContent, router)
	mockConfigByFile("./data/mockdata/appconfigfile3.json", "./data/mockdata/appconfigcontent3.toml", urlFile, urlContent, router)
	mockConfigByFile("./data/mockdata/appconfigfile4.json", "./data/mockdata/appconfigcontent4.toml", urlFile, urlContent, router)
}

func mockConfigByFile(configFile string, configContent string, urlFile string, urlContent string, router *echo.Echo) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	// 发起post请求，以表单形式传递参数
	body := PostForm(urlFile, string(file), router)
	fmt.Println("urlFile------>", configFile)
	fmt.Println(string(body))
	resp := struct {
		Data db.CmcApp `json:"data"`
	}{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		panic(err)
	}

	file, err = ioutil.ReadFile(configContent)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
	param1 := map[string]interface{}{
		"caid":  resp.Data.ID,
		"key":   "default",
		"value": string(file),
	}

	tt, _ := json.Marshal(param1)
	// 发起post请求，以表单形式传递参数
	body = PostForm(urlContent, string(tt), router)
	fmt.Println(string(body))
}

func mockParse(url string, router *echo.Echo) {
	// 发起post请求，以表单形式传递参数
	PostForm(url, "", router)
}

func mockCreateUser(url string, router *echo.Echo) {
	// 发起post请求，以表单形式传递参数
	param := `{
	"username": "admin",
	"nickname": "admin",
	"password": "21232f297a57a5a743894a0e4a801fc3",
	"access": "admin"
}`

	body := PostForm(url, param, router)
	fmt.Println(string(body))

	PostForm(url,
		`{"username": "default", "nickname": "default","password": "c21f969b5f03d33d43e04f8f136e7682","access": "user"}`,
		router,
	)
}

func mockMonitorAPI(url string, router *echo.Echo) {
	// 发起post请求，以表单形式传递参数
	param := `{
		"id": 0,
	"sysType":2,
	"setCate": "API",
	"setStr": "http://jupiterconsole.douyu.com/grafana/d/api"
}`
	body := PostForm(url, param, router)
	fmt.Println(string(body))
}

func mockMonitorInstance(url string, router *echo.Echo) {
	// 发起post请求，以表单形式传递参数
	param := `{
		"id": 0,
		"sysType":2,
	"setCate": "Instance",
	"setStr": "http://jupiterconsole.douyu.com/grafana/d/instance"
}`
	body := PostForm(url, param, router)
	fmt.Println(string(body))
}

func mockMonitorOverview(url string, router *echo.Echo) {
	// 发起post请求，以表单形式传递参数
	param := `{
	"id": 0,
	"sysType":2,
	"setCate": "Overview",
	"setStr": "http://jupiterconsole.douyu.com/grafana/d/overview"
}`
	body := PostForm(url, param, router)
	fmt.Println(string(body))

}
