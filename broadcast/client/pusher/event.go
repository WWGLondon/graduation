package pusher

const (
	SubscribeEvent = "pusher:subscribe"
)

type OutMessage struct {
	Event string `json:"event"`
	Data  Data   `json:"data"`
}

type InMessage struct {
	Event   string `json:"event"`
	Channel string `json:"channel,ommitempty"`
	Data    string
}

type Data struct {
	Channel     string `json:"channel"`
	Auth        string `json:"auth,ommitempty"`
	ChannelData string `json:"channel_data,ommitempty"`
}
