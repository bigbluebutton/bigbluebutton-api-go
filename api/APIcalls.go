package api

import (
	"log"
	"net/url"
	"github.com/bigbluebutton-api-go/dataStructs"
	"github.com/bigbluebutton-api-go/helpers"

	"strconv"
)

// var BASE_URL = "http://10.118.45.220/bigbluebutton/api/"
// var SALT = "836fcbf304c917f91175c3b34f8c3347"
var BASE_URL 	string
var SALT 			string

func SetAPI(url string, salt string){
	BASE_URL =  url
	SALT = salt
}

//TODO: needs things
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

	// var welcome string
	// var logoutURL string
	// var record string
	// var duration string
	// var moderatorOnlyMessage string
	// var allowStartStopRecording string
	// var voiceBridge string

	welcome := "&welcome=" + url.QueryEscape(meetingRoom.Welcome)

	dialNumber := "&dialNumber=" + url.QueryEscape(meetingRoom.DialNumber)

	logoutURL := "&logoutURL=" + url.QueryEscape(meetingRoom.LogoutURL)

	record := "&record=" + url.QueryEscape(meetingRoom.Record)

	duration := "&duration=" + url.QueryEscape(strconv.Itoa(meetingRoom.Duration))

	allowStartStopRecording := "&allowStartStopRecording=" +
		url.QueryEscape(strconv.FormatBool(meetingRoom.AllowStartStopRecording))

	moderatorOnlyMessage := "&moderatorOnlyMessage=" +
		url.QueryEscape(meetingRoom.ModeratorOnlyMessage)

	meta_bn_recording_ready_url := "&meta_bn-recording-ready-url=" +
		url.QueryEscape(meetingRoom.Meta_bn_recording_ready_url)

	meta_channelid := "&meta_channelid=" +
			url.QueryEscape(meetingRoom.Meta_channelid)

	meta_endcallback := 	"&meta_endcallbackurl=" +
		url.QueryEscape(meetingRoom.Meta_endcallbackurl)

	voiceBridge := "&voiceBridge=" + url.QueryEscape(meetingRoom.VoiceBridge)



	createParam := name + meetingID + attendeePW + moderatorPW + welcome + dialNumber +
		voiceBridge + logoutURL + record + duration + moderatorOnlyMessage + meta_bn_recording_ready_url +meta_channelid+
		meta_endcallback +allowStartStopRecording

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

// we need participants
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
	joinviahtml := "&joinViaHtml5=true"

	joinParam := fullName + meetingID + password + createTime + userID +
		configToken + avatarURL + redirect + clientURL +joinviahtml

	checksum := helpers.GetChecksum("join" + joinParam + SALT)
	joinUrl := BASE_URL + "join?" + joinParam + "&checksum=" + checksum
	participants.JoinURL = joinUrl

	return joinUrl
}

//only returns true when someone has joined the meeting
func IsMeetingRunning(meetingID string) bool {
	checksum := helpers.GetChecksum("isMeetingRunning" + "meetingID=" + meetingID + SALT)
	getURL := BASE_URL + "isMeetingRunning?" + "meetingID=" + meetingID+ "&checksum=" + checksum
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

func EndMeeting(meeting_ID string, mod_PW string) string {
	log.Println("*** ending meeting ***")


	meetingID := "meetingID=" + url.QueryEscape(meeting_ID)
	modPW := "&password=" + url.QueryEscape(mod_PW)
	param := meetingID + modPW
	checksum := helpers.GetChecksum("end" + param + SALT)

	getURL := BASE_URL + "end?" + param + "&checksum=" + checksum

	response := helpers.HttpGet(getURL)

	//	log.Println(response)
	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")
		return "Could not end meeting " + meeting_ID
	}
	var XMLResp dataStructs.EndResponse

	err := helpers.ReadXML(response, &XMLResp)
	if nil != err {
		return "Could not end meeting " + meeting_ID
	}
	//log.Println("*** ", XMLResp, " ***")

	if "SUCCESS" == XMLResp.ReturnCode {

		return "Successfully ended meeting " + meeting_ID
	} else {
		return "Could not end meeting " + meeting_ID
	}

}
//TODO: only needs meeting id and mod pw, and repsonse struct
func GetMeetingInfo(meeting_ID string, mod_PW string, responseXML *dataStructs.GetMeetingInfoResponse) string {
	log.Println("*** Getting meeting info ***")

	meetingID := "meetingID=" + url.QueryEscape(meeting_ID)
	modPW := "&password=" + url.QueryEscape(mod_PW)
	param := meetingID + modPW
	checksum := helpers.GetChecksum("getMeetingInfo" + param + SALT)

	getURL := BASE_URL + "getMeetingInfo?" + param + "&checksum=" + checksum
	//log.Println(getURL)
	response := helpers.HttpGet(getURL)

	//	log.Println(response)
	if "ERROR" == response {
		log.Println("ERROR: HTTP ERROR.")
		return "FAILED"
	}

	err := helpers.ReadXML(response, responseXML)
	if nil != err {
		return "FAILED"
	}
	//log.Println("*** ", XMLResp, " ***")

	if "SUCCESS" == responseXML.ReturnCode {
		println("Successfully got meeting info")
		return "Successfully got meeting info" + meeting_ID
	} else {
		println("Could not get meeting info ")
		return "FAILED"
	}

}

// TODO: doesn't need anything
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
func GetRecordings(meeting_id string, record_id string, metachannelid string) (dataStructs.GetRecordingsResponse, string){

	meetingID := "meetingID=" + url.QueryEscape(meeting_id)
	recordid := "&recordID=" + url.QueryEscape(record_id)
	var param string
	if (metachannelid != ""){
		meta_channelid := "meta_channelid=" +
				url.QueryEscape(metachannelid)
		param =  meta_channelid
	}else if (meeting_id != "" && record_id != ""){
		param = meetingID + recordid
	}else if (meeting_id != ""){
		param = meetingID
	}


	checksum := helpers.GetChecksum("getRecordings"  + param + SALT)

	getURL := BASE_URL + "getRecordings?"  +param + "&checksum=" + checksum
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
	return XMLResp,response
}

func PublishRecordings(recordid string, publish string) (dataStructs.PublishRecordingsResponse){
	recordID := "recordID=" + url.QueryEscape(recordid)
	Publish := "&publish=" + url.QueryEscape(publish)

	param := recordID + Publish
	checksum := helpers.GetChecksum("publishRecordings" + param + SALT)

	getURL := BASE_URL + "publishRecordings?" + param + "&checksum=" + checksum
	//log.Println(getURL)
	response := helpers.HttpGet(getURL)
	var XMLResp dataStructs.PublishRecordingsResponse

	helpers.ReadXML(response, &XMLResp)

	return XMLResp
}

func DeleteRecordings(recordid string) (dataStructs.DeleteRecordingsResponse){
	recordID := "recordID=" + url.QueryEscape(recordid)
	param := recordID
	checksum := helpers.GetChecksum("deleteRecordings" + param + SALT)

	getURL := BASE_URL + "deleteRecordings?" + param + "&checksum=" + checksum
	//log.Println(getURL)
	response := helpers.HttpGet(getURL)
	var XMLResp dataStructs.DeleteRecordingsResponse

	helpers.ReadXML(response, &XMLResp)

	return XMLResp
}
