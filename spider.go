package main

import (
	"fmt"
	"strconv"
	"os"
	"net/http"
	"regexp"
	"io/ioutil"
	"io"
	"bytes"
	"time"
)

var n=0

func HttpGet(url string)(result string,err error)  {
	req,err1 :=http.Get(url)
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	if err1 != nil{
		fmt.Println("http get err>>>",err)
		err =err1
	}
	buf:=make([]byte,1024*40)

	for{
		n,err2:=req.Body.Read(buf)
		if n == 0{
			fmt.Println("req Body Read err >>>",err2)
			break
		}
		result +=string(buf[:n])

	}
	return
}
func SaveImage(imgurl string){
	//imgurl = "https://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
	//imgurl="https://img1.mm131.me/pic/3118/0.jpg"
	//图片正则
	n++
	//reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)

	name := "./wstudy/img/0"+strconv.Itoa(n)+".jpg"
	fmt.Print(name)
	//通过http请求获取图片的流文件
	resp, _ := http.Get(imgurl)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(name)
	io.Copy(out, bytes.NewReader(body))
	return
}

func DoWork(){
       for k:=1;k<100;k++{
	    //url:="http://trend.caipiao.163.com/cqssc/"
		//url:="http://www.27270.com/ent/meinvtupian/"
		url:="http://www.27270.com/ent/meinvtupian/list_11_"+strconv.Itoa(k)+".html"

		result,err:=HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err >>>",err)
			continue
		}
		fileName :=strconv.Itoa(12)+".html"
		f,err1 :=os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.Create err >>>",err1)
			continue
		}

		//rel := regexp.MustCompile(`<tr data-period="(?s:(.*?))" data-award="(?s:(.*?))">`)
		//rel:=regexp.MustCompile("<html>(?s:(.*?))</html>")
		rel:=regexp.MustCompile("<img src=\"(?s:(.*?))\" w")
		//rel := regexp.MustCompile(`<td><span id='qs_1' class='xs0'>(?s:(.*?))</span></td>`)
		result1:=rel.FindAllStringSubmatch(result,100000)

		for i:=0;i<len(result1);i++{
			f.WriteString(result1[i][1]+"====>>")
			f.WriteString(result1[i][1]+"\r\n")
			if i != 0{
				if result1[i][1] !="http://www.27270.comFailed to open output stream."{
					SaveImage(result1[i][1])
					continue
				}
				continue
			}
		}

		f.Close()

	   }

	return
}


func main(){
	//SaveImage("as")
	n:=0
 //for{
 	//if time.Now().Second() == 9{
 		n++;
 		fmt.Println("第 %v 次请求",n)
		DoWork()
	//}
 //}

}
