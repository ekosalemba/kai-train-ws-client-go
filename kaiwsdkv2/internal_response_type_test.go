package kaiwsdkv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInternalGetOriginationRSOK is a positive test function for "InternalGetOriginationRS" internal response type <- mapping from "data.get_org"
func TestInternalGetOriginationRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0", "errMsg": null, "return": [{"originCode": "BD", "originName":"BANDUNG"},{"originCode": "YK", "originName":"YOGYAKARTA"}]}`

	// test expected values
	errCode := "0"
	orgLen := 2
	originCode := "BD"
	originName := "BANDUNG"

	// test variable
	var vRS InternalGetOriginationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, orgLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, originCode, vRS.Return[0].OriginCode, "should be equal")
	assert.Equal(t, originName, vRS.Return[0].OriginName, "should be equal")

	assert.Nil(t, err)
}

// TestInternalGetDestinationRSOK is a positive test function for "InternalGetDestinationRS" internal response type <- mapping from "data.get_des"
func TestInternalGetDestinationRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0", "errMsg": null, "return": [{"destCode": "BD", "destName":"BANDUNG"},{"destCode": "YK", "destName":"YOGYAKARTA"}]}`

	// test expected values
	errCode := "0"
	desLen := 2
	destCode := "YK"
	destName := "YOGYAKARTA"

	// test variable
	var vRS InternalGetDestinationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, desLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, destCode, vRS.Return[1].DestCode, "should be equal")
	assert.Equal(t, destName, vRS.Return[1].DestName, "should be equal")

	assert.Nil(t, err)
}

// TestInternalGetPayTypeRSOK is a positive test function for "InternalGetPayTypeRS" internal response type <- mapping from "data.get_pay_type"
func TestInternalGetPayTypeRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0", "errMsg": null, "return": [{"name": "TUNAI"},{"name": "EDC BNI"}]}`

	// test expected values
	errCode := "0"
	payLen := 2
	name := "TUNAI"

	// test variable
	var vRS InternalGetPayTypeRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, payLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, name, vRS.Return[0].Name, "should be equal")

	assert.Nil(t, err)
}

// TestInternalGetScheduleRSOK is a positive test function for "InternalGetScheduleRSOK" internal response type <- mapping from "information.get_schedule"
func TestInternalGetScheduleRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0","errMsg":null,"return":{"origin":"BD","destination":"GMR","departureDate":"20190919","schedules":[{"trainNo":"10501","trainName":"ARGO PARAHYANGAN PREMIUM","departureTime":"0415","arriveTime":"0725","availSubClass":[{"subClass":"C","seatAvailable":800,"seatClass":"K","adultPrice":100000,"childPrice":0,"infantPrice":0}]},{"trainNo":"710","trainName":"RANGKAS JAYA","departureTime":"0800","arriveTime":"1115","availSubClass":[{"subClass":"C","seatAvailable":424,"seatClass":"K","adultPrice":80000,"childPrice":0,"infantPrice":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	schedLen := 2
	trainNo0 := "10501"
	trainName0 := "ARGO PARAHYANGAN PREMIUM"

	// test variable
	var vRS InternalGetScheduleRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, schedLen, len(vRS.Return.Schedules), "should be equal!")
	assert.Equal(t, trainNo0, vRS.Return.Schedules[0].TrainNo, "should be equal!")
	assert.Equal(t, trainName0, vRS.Return.Schedules[0].TrainName, "should be equal!")

	assert.Nil(t, err)
}

