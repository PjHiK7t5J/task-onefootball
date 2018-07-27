package vintagemonster

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Vintagemonster is a REST service provided by Onefootball
// to fetch information about teams.
type Vintagemonster struct {
	host string
}

// New return new instance of Vintagemonster
func New() *Vintagemonster {
	return &Vintagemonster{
		host: "https://vintagemonster.onefootball.com",
	}
}

// Response describe structure of the response for team endpoint
type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		Team struct {
			Name    string `json:"name"`
			Players []struct {
				Name string `json:"name"`
				Age  string `json:"age"`
			} `json:"players"`
		} `json:"team"`
	} `json:"data"`
}

// SetHost sets the host of Vintagemonster service
func (v *Vintagemonster) SetHost(host string) {
	v.host = host
}

// Get returns Response for the team endpoint, error if failed
func (v *Vintagemonster) Get(id int) (res Response, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/teams/en/%d.json", v.host, id))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	return
}
