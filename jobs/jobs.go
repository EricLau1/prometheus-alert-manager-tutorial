package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"prometheus-alert-manager-tutorial/api/httpext"
	"prometheus-alert-manager-tutorial/api/types"
	"prometheus-alert-manager-tutorial/api/utils"
	"time"
)

const baseURI = "http://localhost:8080/todos"

func createTodo() {
	todo := &types.Todo{
		Title:       faker.Word(),
		Description: faker.Paragraph(),
		Done:        faker.RandomUnixTime()%2 == 0,
	}
	b, err := json.Marshal(todo)
	if err != nil {
		log.Println("unable to encode Todo to JSON:", err.Error())
		return
	}
	_, _ = doRequest(http.MethodPost, baseURI, bytes.NewReader(b))
}

func getTodos() []*types.Todo {
	var todos []*types.Todo
	body, _ := doRequest(http.MethodGet, baseURI, nil)
	err := json.Unmarshal(body, &todos)
	if err != nil {
		log.Println("unable to decode todos:", err.Error())
		return nil
	}
	return todos
}

func updateTodo() {
	todos := getTodos()
	if len(todos) > 0 {
		var (
			index = utils.Random().Intn(len(todos))
			todo  = todos[index]
			uri   = baseURI + fmt.Sprintf("/%d", todo.ID)
		)
		todo.Done = !todo.Done
		todo.Title = faker.Word()
		todo.Description = faker.Paragraph()
		b, err := json.Marshal(todo)
		if err != nil {
			log.Println("unable to encode Todo to JSON:", err.Error())
			return
		}
		_, _ = doRequest(http.MethodPut, uri, bytes.NewReader(b))
	}
}

func deleteTodo() {
	todos := getTodos()
	if len(todos) > 0 {
		var (
			index = utils.Random().Intn(len(todos))
			todo  = todos[index]
			uri   = baseURI + fmt.Sprintf("/%d", todo.ID)
		)
		_, _ = doRequest(http.MethodDelete, uri, nil)
	}
}

func doRequest(method string, uri string, body io.Reader) ([]byte, string) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		log.Println("unable to create request:", err.Error())
		return nil, ""
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("request failed:", err.Error())
		return nil, ""
	}
	if res.StatusCode >= http.StatusBadRequest {
		httpext.Dump(req, res)
		return nil, res.Status
	}
	defer utils.HandleClose(res.Body)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("unable to read response body:", err.Error())
	}
	return b, res.Status
}

type Job struct {
	stop chan struct{}
}

func New() *Job {
	return &Job{stop: make(chan struct{})}
}

func (j *Job) Close() {
	j.stop <- struct{}{}
}

func (j *Job) Run() {
	go onLoop(createTodo, time.Millisecond*100, j.stop)
	go onLoop(updateTodo, time.Millisecond*200, j.stop)
	go onLoop(deleteTodo, time.Second, j.stop)
}

func onLoop(fn func(), wait time.Duration, stop <-chan struct{}) {
	time.Sleep(wait)
Loop:
	for {
		select {
		case _ = <-stop:
			break Loop
		default:
			fn()
			time.Sleep(wait)
		}
	}
}
