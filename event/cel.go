package event

// ChannelEventLog is the CEL event structure
type ChannelEventLog struct {
	EventName     string `AMI:"Eventname"`
	AccountCode   string `AMI:"Accountcode"`
	CallerIDnum   string `AMI:"Calleridnum"`
	CallerIDname  string `AMI:"Calleridname"`
	CallerIDani   string `AMI:"Calleridani"`
	CallerIDrdnis string `AMI:"Calleridrdnis"`
	CallerIDdnid  string `AMI:"Calleriddnid"`
	Exten         string `AMI:"Exten"`
	Context       string `AMI:"Context"`
	Application   string `AMI:"Application"`
	AppData       string `AMI:"Appdata"`
	EventTime     string `AMI:"Eventtime"`
	AMAFlags      int64  `AMI:"Amaflags"`
	UniqueID      string `AMI:"Uniqueid"`
	LinkedID      string `AMI:"Linkedid"`
	UserField     string `AMI:"Userfield"`
	Peer          string `AMI:"Peer"`
	PeerAccount   string `AMI:"Peeraccount"`
	Extra         string `AMI:"Extra"`
}

func init() {
	eventTrap["CEL"] = ChannelEventLog{}
}
