package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/14-bits/crowd/pkg/sink"
	"github.com/julienschmidt/httprouter"
)

var port int
var endpoint string
var sinkType string
var sinkUrl string

func init() {
	flag.IntVar(&port, "port", 2104, "HTTP Server Port")
	flag.StringVar(&endpoint, "endpoint", "/api/foo", "HTTP path to receive POST requests")
	flag.StringVar(&sinkType, "sink-type", "void", "Sink adapter, e.g.: void, sqs, redis...")
	flag.StringVar(&sinkUrl, "sink-url", "void", "Connection String/URL for the the sink")
	flag.Parse()
}

type Sink interface {
	Push(payload []byte) error
}

type Crowd struct {
	s Sink
}

func selectSink(name string, url string) (Sink, error) {
	switch name {
	case "sqs":
		return sink.NewSQS(url), nil
	case "void":
		return sink.NewVoid(url), nil
	default:
		return nil, errors.New("sink is not supported")
	}
}

func (c *Crowd) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = c.s.Push(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	s, err := selectSink(sinkType, sinkUrl)
	if err != nil {
		log.Fatal(err)
	}

	crowd := Crowd{s}

	router := httprouter.New()
	router.POST(endpoint, crowd.Handle)

	address := fmt.Sprintf(":%d", port)
	log.Println("crowd is running at", address, "with", endpoint, "-->", sinkType)

	http.ListenAndServe(address, router)
}
