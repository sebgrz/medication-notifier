package db

import (
	"errors"
	"medication-notifier/data"
	"slices"
)

type DummyMedicationDataService struct {
	medications []data.Medication
}

func NewDummyMedicationDataService() DummyMedicationDataService {
	return DummyMedicationDataService{
		medications: []data.Medication{},
	}
}

func (s *DummyMedicationDataService) Add(medication data.Medication) error {
	s.medications = append(s.medications, medication)

	return nil
}

func (s *DummyMedicationDataService) FindByUserId(userId string) []data.Medication {
	medications := []data.Medication{}
	for _, medication := range s.medications {
		if medication.UserId == userId {
			medications = append(medications, medication)
		}
	}

	return medications
}

func (s *DummyMedicationDataService) FindById(id string) (*data.Medication, error) {
	idx := slices.IndexFunc(s.medications, func(u data.Medication) bool {
		return u.Id == id
	})

	if idx < 0 {
		return nil, errors.New("medication_not_exists")
	}

	medicationCopy := s.medications[idx]
	return &medicationCopy, nil
}

func (s *DummyMedicationDataService) RemoveById(id string) error {
	_ = slices.DeleteFunc(s.medications, func(u data.Medication) bool {
		return u.Id == id
	})

	return nil
}

func (s *DummyMedicationDataService) RemoveByUserId(userId string) error {
	_ = slices.DeleteFunc(s.medications, func(u data.Medication) bool {
		return u.UserId == userId
	})

	return nil
}
