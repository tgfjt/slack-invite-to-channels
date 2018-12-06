package lib

import (
	"sync"
)

type Apis struct {
	MembersApi  string
	ChannelsApi string
	InviteApi   string
}

var once sync.Once
var apis *Apis

// TODO: we just need some const??
func GetApis() *Apis {
	once.Do(func() {
		apis = &Apis{
			MembersApi:  "https://slack.com/api/users.list",
			ChannelsApi: "https://slack.com/api/channels.list",
			InviteApi:   "https://slack.com/api/channels.invite",
		}
	})

	return apis
}
