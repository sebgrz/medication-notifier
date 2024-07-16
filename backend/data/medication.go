package data

type Medication struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Day       string `json:"day"`
	TimeOfDay string `json:"time_of_day"`
}

type MedicationDataService interface {
	Add(Medication) error
	FindByUserId(string) []Medication
	FindById(string) (*Medication, error)
	RemoveById(string) error
	RemoveByUserId(string) error
}
