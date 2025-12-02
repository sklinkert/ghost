package ghost

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Members struct {
	Members []Member   `json:"members"`
	Meta    Pagination `json:"meta,omitempty"`
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
	Subscriptions    []Subscription   `json:"subscriptions"`
	AvatarImage      string           `json:"avatar_image"`
	Comped           bool             `json:"comped"`
	EmailCount       int              `json:"email_count"`
	EmailOpenedCount int              `json:"email_opened_count"`
	EmailOpenRate    *float64         `json:"email_open_rate"`
	Status           string           `json:"status"`
	LastSeenAt       *time.Time       `json:"last_seen_at"`
	UnsubscribeUrl   string           `json:"unsubscribe_url"`
	Tiers            []Tier           `json:"tiers"`
	EmailSuppression EmailSuppression `json:"email_suppression"`
	Newsletters      []Newsletter     `json:"newsletters"`
	Attribution      *Attribution     `json:"attribution,omitempty"`
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

type Tier struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Newsletter struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type EmailSuppression struct {
	Suppressed bool        `json:"suppressed"`
	Info       interface{} `json:"info"`
}

type Attribution struct {
	Id             interface{} `json:"id"`
	Type           string      `json:"type"`
	Url            string      `json:"url"`
	Title          string      `json:"title"`
	ReferrerSource string      `json:"referrer_source"`
	ReferrerMedium interface{} `json:"referrer_medium"`
	ReferrerUrl    interface{} `json:"referrer_url"`
}

func (g *Ghost) AdminGetMembers() (Members, error) {
	const limit = 100
	var allMembers Members
	page := 1

	for {
		url := fmt.Sprintf("%s/ghost/api/v3/admin/members/?limit=%d&page=%d", g.url, limit, page)

		var pageMembers Members
		if err := g.getJson(url, &pageMembers); err != nil {
			return allMembers, err
		}

		allMembers.Members = append(allMembers.Members, pageMembers.Members...)

		if pageMembers.Meta.Pagination.Next == nil {
			break
		}
		page = *pageMembers.Meta.Pagination.Next
	}

	return allMembers, nil
}

func (g *Ghost) AdminCreateMember(member NewMember) (Members, error) {
	const ghostPostsURLSuffix = "%s/ghost/api/v3/admin/members/?key=%s"
	var members Members

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

func (g *Ghost) AdminGetMember(memberId string) (Members, error) {
	var members Members
	url := fmt.Sprintf("%s/ghost/api/v3/admin/members/%s/?include=tiers", g.url, memberId)

	if err := g.getJson(url, &members); err != nil {
		return members, err
	}
	return members, nil
}

func (g *Ghost) AdminDeleteMember(memberId string) error {
	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	deleteURL := fmt.Sprintf("%s/ghost/api/v3/admin/members/%s/", g.url, memberId)

	req, err := http.NewRequest(http.MethodDelete, deleteURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+g.jwtToken)
	resp, err := g.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	content, _ := io.ReadAll(resp.Body)
	responseBody := string(content)
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}
