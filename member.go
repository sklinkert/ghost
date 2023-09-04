package ghost

import (
	"encoding/json"
	"fmt"
	"time"
)

type Members struct {
	Members []Member `json:"members"`
}

type NewMembers struct {
	Members []NewMember `json:"members"`
}

type Member struct {
	Id          string      `json:"id"`
	Uuid        string      `json:"uuid"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	Note        interface{} `json:"note"`
	Geolocation interface{} `json:"geolocation"`
	Subscribed  bool        `json:"subscribed"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Labels      []struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"labels"`
	Subscriptions    []Subscription `json:"subscriptions"`
	AvatarImage      string         `json:"avatar_image"`
	EmailCount       int            `json:"email_count"`
	EmailOpenedCount int            `json:"email_opened_count"`
	EmailOpenRate    float64        `json:"email_open_rate"`
	Status           string         `json:"status"`
}

type NewMember struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Subscription struct {
	Id       string `json:"id"`
	Customer struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"customer"`
	Status                  string    `json:"status"`
	StartDate               time.Time `json:"start_date"`
	DefaultPaymentCardLast4 string    `json:"default_payment_card_last4"`
	CancelAtPeriodEnd       bool      `json:"cancel_at_period_end"`
	CancellationReason      string    `json:"cancellation_reason"`
	CurrentPeriodEnd        time.Time `json:"current_period_end"`
	Price                   struct {
		Id       string `json:"id"`
		PriceId  string `json:"price_id"`
		Nickname string `json:"nickname"`
		Amount   int    `json:"amount"`
		Interval string `json:"interval"`
		Type     string `json:"type"`
		Currency string `json:"currency"`
	} `json:"price"`
}

func (g *Ghost) AdminGetMembers() (Members, error) {
	var membersResponse Members
	const ghostPostsURLSuffix = "%s/ghost/api/v3/admin/members/?key=%s&limit=all"

	if err := g.checkAndRenewJWT(); err != nil {
		return membersResponse, err
	}

	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.adminAPIToken)
	if err := g.getJson(url, &membersResponse); err != nil {
		return membersResponse, err
	}
	return membersResponse, nil
}

func (g *Ghost) AdminCreateMember(member NewMember) (Members, error) {
	const ghostPostsURLSuffix = "%s/ghost/api/v3/admin/members/?key=%s"
	var members Members

	if err := g.checkAndRenewJWT(); err != nil {
		return members, err
	}

	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.adminAPIToken)
	data, err := json.Marshal(&NewMembers{Members: []NewMember{member}})
	if err != nil {
		return members, err
	}

	if err := g.postJson(url, data, &members); err != nil {
		return members, err
	}
	return members, nil
}
