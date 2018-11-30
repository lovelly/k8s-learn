package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"k8s_learn/utils"
	"net/http"
	"os"
)



func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := bytes.Buffer{}
		data.WriteString("v1 ")
		data.WriteString("hello world ")
		data.WriteString(utils.GetSelfAddr())
		data.WriteString("\n")
		env:=os.Getenv("TEST_ENV")
		data.WriteString("env: ")
		data.WriteString(env)
		data.WriteString("\n")
		data.WriteString("secret: ")
		data.WriteString("\n")
		data.WriteString("username: ")
		b,_ :=ioutil.ReadFile("/etc/secret/username")
		data.Write(b)
		data.WriteString("\n")
		data.WriteString("password: ")
		b,_ =ioutil.ReadFile("/etc/secret/password")
		data.Write(b)
		data.WriteString("\n")
		data.WriteString("labels: ")
		b,_ =ioutil.ReadFile("/etc/podinfo/labels")
		data.Write(b)
		data.WriteString("\n")

		_, err := writer.Write(data.Bytes())
		if err != nil {
			fmt.Println("write error: ", err)
		}
	})
	http.ListenAndServe(":8787", nil)
}
