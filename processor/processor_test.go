package processor

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PjHiK7t5J/task-onefootball/vintagemonster"
)

var testResponse = `
{
	"status": "ok",
	"code": 0,
	"data": {
	  "team": {
		"id": 6,
		"optaId": 0,
		"name": "FC Bayern Munich",
		"players": [
		  {
			"id": "149",
			"country": "Spain",
			"firstName": "Thiago",
			"lastName": "Alc\u00e1ntara do Nascimento",
			"name": "Thiago Alc\u00e1ntara",
			"position": "Midfielder",
			"number": 6,
			"birthDate": "1991-04-11",
			"age": "27",
			"height": 174,
			"weight": 70,
			"thumbnailSrc": "https:\/\/image-service.onefootball.com\/resize?fit=crop&h=180&image=https%3A%2F%2Fimages.onefootball.com%2Fplayers%2F149.jpg&q=75&w=180"
		  },
		  {
			"id": "168",
			"country": "France",
			"firstName": "Franck",
			"lastName": "Rib\u00e9ry",
			"name": "Franck Rib\u00e9ry",
			"position": "Midfielder",
			"number": 7,
			"birthDate": "1983-04-07",
			"age": "35",
			"height": 170,
			"weight": 72,
			"thumbnailSrc": "https:\/\/image-service.onefootball.com\/resize?fit=crop&h=180&image=https%3A%2F%2Fimages.onefootball.com%2Fplayers%2F168.jpg&q=75&w=180"
		  }
		]
	  }
	}
  }
`

func TestProcess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, testResponse)
	}))
	defer ts.Close()

	vmonster := vintagemonster.New()
	vmonster.SetHost(ts.URL)

	proc := Processor{vmonster: vmonster}

	list, err := proc.Process(map[string]bool{
		"FC Bayern Munich": true,
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := "Franck Rib\u00e9ry"
	if list[0].Name != expected {
		t.Fatalf("Expected name %q, got %q", expected, list[0].Name)
	}

}

func TestPrint(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, testResponse)
	}))
	defer ts.Close()

	vmonster := vintagemonster.New()
	vmonster.SetHost(ts.URL)

	var b []byte
	buff := bytes.NewBuffer(b)
	proc := Processor{vmonster: vmonster, writer: buff}

	list, err := proc.Process(map[string]bool{
		"FC Bayern Munich": true,
	})
	if err != nil {
		t.Fatal(err)
	}

	proc.Print(list)

	expected := "1. Franck Ribéry; 35; FC Bayern Munich\n2. Thiago Alcántara; 27; FC Bayern Munich\n"

	if buff.String() != expected {
		t.Fatal("Output is not correct")
	}
}
