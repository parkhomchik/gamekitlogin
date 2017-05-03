package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func main(){
	fmt.Println("start login")

	resp, err := http.Get("https://gamekit.com/giveaways/#show=login")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	str := fmt.Sprintf("%s", body)
	index := strings.Index(str, "<input type=\"hidden\" id=\"_token\" name=\"_token\" value=\"")
	x := str[index+53:index+53+45]
//<input type="hidden" id="_token" name="_token" value="wgZHa0tgmM7Kr0EGs9cVD1QamMSnp6OlP2k-u4F6FRs">

	fmt.Println(x)
}