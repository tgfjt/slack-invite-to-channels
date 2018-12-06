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

type ChannelsList struct {
	Ok       bool      `json:"ok"`
	Error    string    `json:"error"`
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type InviteChannel struct {
	Ok      bool    `json:"ok"`
	Error   string  `json:"error"`
	Channel Channel `json:"channel"`
}

func (list *ChannelsList) Find(x string) int {
	for i, n := range list.Channels {
		if x == n.Name {
			return i
		}
	}
	return len(list.Channels)
}

func (list *ChannelsList) GetUidByName(t string) string {
	i := list.Find(t)
	return list.Channels[i].Id
}

func GetChannels(token string) ChannelsList {
	apis := GetApis()
	var list ChannelsList
	var once sync.Once

	once.Do(func() {
		c := utils.GetClient()
		d := url.Values{}
		d.Set("token", token)

		req, _ := http.NewRequest("POST", apis.ChannelsApi, bytes.NewBufferString(d.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := c.Do(req)

		if err != nil {
			log.Fatalf("fail to get channels list")
		}

		json.NewDecoder(res.Body).Decode(&list)

		if !list.Ok {
			log.Fatalf("fail to get channels list: %v", list.Error)
		}
	})

	return list
}

func Invite(token string, uid string, cuid string) {
	apis := GetApis()
	c := utils.GetClient()
	d := url.Values{}
	d.Set("token", token)
	d.Set("user", uid)
	d.Set("channel", cuid)

	req, _ := http.NewRequest("POST", apis.InviteApi, bytes.NewBufferString(d.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)

	var channel InviteChannel

	if err != nil {
		log.Fatalf("fail to get channels list")
	}

	json.NewDecoder(res.Body).Decode(&channel)

	if !channel.Ok {
		log.Fatalf("fail to get channels list: %v", channel.Error)
	}
}
