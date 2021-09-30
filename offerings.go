package revenuecat

import (
	"fmt"
	"net/http"
)

type Package struct {
	Identifier                string `json:"identifier"`
	PlatformProductIdentifier string `json:"platform_product_identifier"`
}

type Offering struct {
	Description string    `json:"description"`
	Identifier  string    `json:"identifier"`
	Packages    []Package `json:"packages"`
}

// OverrideOffering overrides the current Offering for a specific user.
// https://docs.revenuecat.com/reference#override-offering
func (c *Client) OverrideOffering(userID string, offeringUUID string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}
	err := c.call("POST", "subscribers/"+userID+"/offerings/"+offeringUUID+"/override", nil, "", &resp)
	return resp.Subscriber, err
}

// DeleteOfferingOverride reset the offering overrides back to the current offering for a specific user.
// https://docs.revenuecat.com/reference#delete-offering-override
func (c *Client) DeleteOfferingOverride(userID string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}
	err := c.call("DELETE", "subscribers/"+userID+"/offerings/override", nil, "", &resp)
	return resp.Subscriber, err
}

// GetOfferings gets the offerings and current offering ID for your app for a specific user
// https://docs.revenuecat.com/reference#get-offerings
func (c *Client) GetOfferings(userID, platform string) ([]Offering, string, error) {
	resp := struct {
		CurrentOfferingId string     `json:"current_offering_id"`
		Offerings         []Offering `json:"offerings"`
	}{}

	err := c.call(http.MethodGet, fmt.Sprintf("subscribers/%s/offerings", userID), nil, platform, &resp)
	return resp.Offerings, resp.CurrentOfferingId, err
}
