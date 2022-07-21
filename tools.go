package tools

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GetHeader(header string) http.Header {

	HeaderMap := FormatStr(header)
	HEADER := make(http.Header)
	HEADER.Set("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36`)
	for i, j := range HeaderMap {
		HEADER.Set(i, j)
	}

	return HEADER
}

func RequestClient(URL string, METHOD string, HEADER string, DATA string) string {
	HeaderMap := FormatStr(HEADER)
	DataMap := FormatStr(DATA)
	client := &http.Client{}
	if METHOD == "get" {
		METHOD = http.MethodGet
	} else if METHOD == "post" {
		METHOD = http.MethodPost

	}
	FormatData := ""
	for i, j := range DataMap {
		FormatData = FormatData + i + "=" + j + "&"
	}
	if FormatData != "" {
		FormatData = FormatData[:len(FormatData)-1]
	}
	requset, _ := http.NewRequest(
		METHOD,
		URL,
		strings.NewReader(FormatData),
	)
	if METHOD == http.MethodPost {
		requset.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	requset.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Safari/537.36")
	for i, j := range HeaderMap {
		requset.Header.Set(i, j)
	}
	resp, _ := client.Do(requset)
	body_bit, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(body_bit)

}
func Re(Str string, str string) []string {
	digitsRegexp := regexp.MustCompile(str)
	array := digitsRegexp.FindStringSubmatch(Str)
	return array

}
func GetUUID() string {

	u1 := uuid.NewV4()
	return u1.String()
}
func WriteFile(name string, str string) {
	ystr := []byte(str)
	ioutil.WriteFile(name, ystr, 0666)
}
func FormatStr(jsonstr string) map[string]string {
	DataMap := make(map[string]string)
	Nslice := strings.Split(jsonstr, "\n")
	for i := 0; i < len(Nslice); i++ {
		if strings.Index(Nslice[i], ":") != -1 {

			a := strings.TrimSpace(Nslice[i][:strings.LastIndex(Nslice[i], ":")])
			b := strings.TrimSpace(Nslice[i][strings.LastIndex(Nslice[i], ":")+1:])
			c := strings.Trim(a, "\"")
			d := strings.Trim(b, "\"")
			DataMap[c] = d
		}
	}
	return DataMap

}
