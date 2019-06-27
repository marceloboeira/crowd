package sink

import (
	"testing"
)

type mockRedis struct {
	Counter int
}

func (m *mockRedis) RPush() (bool, error) {
	m.Counter++

	return true, nil
}

func TestRedis(t *testing.T) {
	mock := mockRedis{Counter: 0}

	s := &Redis{
		Client: &mock,
		//Url:    "http://localhost:9324/sqs/123",
	}

	s.Push([]byte("test"))
	s.Push([]byte("test"))
	s.Push([]byte("test"))

	if mock.Counter != 3 {
		t.Errorf("redis did not receive push")
	}
}
