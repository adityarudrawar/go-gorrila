package validations

import (
	"github.com/adityarudrawar/go-gorilla/structures"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i + 10
	}
	return a
}

func ValidInputGorillaBody(ipg *structures.InputGorilla) structures.OutputGorilla {

	var opg structures.OutputGorilla
	opg.Id = ipg.Id

	if ipg.HasFriends {
		opg.Message = "Monke Good"
	} else {
		opg.Message = "Moke Not Good"
	}

	if ipg.Medicals.BloodPressure < 90 && !ipg.Medicals.HasSugar && ipg.Medicals.Resp == "GOOD" {
		opg.IsHealthy = true
	} else {
		opg.IsHealthy = false
	}

	numOfEle := len(ipg.HasLivedIn)

	opg.Numbering = makeRange(0, numOfEle)
	opg.IsError = false
	opg.Status = "ok"

	return opg
}
