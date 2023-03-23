package entities

const (
	// Onboarding Step:"Invite Other Users"
	InviteOtherUsers = iota
	// Onboarding Step:"Schedule Meetings"
	ScheduleMeetings
	// Onboarding Step:"Company Info"
	CompanyInfo
	// Onboarding Step:"Set Technical Information"
	SetTechnicalInformation
	// Onboarding Step:"Sign Authorization"
	SignAuthorization
	// Onboarding Step:"Upload Policies and Procedures"
	UploadPoliciesAndProcedures
	// Onboarding Step:"Security Awareness Training"
	SecurityAwarenessTraining
)

const (
	OnboardingStatusDraft    string = "DRAFT"
	OnboardingStatusComplete string = "COMPLETE"
	OnboardingStatusCacheKey string = "ONBOARDING_STEPS"
)

/*
Map the onboarding enum item to its name
E.g InviteOtherUsers => Invite Other Users",
*/

const (
	InviteOtherUsersStep            string = "Invite Other Users"
	ScheduleMeetingsStep            string = "Schedule Meetings"
	CompanyInfoStep                 string = "Enter Company Information"
	SetTechnicalInformationStep     string = "Enter Technology Information"
	SignAuthorizationStep           string = "Sign Authorization Forms"
	UploadPoliciesAndProceduresStep string = "Upload Policies and Procedures"
	SecurityAwarenessTrainingStep   string = "Setup Security Awareness Training"
)

// TODO update onboarding client to use constants insted of map logic
var OnboardinStepToName map[int]string = map[int]string{
	InviteOtherUsers:            InviteOtherUsersStep,
	ScheduleMeetings:            ScheduleMeetingsStep,
	CompanyInfo:                 CompanyInfoStep,
	SetTechnicalInformation:     SetTechnicalInformationStep,
	SignAuthorization:           SignAuthorizationStep,
	UploadPoliciesAndProcedures: UploadPoliciesAndProceduresStep,
	SecurityAwarenessTraining:   SecurityAwarenessTrainingStep,
}
