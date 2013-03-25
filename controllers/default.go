package controllers

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/goredis"
)

const (
	url_encode_prefix = "kissgo"
	url_encode_suffix = "gokiss"
	url_domain        = "http://d.rzj.me"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	this.Ctx.Request.ParseForm()
	original_url := this.Ctx.Request.Form.Get("url")
	if !CheckUrl(original_url) {
		this.Data["json"] = map[string]string{"state": "0", "message": "url格式错误"}
		this.ServeJson()
		return
	}
	short_urls := ShortUrl(original_url, url_encode_prefix, url_encode_suffix)
	short_url := url_domain + "/" + short_urls[0]
	//存储地址到redis
    var client goredis.Client
    val, _ := client.Get(short_urls[0])
    if string(val)!="" {
        this.Data["json"] = map[string]string{"state": "1", "short_url": short_url}
		this.ServeJson()
        return
    }
    
    client.Set(short_urls[0], []byte(original_url))
    client.Set(short_urls[1], []byte(original_url))
    client.Set(short_urls[2], []byte(original_url))
    client.Set(short_urls[3], []byte(original_url))
	
	this.Data["json"] = map[string]string{"state": "1", "short_url": short_url}
	this.ServeJson()
	return
}

func ShortUrl(url string, prefix string, suffix string) [4]string {

	str := prefix + url + suffix

	chars := [62]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
		'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
		'y', 'z', '0', '1', '2', '3', '4', '5',
		'6', '7', '8', '9', 'A', 'B', 'C', 'D',
		'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z'}

	md5str := Md5Encode(str)
	hexstr := []byte(md5str)

	var resUrl [4]string

	for i := 0; i < 4; i++ {
		j := i * 8
		k := j + 8
		s := "0x" + string(hexstr[j:k])
		hexInt, _ := strconv.ParseInt(s, 0, 64)
		hexInt = 0x3FFFFFFF & hexInt

		var outChars string = ""
		for n := 0; n < 6; n++ {
			index := 0x0000003D & hexInt
			outChars = outChars + string(chars[index])
			hexInt = hexInt >> 5
		}
		resUrl[i] = outChars
	}
	return resUrl
}

func Md5Encode(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", h.Sum(nil))
	return buffer.String()
}

func CheckUrl(url string) bool {
	if ok, _ := regexp.MatchString("^(https?:\\/\\/)?[A-Za-z0-9]+\\.[A-Za-z0-9]+[\\/=\\?%\\-&_~`@[\\]\\':+!]*([^<>\"\"])*$", url); !ok {
		return false
	}
	return true
}
