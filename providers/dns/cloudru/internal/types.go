package internal

import (
	"fmt"
	"time"
)

type Token struct {
	// The bearer token for use in API requests
	AccessToken string `json:"access_token"`
	TokenID     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	// Number in seconds before the expiration
	ExpiresIn       int    `json:"expires_in"`
	NotBeforePolicy int    `json:"not-before-policy"`
	Scope           string `json:"scope"`

	Deadline time.Time `json:"-"`
}

type authResponseError struct {
	ErrorMsg         string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (a authResponseError) Error() string {
	return fmt.Sprintf("%s: %s", a.ErrorMsg, a.ErrorDescription)
}

type APIResponse[T any] struct {
	Items []T `json:"items"`
}

type ZoneResponse struct {
	Zone Zone `json:"zone"`
}

type Zone struct {
	Meta            ZoneMeta  `json:"meta"`
	ProjectID       string    `json:"projectId"`
	Name            string    `json:"name"`
	Domain          string    `json:"domain"`
	ReadOnly        bool      `json:"readOnly"`
	State           string    `json:"state"`
	CountRecords    string    `json:"countRecords"`
	CountValues     string    `json:"countValues"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	ConfirmState    string    `json:"confirmState"`
	NameServers     []string  `json:"NameServers"`
	OwnerProductCode string   `json:"ownerProductCode"`
	IsProductOwner  bool      `json:"isProductOwner"`
}

type ZoneMeta struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"taskId"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Record struct {
	ZoneID  string   `json:"zoneId,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    string   `json:"type,omitempty"`
	Values  []string `json:"values,omitempty"`
	TTL     int      `json:"ttl,omitempty"`
}
type RecordMeta struct {
	Task RecordTask `json:"task"`
}

type RecordTask struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	EntityID  string    `json:"entityId"`
	Entity    string    `json:"entity"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
}
