package event

// ChannelEventLog is the CEL event structure
type ChannelEventLog struct {
	EventName     string `AMI:"EventName"`
	AccountCode   string `AMI:"AccountCode"`
	CallerIDnum   string `AMI:"CallerIDnum"`
	CallerIDname  string `AMI:"CallerIDname"`
	CallerIDani   string `AMI:"CallerIDani"`
	CallerIDrdnis string `AMI:"CallerIDrdnis"`
	CallerIDdnid  string `AMI:"CallerIDdnid"`
	Exten         string `AMI:"Exten"`
	Context       string `AMI:"Context"`
	Application   string `AMI:"Application"`
	AppData       string `AMI:"AppData"`
	EventTime     string `AMI:"EventTime"`
	AMAFlags      string `AMI:"AMAFlags"`
	UniqueID      string `AMI:"UniqueID"`
	LinkedID      string `AMI:"LinkedID"`
	UserField     string `AMI:"UserField"`
	Peer          string `AMI:"Peer"`
	PeerAccount   string `AMI:"PeerAccount"`
	Extra         string `AMI:"Extra"`
}

func init() {
	eventTrap["CEL"] = ChannelEventLog{}
}
