package main

import (
	"fmt"
	"log"

	"github.com/bigbluebutton-api-go/api"
	"github.com/bigbluebutton-api-go/dataStructs"
	"github.com/bigbluebutton-api-go/webhook"

)

func main() {
	//creates an empty
	var meetingRoom = dataStructs.MeetingRoom{}
	meetingRoom.MeetingID_ = "random-6807874"
	meetingRoom.Name_ = "random-6807874"
	meetingRoom.AttendeePW_ = "ap"
	meetingRoom.ModeratorPW_ = "mp"
	meetingRoom.Record = "true"
	meetingRoom.AllowStartStopRecording = true
	meetingRoom.AutoStartRecording = true

	var participant = dataStructs.Participants{}
	participant.FullName_ = "a name"
	participant.MeetingID_ = "random-6807874"
	participant.Password_ = "mp"

	api.CreateMeeting(&meetingRoom)
	fmt.Println()
	fmt.Println(api.GetJoinURL(&participant))
	fmt.Println()
	if api.IsMeetingRunning(meetingRoom.MeetingID_) {
		log.Println("meeting is running")
	} else {
		log.Println("meeting is not running")
	}
	fmt.Println()
	// fmt.Println(api.EndMeeting(&meetingRoom))
	// fmt.Println()
	//
	// api.CreateMeeting(&meetingRoom)
	// fmt.Println()
	//time.Sleep(30 * time.Second)
	api.GetMeetingInfo(meetingRoom.MeetingID_, meetingRoom.ModeratorPW_, &meetingRoom.MeetingInfo)
	fmt.Println()

	var wh = dataStructs.WebHook{}
	wh.CallBackURL = "http://postcatcher.in/catchers/5aff22a294c447040000000d"
	wh.MeetingId =   meetingRoom.MeetingID_

	fmt.Println("Creating webhook")
	fmt.Println(webhook.CreateHook(&wh))

	fmt.Println(webhook.DestroyHook("1"))
	fmt.Println(webhook.DestroyHook("2"))
	fmt.Println(webhook.DestroyHook("3"))

	var temp = api.GetRecordings()
	fmt.Println(temp)
	//fmt.Println(meetingRoom.MeetingInfo.Attendees)
	//allmeetings := api.GetMeetings()
	//fmt.Println(allmeetings.Meetings.MeetingInfo[0])// <-- to look at each meeting use this
}
