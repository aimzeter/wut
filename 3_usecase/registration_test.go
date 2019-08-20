package zone_test

import (
	"testing"

	zone "github.com/aimzeter/wuts/3_usecase"
	"github.com/stretchr/testify/assert"
)

type CStoreMock struct {
	GetResp zone.Citizen
}

func (s *CStoreMock) Get(string) zone.Citizen {
	return s.GetResp
}

type PStoreMock struct {
	CreatedP zone.Participant
}

func (s *PStoreMock) Get(string) zone.Participant {
	return zone.Participant{}
}

func (s *PStoreMock) Create(p zone.Participant) error {
	s.CreatedP = p
	return nil
}

func (s *PStoreMock) Update(p zone.Participant) error {
	return nil
}

func TestRegister(t *testing.T) {
	nik := "1234567"
	autoReg := true

	cStore := &CStoreMock{}
	pStore := &PStoreMock{}

	cStore.GetResp = zone.Citizen{
		NIK: "1234567",
	}

	app := zone.AppUsecase{CStore: cStore, PStore: pStore}
	err := app.Register(nik, autoReg)

	assert.NoError(t, err)
	assert.Equal(t, nik, pStore.CreatedP.PersonalInfo.NIK)
}
