package entities

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type CompanyAddress struct {
	CompanyAddressUuid uuid.UUID `json:"company_address_uuid" gorm:"primarykey;column:company_address_uuid"`
	Zip                string    `json:"zip" gorm:"column:zip"`
	Country            string    `json:"country" gorm:"column:country"`
	State              string    `json:"state" gorm:"column:state"`
	City               string    `json:"city" gorm:"column:city"`
	Address1           string    `json:"address1" gorm:"column:address1"`
	Address2           string    `json:"address2" gorm:"column:address2"`
}

func (m *CompanyAddress) TableName() string {
	return "company_address"
}

type CompanyFacility struct {
	CompanyFacilityUuid uuid.UUID      `json:"company_facility_uuid" gorm:"primarykey;column:company_facility_uuid"`
	CompanyAddressUuid  uuid.UUID      `json:"-" gorm:"column:company_address_uuid"`
	CompanyAddress      CompanyAddress `json:"-" gorm:"foreignKey:CompanyAddressUuid;references:CompanyAddressUuid"`
	Name                string         `json:"name" gorm:"column:name"`
	Type                string         `json:"-" gorm:"column:type"`
}

func (m *CompanyFacility) TableName() string {
	return "company_facilities"
}

func (f *CompanyFacility) MarshalJSON() ([]byte, error) {
	type Alias CompanyFacility

	o := &struct {
		*Alias
		CompanyAddress string `json:"address"`
	}{
		Alias: (*Alias)(f),
	}

	if f.Type == "Others" {
		o.Name = f.Name
	} else {
		o.Name = f.Type
	}

	var address []string

	if f.CompanyAddress.Address1 != "" {
		address = append(address, f.CompanyAddress.Address1)
	}
	if f.CompanyAddress.Address2 != "" {
		address = append(address, f.CompanyAddress.Address2)
	}
	if f.CompanyAddress.City != "" {
		address = append(address, f.CompanyAddress.City)
	}
	if f.CompanyAddress.State != "" {
		address = append(address, f.CompanyAddress.State)

	}
	if f.CompanyAddress.Country != "" {
		address = append(address, f.CompanyAddress.Country)
	}
	if f.CompanyAddress.Zip != "" {
		address = append(address, f.CompanyAddress.Zip)
	}

	o.CompanyAddress = strings.Join(address, ",")

	return json.Marshal(o)
}
func (f *CompanyFacility) CompanyAddressString() string {
	var address []string

	if f.CompanyAddress.Address1 != "" {
		address = append(address, f.CompanyAddress.Address1)
	}
	if f.CompanyAddress.Address2 != "" {
		address = append(address, f.CompanyAddress.Address2)
	}
	if f.CompanyAddress.City != "" {
		address = append(address, f.CompanyAddress.City)
	}
	if f.CompanyAddress.State != "" {
		address = append(address, f.CompanyAddress.State)

	}
	if f.CompanyAddress.Country != "" {
		address = append(address, f.CompanyAddress.Country)
	}
	if f.CompanyAddress.Zip != "" {
		address = append(address, f.CompanyAddress.Zip)
	}

	return strings.Join(address, ",")
}

type TechInfoWireless struct {
	TechInfoWirelessUuid uuid.UUID         `json:"tech_info_wireless_uuid" gorm:"column:tech_info_wireless_uuid"`
	CompanyUuid          uuid.UUID         `json:"company_uuid" gorm:"column:company_uuid"`
	CompanyFacilityUuid  uuid.UUID         `json:"company_facility_uuid" gorm:"column:company_facility_uuid"`
	CompanyFacility      CompanyFacility   `json:"facility" gorm:"foreignKey:CompanyFacilityUuid;references:CompanyFacilityUuid"`
	Ssid                 string            `json:"ssid" gorm:"column:ssid"`
	Description          string            `json:"description" gorm:"column:description"`
	ProtocolType         string            `json:"protocol_type" gorm:"column:protocol_type"`
	Protocol             string            `json:"protocol" gorm:"column:protocol"`
	SecurityType         string            `json:"security_type" gorm:"column:security_type"`
	Security             string            `json:"security" gorm:"column:security"`
	Status               string            `json:"status" gorm:"column:status"`
	CreatedAt            nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt            nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy            uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy            uuid.UUID         `json:"updated_by" gorm:"column:updated_by"`
}

func (m *TechInfoWireless) TableName() string {
	return "tech_info_wireless"
}

func (t *TechInfoWireless) MarshalJSON() ([]byte, error) {
	type Alias TechInfoWireless

	o := &struct {
		*Alias
		CreatedBy           string `json:"created_by"`
		UpdatedBy           string `json:"updated_by"`
		CompanyFacilityUuid string `json:"company_facility_uuid,omitempty"`
	}{
		Alias:               (*Alias)(t),
		CreatedBy:           t.CreatedBy.String(),
		UpdatedBy:           t.UpdatedBy.String(),
		CompanyFacilityUuid: "",
	}

	if t.CreatedBy == uuid.Nil {
		o.CreatedBy = ""
	}
	if t.UpdatedBy == uuid.Nil {
		o.UpdatedBy = ""
	}

	return json.Marshal(o)
}

type CreateTechInfoWirelessRequest struct {
	CompanyUuid      uuid.UUID         `json:"company_uuid"`
	UserUuid         uuid.UUID         `json:"user_uuid"`
	TechInfoWireless *TechInfoWireless `json:"tech_info_wireless"`
}

type UpdateTechInfoWirelessRequest struct {
	CompanyUuid          uuid.UUID         `json:"company_uuid"`
	UserUuid             uuid.UUID         `json:"user_uuid"`
	TechInfoWirelessUuid uuid.UUID         `json:"tech_info_wireless_uuid"`
	TechInfoWireless     *TechInfoWireless `json:"tech_info_wireless"`
}

type GetAllTechInfoWirelesssRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetTechInfoWirelessByIdRequest struct {
	CompanyUuid          uuid.UUID `json:"company_uuid"`
	UserUuid             uuid.UUID `json:"user_uuid"`
	TechInfoWirelessUuid uuid.UUID `json:"tech_info_wireless_uuid"`
}

type DeleteTechInfoWirelessRequest struct {
	CompanyUuid          uuid.UUID `json:"company_uuid"`
	UserUuid             uuid.UUID `json:"user_uuid"`
	TechInfoWirelessUuid uuid.UUID `json:"tech_info_wireless_uuid"`
}

type CompanyWebsite struct {
	CompanyWebsiteUUID string `json:"company_website_uuid,omitempty"`
	URL                string `json:"url,omitempty"`
}

type UpdateTechInfoWirelessPatchRequestBody struct {
	Status string `json:"status"`
}

type UpdateTechInfoWirelessPatchRequest struct {
	CompanyUuid          uuid.UUID                               `json:"company_uuid"`
	UserUuid             uuid.UUID                               `json:"user_uuid"`
	TechInfoWirelessUuid uuid.UUID                               `json:"tech_info_wireless_uuid"`
	PatchRequestBody     *UpdateTechInfoWirelessPatchRequestBody `json:"patch_request_body"`
}

type UpdateTechInfoWirelessPatchResponse struct {
	TechInfoWirelessUuid uuid.UUID `json:"tech_info_wireless_uuid"`
	Status               string    `json:"status,omitempty"`
}
