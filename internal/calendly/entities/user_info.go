package entities

import "time"

type UserMeResponse struct {
	Resource UserMe `json:"resource"`
}
type UserMe struct {
	AvatarURL           string    `json:"avatar_url"`
	CreatedAt           time.Time `json:"created_at"`
	CurrentOrganization string    `json:"current_organization"`
	Email               string    `json:"email"`
	Name                string    `json:"name"`
	SchedulingURL       string    `json:"scheduling_url"`
	Slug                string    `json:"slug"`
	Timezone            string    `json:"timezone"`
	UpdatedAt           time.Time `json:"updated_at"`
	URI                 string    `json:"uri"`
}
