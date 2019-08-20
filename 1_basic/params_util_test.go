package params_test

import (
	"bytes"
	"testing"

	params "github.com/aimzeter/wut/1_basic"
)

func TestNiK(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantErr bool
		wantNIK string
	}{
		{
			name: "valid body",
			body: `
				{
					"nik": "1234567"
				}
			`,
			wantErr: false,
			wantNIK: "1234567",
		},
		{
			name: "invalid body",
			body: `
				{
					"other_id": 1
				}
			`,
			wantErr: true,
			wantNIK: "",
		},
		{
			name:    "empty body",
			body:    `null`,
			wantErr: true,
			wantNIK: "",
		},
		{
			name: "invalid json",
			body: `
				{
					invalid_field: 1
				}`,
			wantErr: true,
			wantNIK: "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			body := bytes.NewBufferString(tc.body)
			id, err := params.GetNIK(body)
			assertError(t, tc.wantErr, err)
			assertNIK(t, tc.wantNIK, id)
		})
	}
}

func TestGetStudentID(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantErr bool
		wantID  uint64
	}{
		{
			name: "valid body",
			body: `
				{
					"student_id": 1
				}
			`,
			wantErr: false,
			wantID:  1,
		},
		{
			name: "invalid body",
			body: `
				{
					"other_id": 1
				}
			`,
			wantErr: true,
			wantID:  0,
		},
		{
			name:    "empty body",
			body:    `null`,
			wantErr: true,
			wantID:  0,
		},
		{
			name: "invalid json",
			body: `
				{
					invalid_field: 1
				}`,
			wantErr: true,
			wantID:  0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			body := bytes.NewBufferString(tc.body)
			id, err := params.GetStudentID(body)
			assertError(t, tc.wantErr, err)
			assertID(t, tc.wantID, id)
		})
	}
}

func assertError(t *testing.T, want bool, err error) {
	t.Helper()
	got := err != nil
	if want != got {
		if want {
			t.Fatalf("❌ FAIL ❌: GetStudentID should return error\n")
		}
		t.Fatalf("❌ FAIL ❌: GetStudentID should not return error, got error %s\n", err.Error())
	}
}

func assertID(t *testing.T, want, got uint64) {
	t.Helper()
	if want != got {
		t.Fatalf("❌ FAIL ❌: GetStudentID did not return corrent id.\n"+
			"\twant\t:\t%d\n"+
			"\tgot\t:\t%d\n", want, got)
	}
}

func assertNIK(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Fatalf("❌ FAIL ❌: GetNIK did not return corrent id.\n"+
			"\twant\t:\t%s\n"+
			"\tgot\t:\t%s\n", want, got)
	}
}
