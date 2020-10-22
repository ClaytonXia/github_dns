package main

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

func main() {
	getFastlyAddr()
	getAssetAddr()
	getGithubAddr()
	fmt.Println("dont't forget to restart networking")
}

func getFastlyAddr() {
	url := "https://fastly.net.ipaddress.com/github.global.ssl.fastly.net#ipinfo"
	addr := queryDOC(url, "/html/body/div[1]/main/section[1]/table/tbody/tr[2]/td/ul/li")
	fmt.Printf("%s github.global.ssl.fastly.net\n", addr)
}

func getAssetAddr() {
	url := "https://github.com.ipaddress.com/assets-cdn.github.com"
	addr := queryDOC(url, "/html/body/div[1]/main/section[2]/table/tbody/tr[1]/td/div/ul/li[1]/a")
	fmt.Printf("%s assets-cdn.github.com\n", addr)
}

func getGithubAddr() {
	url := "http://github.com.ipaddress.com/#ipinfo "
	addr := queryDOC(url, "/html/body/div[1]/main/section[1]/table/tbody/tr[6]/td/ul/li")
	fmt.Printf("%s github.com\n", addr)
}

func queryDOC(url, xpathExpr string) (res string) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		panic(err)
	}

	exprNode := htmlquery.FindOne(doc, xpathExpr)
	res = htmlquery.InnerText(exprNode)

	return
}
