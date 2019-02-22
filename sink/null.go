package sink

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Null struct{}

func NewNull() *Null {
	return &Null{}
}

func (n *Null) Push(reader io.Reader) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
