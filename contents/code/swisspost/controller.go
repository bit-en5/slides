package swisspost

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/go-resty/resty/v2"
)

//START1 OMIT

func (s Swisspost) getURL(bfsNr int) (url string) {
	url = apiURL + dataset +
		"?" + dSelect +
		"&" + dWhere + fmt.Sprint(bfsNr) +
		"&" + dLimit +
		"&" + dOffset +
		"&" + dTZ +
		"&" + dKey + s.apiKey

	return
}

//END1 OMIT
//START2 OMIT

func (s Swisspost) GetTown(bfsNr int) (string, string, error) {
	url := s.getURL(bfsNr)

	resp, err := resty.New().SetRetryCount(3).R().Get(url)
	if err != nil {
		return "", "", err
	}

	var response Response

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return "", "", errors.New("could not unmarshal the JSON structure")
	}

	if len(response.Records) == 0 {
		return "", "", errors.New("no data found")
	}
	f := response.Records[0].Record.Fields

	return f.Kanton, f.GemeindeName, nil
}

//END2 OMIT

func (s Swisspost) GetTown2(bfsNr int) (string, string, error) {
	url := s.getURL(bfsNr)

	//START3 OMIT

	resp, err := resty.New().SetRetryCount(3).R().Get(url)
	if err != nil {
		return "", "", err
	}

	//END3 OMIT
	//START4 OMIT

	var response Response

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return "", "", errors.New("could not unmarshal the JSON structure")
	}

	//END4 OMIT

	if len(response.Records) == 0 {
		return "", "", errors.New("no data found")
	}
	f := response.Records[0].Record.Fields

	return f.Kanton, f.GemeindeName, nil
}

func (s Swisspost) getURL2(bfsNr int) (url string) {
	url = "" + `
	//START5 OMIT
	https://swisspost.opendatasoft.com/api/v2/catalog/datasets/politische-gemeinden_v2/records?
	where=bfsnr%3D<nr>&apiKey=<key>
	//END5 OMIT
	`

	return
}

func (s Swisspost) randomize() (res []string) {
	//START6 OMIT

	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	//END6 OMIT
	return
}
