package conf

import (
"bufio"
"io"
"os"
"strings"
)



type Config struct {
	Mymap  map[string]string
	strcet string
}

func (c *Config) InitConfig(path string)(map[string]map[string]string) {
	 m:= make(map[string]map[string]string)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	key:=""
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		if s==""{
			continue
		}else{
			//fmt.Println("<<<<",s)
			n1 := strings.Index(s, "[")
			n2 := strings.LastIndex(s, "]")

			if n1>-1 && n2>-1{
			//	fmt.Println("####",n1,n2,s[n1+1:n2-1])
				key=s[n1+1:n2]
				m[key]=make(map[string]string)
				//fmt.Println(">>22>>",m[key])
			}else{
				n3 := strings.Index(s, ":")
				//fmt.Println("@@@@",n3,s[0:n3],s[n3+1:])
				k:=s[0:n3]
				m[key][k]=s[n3+1:]
			//	fmt.Println(">>>>",m[key])
			}

		}

	}
	return m
}

