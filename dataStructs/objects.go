package dataStructs

//the following structs are the "things" we can create.
// ie participants, meetingRooms, recordings


type Recording struct {
  MeetingID     string
  RecordID      string
  State         string
  Meta          string
  Publish       string
}

type Participants struct {
	IsAdmin_     int
	FullName_    string
	MeetingID_   string
	Password_    string
	CreateTime   string
	UserID       string
	WebVoiceConf string
	ConfigToken  string
	AvatarURL    string
	Redirect     string
	ClientURL    string
	JoinURL      string
}

type MeetingRoom struct {
	Name_                   string
	MeetingID_              string
	AttendeePW_             string
	ModeratorPW_            string
	Welcome                 string
	DialNumber              string
	VoiceBridge             string
	WebVoice                string
	LogoutURL               string
	Record                  string
	Duration                int
	Meta                    string
	ModeratorOnlyMessage    string
	AutoStartRecording      bool
	AllowStartStopRecording bool
  Created                 bool

	CreateMeetingResponse CreateMeetingResponse
	MeetingInfo           GetMeetingInfoResponse

}

type WebHook struct {
  HookID      string
  CallBackURL string
  MeetingId   string

  WebhookResponse CreateWebhookResponse
}
