package params

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func GetNIK(body io.Reader) (string, error) {
	var nik string
	bodyMap, err := extractBody("nik", body)
	if err != nil {
		return nik, err
	}

	err = json.Unmarshal(bodyMap["nik"], &nik)
	return nik, err
}

func GetStudentID(body io.Reader) (uint64, error) {
	return extractId("student_id", body)
}

func extractId(fieldName string, body io.Reader) (uint64, error) {
	var id uint64
	bodyMap, err := extractBody(fieldName, body)
	if err != nil {
		return id, err
	}

	err = json.Unmarshal(bodyMap[fieldName], &id)
	return id, err
}

func extractBody(fieldName string, body io.Reader) (map[string]json.RawMessage, error) {
	var bodyMap map[string]json.RawMessage
	var err error

	err = json.NewDecoder(body).Decode(&bodyMap)
	if err != nil {
		return bodyMap, err
	}

	if bodyMap == nil {
		return bodyMap, errors.New("body cant be empty")
	}

	if bodyMap[fieldName] == nil {
		m := fmt.Sprint(fieldName, " cant be empty")
		return bodyMap, errors.New(m)
	}

	return bodyMap, nil
}
