package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	http.HandleFunc("/write", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		data := " not found "
		if len(request.Form["data"]) > 0 {
			data = request.Form["data"][0]
		}
		err := ioutil.WriteFile("/data/test",[]byte(data),0644)
		if err != nil {
			fmt.Println(err)
			writer.Write([]byte(err.Error()))
		}
	})
	http.HandleFunc("/read", func(writer http.ResponseWriter, request *http.Request) {
		data, err := ioutil.ReadFile("/data/test")
		if err != nil {
			fmt.Println(err)
			writer.Write([]byte(err.Error()))
		}
		writer.Write([]byte(" read: " + string(data) + "\n"))
	})
	http.ListenAndServe(":8788", nil)
}
