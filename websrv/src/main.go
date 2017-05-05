/*************************************************************************
# File Name: main.go
# Author: lizao2
# Mail: lizao2@163.com    
# Created Time: 2017-04-21 16:32:41
# Last modified: 2017-04-21 18:16:01
************************************************************************/
package main

import "fmt"
import "runtime"
import "net/http"
import "github.com/gorilla/mux"
import "logger"

func main() {
	fmt.Println(runtime.NumCPU())

	r := mux.NewRouter()

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", *listenHost, *listenPorti), r); err != nil {
		logger.PostFataMsg("ListenAndServe %s:%d fails.errMsg:%v\n", *listenHost, *listenPort, err)
	}
}
