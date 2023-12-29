package tdd_doubles

type Applicant interface {
	IsOver17() bool
	HoldLicence() bool
}

type NewApplicant struct{}

func (a *NewApplicant) IsOver17() bool {
	//TODO implement me
	panic("implement me")
}

func (a *NewApplicant) HoldLicence() bool {
	//TODO implement me
	panic("implement me")
}
