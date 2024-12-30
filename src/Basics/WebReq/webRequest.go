package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	WebReq()
}
func WebReq() {
	fmt.Println("Web Services Handling")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while Geting Responce")
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Type of resp: %T\n", resp)
	// fmt.Println(resp)

	//Read responce body

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading Responce: ", err)
		return
	}
	fmt.Println("Response: ", string(data))

}
