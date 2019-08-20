package main

import (
	"bytes"
	"fmt"

	params "github.com/aimzeter/wut/1_basic"
)

func main() {
	GetStudentIDSpec()
}

func GetStudentIDSpec() {
	body := bytes.NewBufferString(`
		{
			"student_id": 2
		}
	`)
	id, err := params.GetStudentID(body)

	if err != nil {
		fmt.Printf("❌ FAIL ❌: GetStudentID should not return error, got error %s\n", err.Error())
		return
	}

	if id != 1 {
		fmt.Printf("❌ FAIL ❌: GetStudentID did not return corrent id.\n"+
			"\twant\t:\t%d\n"+
			"\tgot\t:\t%d\n", 1, id)
		return
	}

	fmt.Println("✅ PASS ✅: GetStudentID return expected output")
}
