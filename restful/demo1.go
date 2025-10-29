package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo() {

	var todo Todo
	url := "https://jsonplaceholder.typicode.com/todos/2"
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	dc := json.NewDecoder(response.Body)
	error := dc.Decode(&todo)
	if error != nil {
		fmt.Println("Dönüştüre başarısız", error)
		return
	}
	fmt.Println(todo)

}

func Demo2() {
	// todo := Todo{UserID: 1, ID: 1, Title: "alışveriş yapılacak", Completed: false}
	// jsonTodo, err := json.Marshal(todo)
	url := "https://jsonplaceholder.typicode.com/todos"
	// byteTodo := bytes.NewBuffer(jsonTodo)
	// if err != nil {
	// 	fmt.Println("Dönüştüre başarısız", err)
	// 	return
	// }

	// resp, err2 := http.Post(url, "application/json", byteTodo)
	// if err2 != nil {
	// 	fmt.Println("Dönüştüre başarısız", err2)
	// 	return
	// }
	// fmt.Println(resp)

	todo := Todo{UserID: 1, ID: 1, Title: "alışveriş yapılacak", Completed: false}

	client := http.Client{}

	b, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(req)
	defer response.Body.Close()
}
