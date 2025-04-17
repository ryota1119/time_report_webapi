package entities

import "strconv"

// OrganizationID 組織ID
type OrganizationID uint

type OrganizationName string

type OrganizationCode string

// Organization 組織情報
type Organization struct {
	ID   OrganizationID
	Name OrganizationName
	Code OrganizationCode
}

// NewOrganization は指定された組織名と組織コードから Organization を生成する
func NewOrganization(organizationName, organizationCode string) *Organization {
	return &Organization{
		Name: OrganizationName(organizationName),
		Code: OrganizationCode(organizationCode),
	}
}

// String は OrganizationID をstring型にキャストする
func (uid OrganizationID) String() string {
	return strconv.FormatUint(uint64(uid), 10)
}

// String は OrganizationCode をstring型にキャストする
func (code OrganizationCode) String() string {
	return string(code)
}

// CachedOrganization Redis用キャッシュ組織情報
type CachedOrganization struct {
	ID   OrganizationID   `json:"id"`
	Name OrganizationName `json:"name"`
	Code OrganizationCode `json:"code"`
}
