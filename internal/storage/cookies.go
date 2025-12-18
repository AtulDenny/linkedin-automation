package storage

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFile = "cookies.json"

func SaveCookies(page *rod.Page) error {
	cookies, err := page.Cookies([]string{})
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cookieFile, data, 0644)
}

func LoadCookies(page *rod.Page) error {
	data, err := os.ReadFile(cookieFile)
	if err != nil {

		return nil
	}

	var stored []*proto.NetworkCookie
	if err := json.Unmarshal(data, &stored); err != nil {
		return err
	}

	var params []*proto.NetworkCookieParam
	for _, c := range stored {
		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	return page.SetCookies(params)
}
