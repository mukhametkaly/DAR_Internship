package Internship

type InternshipCollection interface {
	AddInternship(internship *Internship) (*Internship, error)
	GetInternships() ([]*Internship,error)
	GetInternship(id int64) (*Internship, error)
	UpdateInternship(internship *Internship) (*Internship, error)
	DeleteInternship(internship *Internship) error
}

type Internship struct {

	InternshipID int64 `json:"internship_id"`
	Title string `json:"title,omitempty"`
	StartTime string `json:"starttime,omitempty"`
	EndTime string `json:"endtime,omitempty"`
}

