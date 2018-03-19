package main

import (
	"fmt"
	"strconv"
	"os"
	"net/http"
	"regexp"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



func HttpGet2(url string)(result string,err error)  {
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


func DoWork2(ip string){
	for k:=1;k<2;k++{
		url:="https://www.baidu.com/s?wd="+ip
		//url:="https://slamdunk.sports.sina.com.cn/team/stats?tid=583ecb3a-fb46-11e1-82cb-f4ce4684ea4c"
		//url:="http://www.27270.com/ent/meinvtupian/list_11_"+strconv.Itoa(k)+".html"

		result,err:=HttpGet2(url)
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

		//(?s:(.*?))
		rel:=regexp.MustCompile("<span class=\"c-gap-right\">IP地址:&nbsp;(?s:(.*?))</span>(?s:(.*?)) (?s:(.*?))\t")
		//rel:=regexp.MustCompile("<(?s:(.*?))>")
		result1:=rel.FindAllStringSubmatch(result,100000)

		for i:=0;i<len(result1);i++{
			f.WriteString(result1[i][1]+"====>>")
			f.WriteString(result1[i][1]+"\r\n")
			fmt.Println(result1[i][1])
			fmt.Println(result1[i][2])
			fmt.Println(result1[i][3])

			db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/lk?charset=utf8");
			if err != nil {
				fmt.Println(err);
			}

			//关闭数据库，db会被多个goroutine共享，可以不调用
			defer db.Close();
			//插入一行数据
			ret3,err:=db.Exec("insert into ip2(ip,location,type) values( ? ,? ,? )",result1[i][1],result1[i][2],result1[i][3])
			if err != nil {
				fmt.Println(">>>>>>>",err)
			}
			//id := 0;
			//获取插入ID
			//_ , err1:= ret.LastInsertId()
			//fmt.Println(err1)
			//获取影响行数

			del_nums, _ := ret3.RowsAffected()
			fmt.Println(del_nums)
		}

		f.Close()

	}

	return
}


func main(){

	//DoWork2("0.0.0.0")
	fmt.Println("第 1 次请求")
	for a:=1;a<255;a++ {
		for b := 0; b < 255; b++ {
			for c := 0; c < 255; c++ {
				for d := 0; d < 255; d++ {
					DoWork2(strconv.Itoa(a)+"."+strconv.Itoa(b)+"."+strconv.Itoa(c)+"."+strconv.Itoa(d))
				}
			}
		}

	}



}
