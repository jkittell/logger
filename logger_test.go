package main

import (
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jkittell/data/api/client"
	"testing"
)

func TestLogHandler_WriteLog(t *testing.T) {
	log := Log{
		Name: "test",
		Data: gofakeit.Email(),
	}

	data, _ := json.Marshal(log)
	res, err := client.Post("http://127.0.0.1:/log", nil, bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(res))
}

func BenchmarkLogHandler_WriteLog(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		log := Log{
			Name: "test",
			Data: gofakeit.Email(),
		}

		data, _ := json.Marshal(log)
		_, err := client.Post("http://127.0.0.1:/log", nil, bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
	}
}
