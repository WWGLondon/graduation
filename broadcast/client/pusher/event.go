package pusher

const (
	SubscribeEvent = "pusher:subscribe"
)

type Message struct {
	Event string `json:"event"`
	Data  Data   `json:"data"`
}

type Data struct {
	Channel     string `json:"channel"`
	Auth        string `json:"auth,ommitempty"`
	ChannelData string `json:"channel_data,ommitempty"`
}
