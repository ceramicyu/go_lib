package main

import (
	"net/http"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/axgle/mahonia"
)

func reg(result string)  {
	rel:=regexp.MustCompile("<tr>(?s:(.*?))class=\"keyword\">(?s:(.*?))</a></span></td>(?s:(.*?))<td><span(?s:(.*?))><a(?s:(.*?))target=\"_blank\">(?s:(.*?))</a></span></td>(?s:(.*?))<td class=\"colorize\" (?s:(.*?))>(?s:(.*?))</td>(?s:(.*?))class=\"colorize\">(?s:(.*?))</td>(?s:(.*?))>(?s:(.*?))</td>")
	//<td style="text-align:left" class="colorize"><a target="_blank" href="http://vip.stock.finance.sina.com.cn/q/go.php/vInvestConsult/kind/singleqgqp/index.phtml?symbol=603579" title="点击查看该股历史千股千评">均线空头排列，等待趋势反转</a></td>
	//<td class="colorize" id="price_sh603579_14">69.99</td>
	//<td class="colorize">2.91</td>
	//<td class="colorize" id="change_sh603579_14">4.338</td>
	//<td class="colorize">67.08</td>
	//<td class="colorize">66.2</td>
	//<td class="colorize">70.1</td>
	//<td class="colorize">65.8</td>
	//<td>61.25</td>
	//<td>4208.68</td>
	//</tr>")
	result1:=rel.FindAllStringSubmatch(result,100)
	f,_:= os.Create("./wstudy/222.txt")
	defer f.Close()
	for i:=0;i<len(result1);i++{
		f.WriteString(result1[i][2]+"===>")
		f.WriteString(result1[i][6]+"===>")
		f.WriteString(result1[i][9]+"===>")
		f.WriteString(result1[i][11]+"===>")
		f.WriteString(result1[i][13]+"\r\n")
		db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/lk?charset=utf8")
		if err != nil {
			fmt.Println(err)
		}

		//关闭数据库，db会被多个goroutine共享，可以不调用
		defer db.Close()
		//插入一行数据
		ret3,err:=db.Exec("insert into share(`num`,`name`,`price`,`price_a`,`change`) values( ? , ? , ? , ? , ? )",result1[i][2], result1[i][6],result1[i][9],result1[i][11],result1[i][13] )
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
	fmt.Println("结果：\n",result1)
      fmt.Println("over >>>>")
}
func worker(url string){
	req,err:=http.Get(url)
	if err !=nil {
		fmt.Println("http Get error >>> ",err)
	}
	//fmt.Println("result: \n",req)
	buf:=make([]byte,1024*4)
	result := ""
	for{
		n,_ :=	req.Body.Read(buf)
		if n==0{
			fmt.Println(">>>读取完成>>>  \n")
			break
		}
		result += string(buf[:n])
	}
	res ,_:= os.Create("./2222.html")
    srcString:= mahonia.NewDecoder("gbk")
	srcResult:=srcString.ConvertString(result)
	tagString:=mahonia.NewDecoder("utf-8")
	_,tagresult,_:=tagString.Translate([]byte(srcResult),true)
	result=string(tagresult)
	reg(result)
	res.WriteString(result)
	defer res.Close()
}
func main()  {


	url:="http://vip.stock.finance.sina.com.cn/q/go.php/vInvestConsult/kind/qgqp/index.phtml?p="


//ch:=make(chan string)
for i:=1;i<35;i++{
	//go func(ch chan<-string) {
		worker(url+strconv.Itoa(i))
		//ch<-"ok"
	//}(ch)
	fmt.Println(i,"--结束》》》》》》》》》》》》》》》》》》》》》》》》》》》")
	continue
}
fmt.Println("结束》》》》》》》》》》》》》》》》》》》》》》》》》》》")
//for i:=1;i<10;i++{
//	<-ch
//}

}