// TestInternalGetScheduleV2RSOK is a positive test function for "InternalGetScheduleV2RS" internal response type <- mapping from "information.get_schedule_v2"
func TestInternalGetScheduleV2RSOK(t *testing.T) {
	// fake schema
	str := `{"errCode":"0","errMsg":null,"return":{"origin":"BD","destination":"GMR","departureDate":"20190919","schedules":[{"trainNo":"10501","trainName":"ARGO PARAHYANGAN PREMIUM","departureDate":"20190919","arriveDate":"20190919","departureTime":"0415","arriveTime":"0725","availSubClass":[{"subClass":"C","seatAvailable":800,"seatClass":"K","adultPrice":100000,"childPrice":0,"infantPrice":0}]},{"trainNo":"710","trainName":"RANGKAS JAYA","departureDate":"20190919","arriveDate":"20190919","departureTime":"0800","arriveTime":"1115","availSubClass":[{"subClass":"C","seatAvailable":424,"seatClass":"K","adultPrice":80000,"childPrice":0,"infantPrice":0}]},{"trainNo":"77A","trainName":"ARGO GOPAR","departureDate":"20190919","arriveDate":"20190919","departureTime":"1200","arriveTime":"1500","availSubClass":[{"subClass":"A","seatAvailable":49,"seatClass":"E","adultPrice":100000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":50,"seatClass":"B","adultPrice":90000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":49,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]},{"trainNo":"P05","trainName":"ARGO PARAHYANGAN","departureDate":"20190919","arriveDate":"20190919","departureTime":"2000","arriveTime":"2300","availSubClass":[{"subClass":"A","seatAvailable":0,"seatClass":"E","adultPrice":200000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":0,"seatClass":"B","adultPrice":150000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":240,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	schedLen := 4
	trainNo0 := "10501"
	trainName0 := "ARGO PARAHYANGAN PREMIUM"
	arvDate0 := "20190919"
	arvTime0 := "0725"

	// test variable
	var vRS InternalGetScheduleV2RS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, schedLen, len(vRS.Return.Schedules), "should be equal!")
	assert.Equal(t, trainNo0, vRS.Return.Schedules[0].TrainNo, "should be equal!")
	assert.Equal(t, trainName0, vRS.Return.Schedules[0].TrainName, "should be equal!")
	assert.Equal(t, arvDate0, vRS.Return.Schedules[0].ArriveDate, "should be equal!")
	assert.Equal(t, arvTime0, vRS.Return.Schedules[0].ArriveTime, "should be equal!")

	assert.Nil(t, err)
}

// TestInternalGetScheduleLiteRSOK is a positive test function for "InternalGetScheduleLiteRS" internal response type <- mapping from "information.get_schedule_lite"
func TestInternalGetScheduleLiteRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode":"0","errMsg":null,"return":{"origin":"BD","destination":"GMR","departureDate":"20190919","schedules":[{"trainNo":"10501","trainName":"ARGO PARAHYANGAN PREMIUM","departureDate":"20190919","arriveDate":"20190919","departureTime":"0415","arriveTime":"0725","availSubClass":[{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":100000,"childPrice":0,"infantPrice":0}]},{"trainNo":"710","trainName":"RANGKAS JAYA","departureDate":"20190919","arriveDate":"20190919","departureTime":"0800","arriveTime":"1115","availSubClass":[{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":80000,"childPrice":0,"infantPrice":0}]},{"trainNo":"77A","trainName":"ARGO GOPAR","departureDate":"20190919","arriveDate":"20190919","departureTime":"1200","arriveTime":"1500","availSubClass":[{"subClass":"A","seatAvailable":0,"seatClass":"E","adultPrice":100000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":0,"seatClass":"B","adultPrice":90000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]},{"trainNo":"P05","trainName":"ARGO PARAHYANGAN","departureDate":"20190919","arriveDate":"20190919","departureTime":"2000","arriveTime":"2300","availSubClass":[{"subClass":"A","seatAvailable":0,"seatClass":"E","adultPrice":200000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":0,"seatClass":"B","adultPrice":150000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	schedLen := 4
	trainNo0 := "10501"
	trainName0 := "ARGO PARAHYANGAN PREMIUM"
	arvDate0 := "20190919"
	arvTime0 := "0725"
	var seatAvailable0_0 float64 // = 0

	// test variable
	var vRS InternalGetScheduleLiteRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, schedLen, len(vRS.Return.Schedules), "should be equal!")
	assert.Equal(t, trainNo0, vRS.Return.Schedules[0].TrainNo, "should be equal!")
	assert.Equal(t, trainName0, vRS.Return.Schedules[0].TrainName, "should be equal!")
	assert.Equal(t, arvDate0, vRS.Return.Schedules[0].ArriveDate, "should be equal!")
	assert.Equal(t, arvTime0, vRS.Return.Schedules[0].ArriveTime, "should be equal!")
	assert.Equal(t, seatAvailable0_0, vRS.Return.Schedules[0].AvailSubClass[0].SeatAvailable, "should be equal!")

	assert.Nil(t, err)
}

// TestInternalGetSeatMapRSOK is a positive test function for "InternalGetSeatMapRS" internal response type <- mapping from "information.get_seat_map"
func TestInternalGetSeatMapRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode":"0","errMsg":null,"return":{"origin":"BD","destination":"GMR","trainNo":"10501","departureDate":"20190919","seatMaps":[{"wagonCode":"PREMIUM","wagonNo":1,"seats":[{"row":1,"column":1,"seatRow":1,"seatColumn":"A","seatClass":"C","status":0},{"row":1,"column":2,"seatRow":1,"seatColumn":"B","seatClass":"C","status":0},{"row":1,"column":3,"seatRow":1,"seatColumn":"","seatClass":"","status":0},{"row":1,"column":4,"seatRow":1,"seatColumn":"C","seatClass":"C","status":0},{"row":1,"column":5,"seatRow":1,"seatColumn":"D","seatClass":"C","status":0},{"row":1,"column":6,"seatRow":1,"seatColumn":"E","seatClass":"C","status":0},{"row":2,"column":1,"seatRow":2,"seatColumn":"A","seatClass":"C","status":0},{"row":2,"column":2,"seatRow":2,"seatColumn":"B","seatClass":"C","status":0},{"row":2,"column":3,"seatRow":2,"seatColumn":"","seatClass":"","status":0},{"row":2,"column":4,"seatRow":2,"seatColumn":"C","seatClass":"C","status":0},{"row":2,"column":5,"seatRow":2,"seatColumn":"D","seatClass":"C","status":0},{"row":2,"column":6,"seatRow":2,"seatColumn":"E","seatClass":"C","status":0},{"row":3,"column":1,"seatRow":3,"seatColumn":"A","seatClass":"C","status":0},{"row":3,"column":2,"seatRow":3,"seatColumn":"B","seatClass":"C","status":0},{"row":3,"column":3,"seatRow":3,"seatColumn":"","seatClass":"","status":0},{"row":3,"column":4,"seatRow":3,"seatColumn":"C","seatClass":"C","status":0},{"row":3,"column":5,"seatRow":3,"seatColumn":"D","seatClass":"C","status":0},{"row":3,"column":6,"seatRow":3,"seatColumn":"E","seatClass":"C","status":0},{"row":4,"column":1,"seatRow":4,"seatColumn":"A","seatClass":"C","status":0},{"row":4,"column":2,"seatRow":4,"seatColumn":"B","seatClass":"C","status":0},{"row":4,"column":3,"seatRow":4,"seatColumn":"","seatClass":"","status":0},{"row":4,"column":4,"seatRow":4,"seatColumn":"C","seatClass":"C","status":0},{"row":4,"column":5,"seatRow":4,"seatColumn":"D","seatClass":"C","status":0},{"row":4,"column":6,"seatRow":4,"seatColumn":"E","seatClass":"C","status":0},{"row":5,"column":1,"seatRow":5,"seatColumn":"A","seatClass":"C","status":0},{"row":5,"column":2,"seatRow":5,"seatColumn":"B","seatClass":"C","status":0},{"row":5,"column":3,"seatRow":5,"seatColumn":"","seatClass":"","status":0},{"row":5,"column":4,"seatRow":5,"seatColumn":"C","seatClass":"C","status":0},{"row":5,"column":5,"seatRow":5,"seatColumn":"D","seatClass":"C","status":0},{"row":5,"column":6,"seatRow":5,"seatColumn":"E","seatClass":"C","status":0},{"row":6,"column":1,"seatRow":6,"seatColumn":"A","seatClass":"C","status":0},{"row":6,"column":2,"seatRow":6,"seatColumn":"B","seatClass":"C","status":0},{"row":6,"column":3,"seatRow":6,"seatColumn":"","seatClass":"","status":0},{"row":6,"column":4,"seatRow":6,"seatColumn":"C","seatClass":"C","status":0},{"row":6,"column":5,"seatRow":6,"seatColumn":"D","seatClass":"C","status":0},{"row":6,"column":6,"seatRow":6,"seatColumn":"E","seatClass":"C","status":0},{"row":7,"column":1,"seatRow":7,"seatColumn":"A","seatClass":"C","status":0},{"row":7,"column":2,"seatRow":7,"seatColumn":"B","seatClass":"C","status":0},{"row":7,"column":3,"seatRow":7,"seatColumn":"","seatClass":"","status":0},{"row":7,"column":4,"seatRow":7,"seatColumn":"C","seatClass":"C","status":0},{"row":7,"column":5,"seatRow":7,"seatColumn":"D","seatClass":"C","status":0},{"row":7,"column":6,"seatRow":7,"seatColumn":"E","seatClass":"C","status":0},{"row":8,"column":1,"seatRow":8,"seatColumn":"A","seatClass":"C","status":0},{"row":8,"column":2,"seatRow":8,"seatColumn":"B","seatClass":"C","status":0},{"row":8,"column":3,"seatRow":8,"seatColumn":"","seatClass":"","status":0},{"row":8,"column":4,"seatRow":8,"seatColumn":"C","seatClass":"C","status":0},{"row":8,"column":5,"seatRow":8,"seatColumn":"D","seatClass":"C","status":0},{"row":8,"column":6,"seatRow":8,"seatColumn":"E","seatClass":"C","status":0},{"row":9,"column":1,"seatRow":9,"seatColumn":"A","seatClass":"C","status":0},{"row":9,"column":2,"seatRow":9,"seatColumn":"B","seatClass":"C","status":0},{"row":9,"column":3,"seatRow":9,"seatColumn":"","seatClass":"","status":0},{"row":9,"column":4,"seatRow":9,"seatColumn":"C","seatClass":"C","status":0},{"row":9,"column":5,"seatRow":9,"seatColumn":"D","seatClass":"C","status":0},{"row":9,"column":6,"seatRow":9,"seatColumn":"E","seatClass":"C","status":0},{"row":10,"column":1,"seatRow":10,"seatColumn":"A","seatClass":"C","status":0},{"row":10,"column":2,"seatRow":10,"seatColumn":"B","seatClass":"C","status":0},{"row":10,"column":3,"seatRow":10,"seatColumn":"","seatClass":"","status":0},{"row":10,"column":4,"seatRow":10,"seatColumn":"C","seatClass":"C","status":0},{"row":10,"column":5,"seatRow":10,"seatColumn":"D","seatClass":"C","status":0},{"row":10,"column":6,"seatRow":10,"seatColumn":"E","seatClass":"C","status":0},{"row":11,"column":1,"seatRow":11,"seatColumn":"A","seatClass":"C","status":0},{"row":11,"column":2,"seatRow":11,"seatColumn":"B","seatClass":"C","status":0},{"row":11,"column":3,"seatRow":11,"seatColumn":"","seatClass":"","status":0},{"row":11,"column":4,"seatRow":11,"seatColumn":"C","seatClass":"C","status":0},{"row":11,"column":5,"seatRow":11,"seatColumn":"D","seatClass":"C","status":0},{"row":11,"column":6,"seatRow":11,"seatColumn":"E","seatClass":"C","status":0},{"row":12,"column":1,"seatRow":12,"seatColumn":"A","seatClass":"C","status":0},{"row":12,"column":2,"seatRow":12,"seatColumn":"B","seatClass":"C","status":0},{"row":12,"column":3,"seatRow":12,"seatColumn":"","seatClass":"","status":0},{"row":12,"column":4,"seatRow":12,"seatColumn":"C","seatClass":"C","status":0},{"row":12,"column":5,"seatRow":12,"seatColumn":"D","seatClass":"C","status":0},{"row":12,"column":6,"seatRow":12,"seatColumn":"E","seatClass":"C","status":0},{"row":13,"column":1,"seatRow":13,"seatColumn":"A","seatClass":"C","status":0},{"row":13,"column":2,"seatRow":13,"seatColumn":"B","seatClass":"C","status":0},{"row":13,"column":3,"seatRow":13,"seatColumn":"","seatClass":"","status":0},{"row":13,"column":4,"seatRow":13,"seatColumn":"C","seatClass":"C","status":0},{"row":13,"column":5,"seatRow":13,"seatColumn":"D","seatClass":"C","status":0},{"row":13,"column":6,"seatRow":13,"seatColumn":"E","seatClass":"C","status":0},{"row":14,"column":1,"seatRow":14,"seatColumn":"A","seatClass":"C","status":0},{"row":14,"column":2,"seatRow":14,"seatColumn":"B","seatClass":"C","status":0},{"row":14,"column":3,"seatRow":14,"seatColumn":"","seatClass":"","status":0},{"row":14,"column":4,"seatRow":14,"seatColumn":"C","seatClass":"C","status":0},{"row":14,"column":5,"seatRow":14,"seatColumn":"D","seatClass":"C","status":0},{"row":14,"column":6,"seatRow":14,"seatColumn":"E","seatClass":"C","status":0},{"row":15,"column":1,"seatRow":15,"seatColumn":"A","seatClass":"C","status":0},{"row":15,"column":2,"seatRow":15,"seatColumn":"B","seatClass":"C","status":0},{"row":15,"column":3,"seatRow":15,"seatColumn":"","seatClass":"","status":0},{"row":15,"column":4,"seatRow":15,"seatColumn":"C","seatClass":"C","status":0},{"row":15,"column":5,"seatRow":15,"seatColumn":"D","seatClass":"C","status":0},{"row":15,"column":6,"seatRow":15,"seatColumn":"E","seatClass":"C","status":0},{"row":16,"column":1,"seatRow":16,"seatColumn":"A","seatClass":"C","status":0},{"row":16,"column":2,"seatRow":16,"seatColumn":"B","seatClass":"C","status":0},{"row":16,"column":3,"seatRow":16,"seatColumn":"","seatClass":"","status":0},{"row":16,"column":4,"seatRow":16,"seatColumn":"C","seatClass":"C","status":0},{"row":16,"column":5,"seatRow":16,"seatColumn":"D","seatClass":"C","status":0},{"row":16,"column":6,"seatRow":16,"seatColumn":"E","seatClass":"C","status":0},{"row":17,"column":1,"seatRow":17,"seatColumn":"A","seatClass":"C","status":0},{"row":17,"column":2,"seatRow":17,"seatColumn":"B","seatClass":"C","status":0},{"row":17,"column":3,"seatRow":17,"seatColumn":"","seatClass":"","status":0},{"row":17,"column":4,"seatRow":17,"seatColumn":"C","seatClass":"C","status":0},{"row":17,"column":5,"seatRow":17,"seatColumn":"D","seatClass":"C","status":0},{"row":17,"column":6,"seatRow":17,"seatColumn":"E","seatClass":"C","status":0}]},{"wagonCode":"PREMIUM","wagonNo":2,"seats":[{"row":1,"column":1,"seatRow":1,"seatColumn":"A","seatClass":"C","status":0},{"row":1,"column":2,"seatRow":1,"seatColumn":"B","seatClass":"C","status":0},{"row":1,"column":3,"seatRow":1,"seatColumn":"","seatClass":"","status":0},{"row":1,"column":4,"seatRow":1,"seatColumn":"C","seatClass":"C","status":0},{"row":1,"column":5,"seatRow":1,"seatColumn":"D","seatClass":"C","status":0},{"row":1,"column":6,"seatRow":1,"seatColumn":"E","seatClass":"C","status":0},{"row":2,"column":1,"seatRow":2,"seatColumn":"A","seatClass":"C","status":0},{"row":2,"column":2,"seatRow":2,"seatColumn":"B","seatClass":"C","status":0},{"row":2,"column":3,"seatRow":2,"seatColumn":"","seatClass":"","status":0},{"row":2,"column":4,"seatRow":2,"seatColumn":"C","seatClass":"C","status":0},{"row":2,"column":5,"seatRow":2,"seatColumn":"D","seatClass":"C","status":0},{"row":2,"column":6,"seatRow":2,"seatColumn":"E","seatClass":"C","status":0},{"row":3,"column":1,"seatRow":3,"seatColumn":"A","seatClass":"C","status":0},{"row":3,"column":2,"seatRow":3,"seatColumn":"B","seatClass":"C","status":0},{"row":3,"column":3,"seatRow":3,"seatColumn":"","seatClass":"","status":0},{"row":3,"column":4,"seatRow":3,"seatColumn":"C","seatClass":"C","status":0},{"row":3,"column":5,"seatRow":3,"seatColumn":"D","seatClass":"C","status":0},{"row":3,"column":6,"seatRow":3,"seatColumn":"E","seatClass":"C","status":0},{"row":4,"column":1,"seatRow":4,"seatColumn":"A","seatClass":"C","status":0},{"row":4,"column":2,"seatRow":4,"seatColumn":"B","seatClass":"C","status":0},{"row":4,"column":3,"seatRow":4,"seatColumn":"","seatClass":"","status":0},{"row":4,"column":4,"seatRow":4,"seatColumn":"C","seatClass":"C","status":0},{"row":4,"column":5,"seatRow":4,"seatColumn":"D","seatClass":"C","status":0},{"row":4,"column":6,"seatRow":4,"seatColumn":"E","seatClass":"C","status":0},{"row":5,"column":1,"seatRow":5,"seatColumn":"A","seatClass":"C","status":0},{"row":5,"column":2,"seatRow":5,"seatColumn":"B","seatClass":"C","status":0},{"row":5,"column":3,"seatRow":5,"seatColumn":"","seatClass":"","status":0},{"row":5,"column":4,"seatRow":5,"seatColumn":"C","seatClass":"C","status":0},{"row":5,"column":5,"seatRow":5,"seatColumn":"D","seatClass":"C","status":0},{"row":5,"column":6,"seatRow":5,"seatColumn":"E","seatClass":"C","status":0},{"row":6,"column":1,"seatRow":6,"seatColumn":"A","seatClass":"C","status":0},{"row":6,"column":2,"seatRow":6,"seatColumn":"B","seatClass":"C","status":0},{"row":6,"column":3,"seatRow":6,"seatColumn":"","seatClass":"","status":0},{"row":6,"column":4,"seatRow":6,"seatColumn":"C","seatClass":"C","status":0},{"row":6,"column":5,"seatRow":6,"seatColumn":"D","seatClass":"C","status":0},{"row":6,"column":6,"seatRow":6,"seatColumn":"E","seatClass":"C","status":0},{"row":7,"column":1,"seatRow":7,"seatColumn":"A","seatClass":"C","status":0},{"row":7,"column":2,"seatRow":7,"seatColumn":"B","seatClass":"C","status":0},{"row":7,"column":3,"seatRow":7,"seatColumn":"","seatClass":"","status":0},{"row":7,"column":4,"seatRow":7,"seatColumn":"C","seatClass":"C","status":0},{"row":7,"column":5,"seatRow":7,"seatColumn":"D","seatClass":"C","status":0},{"row":7,"column":6,"seatRow":7,"seatColumn":"E","seatClass":"C","status":0},{"row":8,"column":1,"seatRow":8,"seatColumn":"A","seatClass":"C","status":0},{"row":8,"column":2,"seatRow":8,"seatColumn":"B","seatClass":"C","status":0},{"row":8,"column":3,"seatRow":8,"seatColumn":"","seatClass":"","status":0},{"row":8,"column":4,"seatRow":8,"seatColumn":"C","seatClass":"C","status":0},{"row":8,"column":5,"seatRow":8,"seatColumn":"D","seatClass":"C","status":0},{"row":8,"column":6,"seatRow":8,"seatColumn":"E","seatClass":"C","status":0},{"row":9,"column":1,"seatRow":9,"seatColumn":"A","seatClass":"C","status":0},{"row":9,"column":2,"seatRow":9,"seatColumn":"B","seatClass":"C","status":0},{"row":9,"column":3,"seatRow":9,"seatColumn":"","seatClass":"","status":0},{"row":9,"column":4,"seatRow":9,"seatColumn":"C","seatClass":"C","status":0},{"row":9,"column":5,"seatRow":9,"seatColumn":"D","seatClass":"C","status":0},{"row":9,"column":6,"seatRow":9,"seatColumn":"E","seatClass":"C","status":0},{"row":10,"column":1,"seatRow":10,"seatColumn":"A","seatClass":"C","status":0},{"row":10,"column":2,"seatRow":10,"seatColumn":"B","seatClass":"C","status":0},{"row":10,"column":3,"seatRow":10,"seatColumn":"","seatClass":"","status":0},{"row":10,"column":4,"seatRow":10,"seatColumn":"C","seatClass":"C","status":0},{"row":10,"column":5,"seatRow":10,"seatColumn":"D","seatClass":"C","status":0},{"row":10,"column":6,"seatRow":10,"seatColumn":"E","seatClass":"C","status":0},{"row":11,"column":1,"seatRow":11,"seatColumn":"A","seatClass":"C","status":0},{"row":11,"column":2,"seatRow":11,"seatColumn":"B","seatClass":"C","status":0},{"row":11,"column":3,"seatRow":11,"seatColumn":"","seatClass":"","status":0},{"row":11,"column":4,"seatRow":11,"seatColumn":"C","seatClass":"C","status":0},{"row":11,"column":5,"seatRow":11,"seatColumn":"D","seatClass":"C","status":0},{"row":11,"column":6,"seatRow":11,"seatColumn":"E","seatClass":"C","status":0},{"row":12,"column":1,"seatRow":12,"seatColumn":"A","seatClass":"C","status":0},{"row":12,"column":2,"seatRow":12,"seatColumn":"B","seatClass":"C","status":0},{"row":12,"column":3,"seatRow":12,"seatColumn":"","seatClass":"","status":0},{"row":12,"column":4,"seatRow":12,"seatColumn":"C","seatClass":"C","status":0},{"row":12,"column":5,"seatRow":12,"seatColumn":"D","seatClass":"C","status":0},{"row":12,"column":6,"seatRow":12,"seatColumn":"E","seatClass":"C","status":0},{"row":13,"column":1,"seatRow":13,"seatColumn":"A","seatClass":"C","status":0},{"row":13,"column":2,"seatRow":13,"seatColumn":"B","seatClass":"C","status":0},{"row":13,"column":3,"seatRow":13,"seatColumn":"","seatClass":"","status":0},{"row":13,"column":4,"seatRow":13,"seatColumn":"C","seatClass":"C","status":0},{"row":13,"column":5,"seatRow":13,"seatColumn":"D","seatClass":"C","status":0},{"row":13,"column":6,"seatRow":13,"seatColumn":"E","seatClass":"C","status":0},{"row":14,"column":1,"seatRow":14,"seatColumn":"A","seatClass":"C","status":0},{"row":14,"column":2,"seatRow":14,"seatColumn":"B","seatClass":"C","status":0},{"row":14,"column":3,"seatRow":14,"seatColumn":"","seatClass":"","status":0},{"row":14,"column":4,"seatRow":14,"seatColumn":"C","seatClass":"C","status":0},{"row":14,"column":5,"seatRow":14,"seatColumn":"D","seatClass":"C","status":0},{"row":14,"column":6,"seatRow":14,"seatColumn":"E","seatClass":"C","status":0},{"row":15,"column":1,"seatRow":15,"seatColumn":"A","seatClass":"C","status":0},{"row":15,"column":2,"seatRow":15,"seatColumn":"B","seatClass":"C","status":0},{"row":15,"column":3,"seatRow":15,"seatColumn":"","seatClass":"","status":0},{"row":15,"column":4,"seatRow":15,"seatColumn":"C","seatClass":"C","status":0},{"row":15,"column":5,"seatRow":15,"seatColumn":"D","seatClass":"C","status":0},{"row":15,"column":6,"seatRow":15,"seatColumn":"E","seatClass":"C","status":0},{"row":16,"column":1,"seatRow":16,"seatColumn":"A","seatClass":"C","status":0},{"row":16,"column":2,"seatRow":16,"seatColumn":"B","seatClass":"C","status":0},{"row":16,"column":3,"seatRow":16,"seatColumn":"","seatClass":"","status":0},{"row":16,"column":4,"seatRow":16,"seatColumn":"C","seatClass":"C","status":0},{"row":16,"column":5,"seatRow":16,"seatColumn":"D","seatClass":"C","status":0},{"row":16,"column":6,"seatRow":16,"seatColumn":"E","seatClass":"C","status":0},{"row":17,"column":1,"seatRow":17,"seatColumn":"A","seatClass":"C","status":0},{"row":17,"column":2,"seatRow":17,"seatColumn":"B","seatClass":"C","status":0},{"row":17,"column":3,"seatRow":17,"seatColumn":"","seatClass":"","status":0},{"row":17,"column":4,"seatRow":17,"seatColumn":"C","seatClass":"C","status":0},{"row":17,"column":5,"seatRow":17,"seatColumn":"D","seatClass":"C","status":0},{"row":17,"column":6,"seatRow":17,"seatColumn":"E","seatClass":"C","status":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	seatMapLen := 2
	trainNo := "10501"
	wagonCode0 := "PREMIUM"
	var wagonNo0 float64 = 1
	var row0_0 float64 = 1

	// test variable
	var vRS InternalGetSeatMapRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, seatMapLen, len(vRS.Return.SeatMaps), "should be equal!")
	assert.Equal(t, trainNo, vRS.Return.TrainNo, "should be equal!")

	assert.Equal(t, wagonCode0, vRS.Return.SeatMaps[0].WagonCode, "should be equal!")
	assert.Equal(t, wagonNo0, vRS.Return.SeatMaps[0].WagonNo, "should be equal!")
	assert.Equal(t, row0_0, vRS.Return.SeatMaps[0].Seats[0].Row, "should be equal!")

	assert.Nil(t, err)
}

// TestInternalGetSeatMapPerSubClassRSOK is a positive test function for "InternalGetSeatMapPerSubClassRS" internal response type <- mapping from "information.get_seat_map_per_subclass"
func TestInternalGetSeatMapPerSubClassRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode":"0","errMsg":null,"return":{"origin":"BD","destination":"GMR","trainNo":"10501","departureDate":"20190919","seatMaps":[{"wagonCode":"PREMIUM","wagonNo":1,"seats":[{"row":1,"column":1,"seatRow":1,"seatColumn":"A","seatClass":"C","status":0},{"row":1,"column":2,"seatRow":1,"seatColumn":"B","seatClass":"C","status":0},{"row":1,"column":4,"seatRow":1,"seatColumn":"C","seatClass":"C","status":0},{"row":1,"column":5,"seatRow":1,"seatColumn":"D","seatClass":"C","status":0},{"row":1,"column":6,"seatRow":1,"seatColumn":"E","seatClass":"C","status":0},{"row":2,"column":1,"seatRow":2,"seatColumn":"A","seatClass":"C","status":0},{"row":2,"column":2,"seatRow":2,"seatColumn":"B","seatClass":"C","status":0},{"row":2,"column":4,"seatRow":2,"seatColumn":"C","seatClass":"C","status":0},{"row":2,"column":5,"seatRow":2,"seatColumn":"D","seatClass":"C","status":0},{"row":2,"column":6,"seatRow":2,"seatColumn":"E","seatClass":"C","status":0},{"row":3,"column":1,"seatRow":3,"seatColumn":"A","seatClass":"C","status":0},{"row":3,"column":2,"seatRow":3,"seatColumn":"B","seatClass":"C","status":0},{"row":3,"column":4,"seatRow":3,"seatColumn":"C","seatClass":"C","status":0},{"row":3,"column":5,"seatRow":3,"seatColumn":"D","seatClass":"C","status":0},{"row":3,"column":6,"seatRow":3,"seatColumn":"E","seatClass":"C","status":0},{"row":4,"column":1,"seatRow":4,"seatColumn":"A","seatClass":"C","status":0},{"row":4,"column":2,"seatRow":4,"seatColumn":"B","seatClass":"C","status":0},{"row":4,"column":4,"seatRow":4,"seatColumn":"C","seatClass":"C","status":0},{"row":4,"column":5,"seatRow":4,"seatColumn":"D","seatClass":"C","status":0},{"row":4,"column":6,"seatRow":4,"seatColumn":"E","seatClass":"C","status":0},{"row":5,"column":1,"seatRow":5,"seatColumn":"A","seatClass":"C","status":0},{"row":5,"column":2,"seatRow":5,"seatColumn":"B","seatClass":"C","status":0},{"row":5,"column":4,"seatRow":5,"seatColumn":"C","seatClass":"C","status":0},{"row":5,"column":5,"seatRow":5,"seatColumn":"D","seatClass":"C","status":0},{"row":5,"column":6,"seatRow":5,"seatColumn":"E","seatClass":"C","status":0},{"row":6,"column":1,"seatRow":6,"seatColumn":"A","seatClass":"C","status":0},{"row":6,"column":2,"seatRow":6,"seatColumn":"B","seatClass":"C","status":0},{"row":6,"column":4,"seatRow":6,"seatColumn":"C","seatClass":"C","status":0},{"row":6,"column":5,"seatRow":6,"seatColumn":"D","seatClass":"C","status":0},{"row":6,"column":6,"seatRow":6,"seatColumn":"E","seatClass":"C","status":0},{"row":7,"column":1,"seatRow":7,"seatColumn":"A","seatClass":"C","status":0},{"row":7,"column":2,"seatRow":7,"seatColumn":"B","seatClass":"C","status":0},{"row":7,"column":4,"seatRow":7,"seatColumn":"C","seatClass":"C","status":0},{"row":7,"column":5,"seatRow":7,"seatColumn":"D","seatClass":"C","status":0},{"row":7,"column":6,"seatRow":7,"seatColumn":"E","seatClass":"C","status":0},{"row":8,"column":1,"seatRow":8,"seatColumn":"A","seatClass":"C","status":0},{"row":8,"column":2,"seatRow":8,"seatColumn":"B","seatClass":"C","status":0},{"row":8,"column":4,"seatRow":8,"seatColumn":"C","seatClass":"C","status":0},{"row":8,"column":5,"seatRow":8,"seatColumn":"D","seatClass":"C","status":0},{"row":8,"column":6,"seatRow":8,"seatColumn":"E","seatClass":"C","status":0},{"row":9,"column":1,"seatRow":9,"seatColumn":"A","seatClass":"C","status":0},{"row":9,"column":2,"seatRow":9,"seatColumn":"B","seatClass":"C","status":0},{"row":9,"column":4,"seatRow":9,"seatColumn":"C","seatClass":"C","status":0},{"row":9,"column":5,"seatRow":9,"seatColumn":"D","seatClass":"C","status":0},{"row":9,"column":6,"seatRow":9,"seatColumn":"E","seatClass":"C","status":0},{"row":10,"column":1,"seatRow":10,"seatColumn":"A","seatClass":"C","status":0},{"row":10,"column":2,"seatRow":10,"seatColumn":"B","seatClass":"C","status":0},{"row":10,"column":4,"seatRow":10,"seatColumn":"C","seatClass":"C","status":0},{"row":10,"column":5,"seatRow":10,"seatColumn":"D","seatClass":"C","status":0},{"row":10,"column":6,"seatRow":10,"seatColumn":"E","seatClass":"C","status":0},{"row":11,"column":1,"seatRow":11,"seatColumn":"A","seatClass":"C","status":0},{"row":11,"column":2,"seatRow":11,"seatColumn":"B","seatClass":"C","status":0},{"row":11,"column":4,"seatRow":11,"seatColumn":"C","seatClass":"C","status":0},{"row":11,"column":5,"seatRow":11,"seatColumn":"D","seatClass":"C","status":0},{"row":11,"column":6,"seatRow":11,"seatColumn":"E","seatClass":"C","status":0},{"row":12,"column":1,"seatRow":12,"seatColumn":"A","seatClass":"C","status":0},{"row":12,"column":2,"seatRow":12,"seatColumn":"B","seatClass":"C","status":0},{"row":12,"column":4,"seatRow":12,"seatColumn":"C","seatClass":"C","status":0},{"row":12,"column":5,"seatRow":12,"seatColumn":"D","seatClass":"C","status":0},{"row":12,"column":6,"seatRow":12,"seatColumn":"E","seatClass":"C","status":0},{"row":13,"column":1,"seatRow":13,"seatColumn":"A","seatClass":"C","status":0},{"row":13,"column":2,"seatRow":13,"seatColumn":"B","seatClass":"C","status":0},{"row":13,"column":4,"seatRow":13,"seatColumn":"C","seatClass":"C","status":0},{"row":13,"column":5,"seatRow":13,"seatColumn":"D","seatClass":"C","status":0},{"row":13,"column":6,"seatRow":13,"seatColumn":"E","seatClass":"C","status":0},{"row":14,"column":1,"seatRow":14,"seatColumn":"A","seatClass":"C","status":0},{"row":14,"column":2,"seatRow":14,"seatColumn":"B","seatClass":"C","status":0},{"row":14,"column":4,"seatRow":14,"seatColumn":"C","seatClass":"C","status":0},{"row":14,"column":5,"seatRow":14,"seatColumn":"D","seatClass":"C","status":0},{"row":14,"column":6,"seatRow":14,"seatColumn":"E","seatClass":"C","status":0},{"row":15,"column":1,"seatRow":15,"seatColumn":"A","seatClass":"C","status":0},{"row":15,"column":2,"seatRow":15,"seatColumn":"B","seatClass":"C","status":0},{"row":15,"column":4,"seatRow":15,"seatColumn":"C","seatClass":"C","status":0},{"row":15,"column":5,"seatRow":15,"seatColumn":"D","seatClass":"C","status":0},{"row":15,"column":6,"seatRow":15,"seatColumn":"E","seatClass":"C","status":0},{"row":16,"column":1,"seatRow":16,"seatColumn":"A","seatClass":"C","status":0},{"row":16,"column":2,"seatRow":16,"seatColumn":"B","seatClass":"C","status":0},{"row":16,"column":4,"seatRow":16,"seatColumn":"C","seatClass":"C","status":0},{"row":16,"column":5,"seatRow":16,"seatColumn":"D","seatClass":"C","status":0},{"row":16,"column":6,"seatRow":16,"seatColumn":"E","seatClass":"C","status":0},{"row":17,"column":1,"seatRow":17,"seatColumn":"A","seatClass":"C","status":0},{"row":17,"column":2,"seatRow":17,"seatColumn":"B","seatClass":"C","status":0},{"row":17,"column":4,"seatRow":17,"seatColumn":"C","seatClass":"C","status":0},{"row":17,"column":5,"seatRow":17,"seatColumn":"D","seatClass":"C","status":0},{"row":17,"column":6,"seatRow":17,"seatColumn":"E","seatClass":"C","status":0}]},{"wagonCode":"PREMIUM","wagonNo":2,"seats":[{"row":1,"column":1,"seatRow":1,"seatColumn":"A","seatClass":"C","status":0},{"row":1,"column":2,"seatRow":1,"seatColumn":"B","seatClass":"C","status":0},{"row":1,"column":4,"seatRow":1,"seatColumn":"C","seatClass":"C","status":0},{"row":1,"column":5,"seatRow":1,"seatColumn":"D","seatClass":"C","status":0},{"row":1,"column":6,"seatRow":1,"seatColumn":"E","seatClass":"C","status":0},{"row":2,"column":1,"seatRow":2,"seatColumn":"A","seatClass":"C","status":0},{"row":2,"column":2,"seatRow":2,"seatColumn":"B","seatClass":"C","status":0},{"row":2,"column":4,"seatRow":2,"seatColumn":"C","seatClass":"C","status":0},{"row":2,"column":5,"seatRow":2,"seatColumn":"D","seatClass":"C","status":0},{"row":2,"column":6,"seatRow":2,"seatColumn":"E","seatClass":"C","status":0},{"row":3,"column":1,"seatRow":3,"seatColumn":"A","seatClass":"C","status":0},{"row":3,"column":2,"seatRow":3,"seatColumn":"B","seatClass":"C","status":0},{"row":3,"column":4,"seatRow":3,"seatColumn":"C","seatClass":"C","status":0},{"row":3,"column":5,"seatRow":3,"seatColumn":"D","seatClass":"C","status":0},{"row":3,"column":6,"seatRow":3,"seatColumn":"E","seatClass":"C","status":0},{"row":4,"column":1,"seatRow":4,"seatColumn":"A","seatClass":"C","status":0},{"row":4,"column":2,"seatRow":4,"seatColumn":"B","seatClass":"C","status":0},{"row":4,"column":4,"seatRow":4,"seatColumn":"C","seatClass":"C","status":0},{"row":4,"column":5,"seatRow":4,"seatColumn":"D","seatClass":"C","status":0},{"row":4,"column":6,"seatRow":4,"seatColumn":"E","seatClass":"C","status":0},{"row":5,"column":1,"seatRow":5,"seatColumn":"A","seatClass":"C","status":0},{"row":5,"column":2,"seatRow":5,"seatColumn":"B","seatClass":"C","status":0},{"row":5,"column":4,"seatRow":5,"seatColumn":"C","seatClass":"C","status":0},{"row":5,"column":5,"seatRow":5,"seatColumn":"D","seatClass":"C","status":0},{"row":5,"column":6,"seatRow":5,"seatColumn":"E","seatClass":"C","status":0},{"row":6,"column":1,"seatRow":6,"seatColumn":"A","seatClass":"C","status":0},{"row":6,"column":2,"seatRow":6,"seatColumn":"B","seatClass":"C","status":0},{"row":6,"column":4,"seatRow":6,"seatColumn":"C","seatClass":"C","status":0},{"row":6,"column":5,"seatRow":6,"seatColumn":"D","seatClass":"C","status":0},{"row":6,"column":6,"seatRow":6,"seatColumn":"E","seatClass":"C","status":0},{"row":7,"column":1,"seatRow":7,"seatColumn":"A","seatClass":"C","status":0},{"row":7,"column":2,"seatRow":7,"seatColumn":"B","seatClass":"C","status":0},{"row":7,"column":4,"seatRow":7,"seatColumn":"C","seatClass":"C","status":0},{"row":7,"column":5,"seatRow":7,"seatColumn":"D","seatClass":"C","status":0},{"row":7,"column":6,"seatRow":7,"seatColumn":"E","seatClass":"C","status":0},{"row":8,"column":1,"seatRow":8,"seatColumn":"A","seatClass":"C","status":0},{"row":8,"column":2,"seatRow":8,"seatColumn":"B","seatClass":"C","status":0},{"row":8,"column":4,"seatRow":8,"seatColumn":"C","seatClass":"C","status":0},{"row":8,"column":5,"seatRow":8,"seatColumn":"D","seatClass":"C","status":0},{"row":8,"column":6,"seatRow":8,"seatColumn":"E","seatClass":"C","status":0},{"row":9,"column":1,"seatRow":9,"seatColumn":"A","seatClass":"C","status":0},{"row":9,"column":2,"seatRow":9,"seatColumn":"B","seatClass":"C","status":0},{"row":9,"column":4,"seatRow":9,"seatColumn":"C","seatClass":"C","status":0},{"row":9,"column":5,"seatRow":9,"seatColumn":"D","seatClass":"C","status":0},{"row":9,"column":6,"seatRow":9,"seatColumn":"E","seatClass":"C","status":0},{"row":10,"column":1,"seatRow":10,"seatColumn":"A","seatClass":"C","status":0},{"row":10,"column":2,"seatRow":10,"seatColumn":"B","seatClass":"C","status":0},{"row":10,"column":4,"seatRow":10,"seatColumn":"C","seatClass":"C","status":0},{"row":10,"column":5,"seatRow":10,"seatColumn":"D","seatClass":"C","status":0},{"row":10,"column":6,"seatRow":10,"seatColumn":"E","seatClass":"C","status":0},{"row":11,"column":1,"seatRow":11,"seatColumn":"A","seatClass":"C","status":0},{"row":11,"column":2,"seatRow":11,"seatColumn":"B","seatClass":"C","status":0},{"row":11,"column":4,"seatRow":11,"seatColumn":"C","seatClass":"C","status":0},{"row":11,"column":5,"seatRow":11,"seatColumn":"D","seatClass":"C","status":0},{"row":11,"column":6,"seatRow":11,"seatColumn":"E","seatClass":"C","status":0},{"row":12,"column":1,"seatRow":12,"seatColumn":"A","seatClass":"C","status":0},{"row":12,"column":2,"seatRow":12,"seatColumn":"B","seatClass":"C","status":0},{"row":12,"column":4,"seatRow":12,"seatColumn":"C","seatClass":"C","status":0},{"row":12,"column":5,"seatRow":12,"seatColumn":"D","seatClass":"C","status":0},{"row":12,"column":6,"seatRow":12,"seatColumn":"E","seatClass":"C","status":0},{"row":13,"column":1,"seatRow":13,"seatColumn":"A","seatClass":"C","status":0},{"row":13,"column":2,"seatRow":13,"seatColumn":"B","seatClass":"C","status":0},{"row":13,"column":4,"seatRow":13,"seatColumn":"C","seatClass":"C","status":0},{"row":13,"column":5,"seatRow":13,"seatColumn":"D","seatClass":"C","status":0},{"row":13,"column":6,"seatRow":13,"seatColumn":"E","seatClass":"C","status":0},{"row":14,"column":1,"seatRow":14,"seatColumn":"A","seatClass":"C","status":0},{"row":14,"column":2,"seatRow":14,"seatColumn":"B","seatClass":"C","status":0},{"row":14,"column":4,"seatRow":14,"seatColumn":"C","seatClass":"C","status":0},{"row":14,"column":5,"seatRow":14,"seatColumn":"D","seatClass":"C","status":0},{"row":14,"column":6,"seatRow":14,"seatColumn":"E","seatClass":"C","status":0},{"row":15,"column":1,"seatRow":15,"seatColumn":"A","seatClass":"C","status":0},{"row":15,"column":2,"seatRow":15,"seatColumn":"B","seatClass":"C","status":0},{"row":15,"column":4,"seatRow":15,"seatColumn":"C","seatClass":"C","status":0},{"row":15,"column":5,"seatRow":15,"seatColumn":"D","seatClass":"C","status":0},{"row":15,"column":6,"seatRow":15,"seatColumn":"E","seatClass":"C","status":0},{"row":16,"column":1,"seatRow":16,"seatColumn":"A","seatClass":"C","status":0},{"row":16,"column":2,"seatRow":16,"seatColumn":"B","seatClass":"C","status":0},{"row":16,"column":4,"seatRow":16,"seatColumn":"C","seatClass":"C","status":0},{"row":16,"column":5,"seatRow":16,"seatColumn":"D","seatClass":"C","status":0},{"row":16,"column":6,"seatRow":16,"seatColumn":"E","seatClass":"C","status":0},{"row":17,"column":1,"seatRow":17,"seatColumn":"A","seatClass":"C","status":0},{"row":17,"column":2,"seatRow":17,"seatColumn":"B","seatClass":"C","status":0},{"row":17,"column":4,"seatRow":17,"seatColumn":"C","seatClass":"C","status":0},{"row":17,"column":5,"seatRow":17,"seatColumn":"D","seatClass":"C","status":0},{"row":17,"column":6,"seatRow":17,"seatColumn":"E","seatClass":"C","status":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	seatMapLen := 2
	trainNo := "10501"
	wagonCode0 := "PREMIUM"
	var wagonNo0 float64 = 1
	var row0_0 float64 = 1

	// test variable
	var vRS InternalGetSeatMapPerSubClassRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, seatMapLen, len(vRS.Return.SeatMaps), "should be equal!")
	assert.Equal(t, trainNo, vRS.Return.TrainNo, "should be equal!")

	assert.Equal(t, wagonCode0, vRS.Return.SeatMaps[0].WagonCode, "should be equal!")
	assert.Equal(t, wagonNo0, vRS.Return.SeatMaps[0].WagonNo, "should be equal!")
	assert.Equal(t, row0_0, vRS.Return.SeatMaps[0].Seats[0].Row, "should be equal!")

	assert.Equal(t, vRS.Return.SeatMaps[0].Seats[0].SubClass, vRS.Return.SeatMaps[1].Seats[1].SubClass, "should be equal!")

	assert.Nil(t, err)
}

// TestInternalGetSeatNullRSOK is a positive test function for "InternalGetSeatNullRS" internal response type <- mapping from "information.get_seat_null"
func TestInternalGetSeatNullRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode":"0","errMsg":null,"return":{"origin":"BD","destination":"GMR","trainNo":"10501","departureDate":"20190919","seatNulls":[{"wagonCode":"PREMIUM","wagonNo":1,"seats":[{"seatRow":1,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"E","seatClass":"C","status":0}]},{"wagonCode":"PREMIUM","wagonNo":2,"seats":[{"seatRow":1,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":1,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":2,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":3,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":4,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":5,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":6,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":7,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":8,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":9,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":10,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":11,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":12,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":13,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":14,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":15,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":16,"seatColumn":"E","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"A","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"B","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"C","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"D","seatClass":"C","status":0},{"seatRow":17,"seatColumn":"E","seatClass":"C","status":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	seatMapLen := 2
	trainNo := "10501"
	wagonCode0 := "PREMIUM"
	var wagonNo0 float64 = 1
	var seatRow0_0 float64 = 1

	// test variable
	var vRS InternalGetSeatNullRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, seatMapLen, len(vRS.Return.SeatNulls), "should be equal!")
	assert.Equal(t, trainNo, vRS.Return.TrainNo, "should be equal!")

	assert.Equal(t, wagonCode0, vRS.Return.SeatNulls[0].WagonCode, "should be equal!")
	assert.Equal(t, wagonNo0, vRS.Return.SeatNulls[0].WagonNo, "should be equal!")
	assert.Equal(t, seatRow0_0, vRS.Return.SeatNulls[0].Seats[0].SeatRow, "should be equal!")

	assert.Nil(t, err)
}
