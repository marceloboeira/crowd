package sink

import "fmt"

type Void struct {
	Debug bool
}

func NewVoid(url string) Void {
	return Void{Debug: true}
}

func (v Void) Push(payload []byte) error {
	if v.Debug {
		fmt.Println("debug: ", string(payload))
	}

	return nil
}
