package lib

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/tgfjt/slack-invite-to-channels/utils"
)

type Members struct {
	Ok      bool     `json:"ok"`
	Error   string   `json:"error"`
	Members []Member `json:"members"`
}

type Member struct {
	Id      string  `json:"id"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	Name string `json:"display_name"`
}

func (members *Members) Find(x string) int {
	for i, n := range members.Members {
		if x == n.Profile.Name {
			return i
		}
	}
	return len(members.Members)
}

func (members *Members) GetUidByName(t string) string {
	i := members.Find(t)
	return members.Members[i].Id
}

func GetMembers(token string) Members {
	apis := GetApis()
	var members Members
	var once sync.Once

	once.Do(func() {
		c := utils.GetClient()
		d := url.Values{}
		d.Set("token", token)

		req, _ := http.NewRequest("POST", apis.MembersApi, bytes.NewBufferString(d.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := c.Do(req)

		if err != nil {
			log.Fatalf("fail to get list")
		}

		json.NewDecoder(res.Body).Decode(&members)

		if !members.Ok {
			log.Fatalf("fail to get members: %v", members.Error)
		}
	})

	return members
}
