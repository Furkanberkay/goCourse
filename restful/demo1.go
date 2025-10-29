package restful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo() {
	getReqUrl := "https://jsonplaceholder.typicode.com/todos/1/"
	response, err := http.Get(getReqUrl)
	if err != nil {
		fmt.Println("get isteği başarısız")
		return
	}
	defer response.Body.Close()

	bodyByte, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyByte)
	fmt.Println(bodyString)

	todo := []Todo{}

	dc := json.NewDecoder(response.Body)
	dc.Decode(&todo)
	fmt.Println(todo)

}
