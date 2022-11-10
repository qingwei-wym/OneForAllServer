package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	oneforall "oneforall/go"
	"os"

	"github.com/gin-gonic/gin"
)

var tokenMapSlice = make(map[string]string, 50)

func valida(token string, domain string) error {

	if len(token) == 0 || len(domain) == 0 {
		return errors.New("token , domain not found")
	}

	value, ok := tokenMapSlice[token]
	if !ok {
		return errors.New("token error")
	}

	println(value)
	return nil
}

func startScanDomain(token string, domain string) error {

	err := valida(token, domain)
	if err != nil {
		return err
	}

	go oneforall.Oneforall(domain)

	return nil
}

type Subdomain struct {
	Id        string `json:"id"`
	Alive     string `json:"alive"`
	Request   string `json:"request"`
	Resolve   string `json:"resolve"`
	Url       string `json:"url"`
	Subdomain string `json:"subdomain"`
	Level     string `json:"level"`
	Cname     string `json:"cname"`
	Ip        string `json:"ip"`
	Public    string `json:"public"`
	Cdn       string `json:"cdn"`
	Port      string `json:"port"`
	Status    string `json:"status"`
	Reason    string `json:"reason"`
	Title     string `json:"title"`
	Banner    string `json:"banner"`
	Cidr      string `json:"cidr"`
	Asn       string `json:"asn"`
	Org       string `json:"org"`
	Addr      string `json:"addr"`
	Isp       string `json:"isp"`
	Source    string `json:"source"`
}

func getDomain(token string, domain string) ([]byte, error) {
	file, err1 := os.Open("results/talentsec.cn.csv")
	if err1 != nil {
		fmt.Println(err1)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	var subdomains []Subdomain
	if len(records) != 0 {
		for _, record := range records[1:] {
			subdomain := Subdomain{
				Id:        record[0],
				Alive:     record[1],
				Request:   record[2],
				Resolve:   record[3],
				Url:       record[4],
				Subdomain: record[5],
				Level:     record[6],
				Cname:     record[7],
				Ip:        record[8],
				Public:    record[9],
				Cdn:       record[10],
				Port:      record[11],
				Status:    record[12],
				Reason:    record[13],
				Title:     record[14],
				Banner:    record[15],
				Cidr:      record[16],
				Asn:       record[17],
				Org:       record[18],
				Addr:      record[19],
				Isp:       record[20],
				Source:    record[21],
			}
			subdomains = append(subdomains, subdomain)
		}
	}
	resBytes, _ := json.Marshal(subdomains)
	err := valida(token, domain)
	if err != nil {
		return []byte{}, err
	}

	//读取文件

	return resBytes, err
}

func main() {

	// getDomain("", "")
	// fmt.Printf("\n\n\n\n----启动-----")

	tokenMapSlice["122"] = "111"

	router := gin.Default()
	//定义路径
	router.GET("/start", func(context *gin.Context) {

		token := context.Query("token")
		domain := context.Query("domain")

		err := startScanDomain(token, domain)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"ok":   false,
				"data": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": "开启成功",
		})
	})

	router.GET("/getDomain", func(context *gin.Context) {

		token := context.Query("token")
		domain := context.Query("domain")

		data, err := getDomain(token, domain)
		fmt.Printf("data: %v\n", string(data))
		if err != nil {
			context.String(200, string(data))
			// context.JSON(http.StatusBadRequest, gin.H{
			// 	"ok":   false,
			// 	"data": err.Error(),
			// })
			return
		}
		// var ass string
		// err1 := json.Unmarshal(data, &ass)
		// if err1 != nil {
		// 	fmt.Printf("err.Error(): %v\n", err.Error())
		// }
		// fmt.Printf("ass: %v\n", ass)
		// context.JSON(http.StatusOK, gin.H{
		// 	"ok":   true,
		// 	"data": "",
		// })
		temp := fmt.Sprintf(`{"ok":   true,"data": %v}`, string(data))
		context.String(200, temp)

	})
	router.Run()
}
