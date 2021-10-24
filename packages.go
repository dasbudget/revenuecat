package revenuecat

import (
	"fmt"
)

type Package struct {
	ID                        string    `json:"id,omitempty"`
	Identifier                string    `json:"identifier"`
	PlatformProductIdentifier string    `json:"platform_product_identifier,omitempty"`
	DisplayName               string    `json:"display_name"`
	OfferingID                string    `json:"offering_id"`
	Products                  []Product `json:"products,omitempty"`
}

type Product struct {
	CreatedAt  string `json:"created_at"`
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Store      string `json:"store"`
}

func (c *Client) CreatePackage(appID string, p *Package) (Package, error) {
	resp := Package{}
	err := c.call("POST", "developers/me/apps/"+appID+"/new_packages", p, "", &resp)
	return resp, err
}

func (c *Client) AttachProduct(appID, pkgID string, productIDs ...string) (Package, error) {
	body := struct {
		ProductsIDs []string `json:"products_ids"`
	}{
		ProductsIDs: productIDs,
	}

	resp := Package{}
	err := c.call("POST", fmt.Sprintf("developers/me/apps/%s/new_packages/%s/attach_products", appID, pkgID), body, "", &resp)
	return resp, err
}

func (c *Client) GetProducts(appID string) ([]Product, error) {
	var resp []Product
	err := c.call("GET", fmt.Sprintf("developers/me/apps/%s/new_products", appID), nil, "", &resp)
	return resp, err
}
