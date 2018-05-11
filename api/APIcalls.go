package api

import (
	"log"
	"net/url"
	"github.com/bigbluebutton-api-go/dataStructs"
	"github.com/bigbluebutton-api-go/helpers"

	"strconv"
)

var BASE_URL = "http://10.118.45.220//bigbluebutton/api/"
var SALT = "836fcbf304c917f91175c3b34f8c3347"

func GetJoinURL(participants *(dataStructs.Participants)) string {
	if "" == participants.FullName_ || "" == participants.MeetingID_ ||
		"" == participants.Password_ {
		return "ERROR: PARAM ERROR."
	}

	fullName := "fullName=" + url.QueryEscape(participants.FullName_)
	meetingID := "&meetingID=" + url.QueryEscape(participants.MeetingID_)
	password := "&password=" + url.QueryEscape(participants.Password_)

	var createTime string
	var userID string
	var configToken string
	var avatarURL string
	var redirect string
	var clientURL string

	if "" != participants.CreateTime {
		createTime = "&createTime=" + url.QueryEscape(participants.CreateTime)
	}

	if "" != participants.UserID {
		userID = "&userID=" + url.QueryEscape(participants.UserID)
	}

	if "" != participants.ConfigToken {
		configToken = "&configToken=" + url.QueryEscape(participants.ConfigToken)
	}

	if "" != participants.AvatarURL {
		avatarURL = "&avatarURL=" + url.QueryEscape(participants.AvatarURL)
	}

	if "" != participants.ClientURL {
		redirect = "&redirect=true"
		clientURL = "&clientURL=" + url.QueryEscape(participants.ClientURL)
	}

	joinParam := fullName + meetingID + password + createTime + userID +
		configToken + avatarURL + redirect + clientURL

	checksum := helpers.GetChecksum("join" + joinParam + SALT)
	joinUrl := BASE_URL + "join?" + joinParam + "&checksum=" + checksum
	participants.JoinURL = joinUrl

	return joinUrl
}

//only returns true when someone has joined the meeting
func IsMeetingRunning(meetingroom *dataStructs.MeetingRoom) bool {
	checksum := helpers.GetChecksum("isMeetingRunning" + "meetingID=" + meetingroom.MeetingID_ + SALT)
	getURL := BASE_URL + "isMeetingRunning?" + "meetingID=" + meetingroom.MeetingID_ + "&checksum=" + checksum
//	log.Println("the url we are GETting to check meeting running is: ", getURL)
	response := helpers.HttpGet(getURL)
	//log.Println(response)
	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")
		return false
	}
	var XMLResp dataStructs.IsMeetingRunningResponse
	//log.Println("**************** TYPE OF RESPONSE  " , reflect.TypeOf(XMLResp))
	// err := xml.Unmarshal([]byte(response),
	// 	&XMLResp)
	err := helpers.ReadXML(response, &XMLResp)
	if nil != err {
		return false
	}
	//log.Println("*** ", XMLResp, " ***")

	return XMLResp.Running
}

func EndMeeting(meetingRoom *dataStructs.MeetingRoom) string {
	log.Println("*** ending meeting ***")
	if meetingRoom.MeetingID_ == "" || meetingRoom.ModeratorPW_ == "" {
		log.Println("Error: PARAM ERROR.")
		return "Error: PARAM ERROR."
	}

	meetingID := "meetingID=" + url.QueryEscape(meetingRoom.MeetingID_)
	modPW := "&password=" + url.QueryEscape(meetingRoom.ModeratorPW_)
	param := meetingID + modPW
	checksum := helpers.GetChecksum("end" + param + SALT)

	getURL := BASE_URL + "end?" + param + "&checksum=" + checksum

	response := helpers.HttpGet(getURL)

	//	log.Println(response)
	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")
		return "Could not end meeting " + meetingRoom.MeetingID_
	}
	var XMLResp dataStructs.EndResponse

	err := helpers.ReadXML(response, &XMLResp)
	if nil != err {
		return "Could not end meeting " + meetingRoom.MeetingID_
	}
	//log.Println("*** ", XMLResp, " ***")

	if "SUCCESS" == XMLResp.ReturnCode {

		return "Successfully ended meeting " + meetingRoom.MeetingID_
	} else {
		return "Could not end meeting " + meetingRoom.MeetingID_
	}

}

