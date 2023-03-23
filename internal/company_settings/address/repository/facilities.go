package repository

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
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
	Status              string         `json:"status" gorm:"column:status"`
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