func GetMeetingInfo(meetingRoom *dataStructs.MeetingRoom) string {
	log.Println("*** Getting meeting info ***")
	if meetingRoom.MeetingID_ == "" || meetingRoom.ModeratorPW_ == "" {
		log.Println("Error: PARAM ERROR.")
		return "Error: PARAM ERROR."
	}
	meetingID := "meetingID=" + url.QueryEscape(meetingRoom.MeetingID_)
	modPW := "&password=" + url.QueryEscape(meetingRoom.ModeratorPW_)
	param := meetingID + modPW
	checksum := helpers.GetChecksum("getMeetingInfo" + param + SALT)

	getURL := BASE_URL + "getMeetingInfo?" + param + "&checksum=" + checksum
	//log.Println(getURL)
	response := helpers.HttpGet(getURL)

	//	log.Println(response)
	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")
		return "Could not get meeting info " + meetingRoom.MeetingID_
	}

	err := helpers.ReadXML(response, &meetingRoom.MeetingInfo)
	if nil != err {
		return "Could not get meeting info " + meetingRoom.MeetingID_
	}
	//log.Println("*** ", XMLResp, " ***")

	if "SUCCESS" == meetingRoom.MeetingInfo.ReturnCode {
		println("Successfully got meeting info")
		return "Successfully got meeting info" + meetingRoom.MeetingID_
	} else {
		println("Could not get meeting info ")
		return "Could not get meeting info " + meetingRoom.MeetingID_
	}

}
func GetMeetings() dataStructs.GetMeetingsResponse{
	checksum := helpers.GetChecksum("getMeetings" + SALT)

	getURL := BASE_URL + "getMeetings?"  + "&checksum=" + checksum
	//log.Println(getURL)
	response := helpers.HttpGet(getURL)

	//	log.Println(response)
	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")

	}
	var XMLResp dataStructs.GetMeetingsResponse

	err := helpers.ReadXML(response, &XMLResp)
	if nil != err {

	}
	//log.Println("*** ", XMLResp, " ***")

	if "SUCCESS" == XMLResp.ReturnCode {
		println("Successfully got meetings info")

	} else {
		println("Could not get meetings info ")
	}
	return XMLResp

}
func GetRecordings() dataStructs.GetRecordingsResponse{
	checksum := helpers.GetChecksum("getRecordings" + SALT)

	getURL := BASE_URL + "getRecordings?"  + "&checksum=" + checksum
	response := helpers.HttpGet(getURL)

	//	log.Println(response)
	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")
	}
	var XMLResp dataStructs.GetRecordingsResponse

	err := helpers.ReadXML(response, &XMLResp)
	if nil != err {

	}
	//log.Println("*** ", XMLResp, " ***")

	if "SUCCESS" == XMLResp.ReturnCode {
		println("Successfully got recordings info")

	} else {
		println("Could not get recordings info ")
	}
	return XMLResp
}

func CreateMeeting(meetingRoom *dataStructs.MeetingRoom) string {
	if meetingRoom.Name_ == "" || meetingRoom.MeetingID_ == "" ||
		meetingRoom.AttendeePW_ == "" || meetingRoom.ModeratorPW_ == "" {
		log.Println("ERROR: PARAM ERROR.")
		return "ERROR: PARAM ERROR."
	}

	name := "name=" + url.QueryEscape(meetingRoom.Name_)
	meetingID := "&meetingID=" + url.QueryEscape(meetingRoom.MeetingID_)
	attendeePW := "&attendeePW=" + url.QueryEscape(meetingRoom.AttendeePW_)
	moderatorPW := "&moderatorPW=" + url.QueryEscape(meetingRoom.ModeratorPW_)

	var welcome string
	var logoutURL string
	var record string
	var duration string
	var moderatorOnlyMessage string
	var allowStartStopRecording string
	var voiceBridge string

	welcome = "&welcome=" + url.QueryEscape(meetingRoom.Welcome)

	logoutURL = "&logoutURL=" + url.QueryEscape(meetingRoom.LogoutURL)

	record = "&record=" + url.QueryEscape(meetingRoom.Record)

	duration = "&duration=" + url.QueryEscape(strconv.Itoa(meetingRoom.Duration))

	allowStartStopRecording = "&allowStartStopRecording=" +
		url.QueryEscape(strconv.FormatBool(meetingRoom.AllowStartStopRecording))

	moderatorOnlyMessage = "&moderatorOnlyMessage=" +
		url.QueryEscape(meetingRoom.ModeratorOnlyMessage)

	voiceBridge = "&voiceBridge=" + url.QueryEscape(meetingRoom.VoiceBridge)

	createParam := name + meetingID + attendeePW + moderatorPW + welcome +
		voiceBridge + logoutURL + record + duration + moderatorOnlyMessage +
		allowStartStopRecording

	checksum := helpers.GetChecksum("create" + createParam + SALT)

	response := helpers.HttpGet(BASE_URL + "create?" + createParam + "&checksum=" +
		checksum)

	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")
		return "ERROR: HTTP ERROR."
	}
	err := helpers.ReadXML(response, &meetingRoom.CreateMeetingResponse)

	if nil != err {
		log.Println("XML PARSE ERROR: " + err.Error())
		return "ERROR: XML PARSE ERROR."
	}

	if "SUCCESS" == meetingRoom.CreateMeetingResponse.Returncode {
		log.Println("SUCCESS CREATE MEETINGROOM. MEETING ID: " +
			meetingRoom.CreateMeetingResponse.MeetingID)
		return meetingRoom.CreateMeetingResponse.MeetingID
	} else {
		log.Println("CREATE MEETINGROOM FAILD: " + response)
		return "FAILED"
	}

	return "ERROR: UNKNOWN."
}
