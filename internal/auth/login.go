package auth

import (
	"strings"
	"time"

	"github.com/go-rod/rod"
)

type Detected struct {
	EmailField    bool
	PasswordField bool
	LoginButton   bool
	Captcha       bool
	Checkpoint    bool
}

func OpenLoginPage(page *rod.Page) error {
	if err := page.Navigate("https://www.linkedin.com/login"); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	return nil
}

func DetectLoginState(page *rod.Page) (*Detected, error) {
	d := &Detected{}

	if _, err := page.Element(`input[name="session_key"], input#username`); err == nil {
		d.EmailField = true
	}

	if _, err := page.Element(`input[name="session_password"], input#password`); err == nil {
		d.PasswordField = true
	}

	if _, err := page.Element(`button[type="submit"]`); err == nil {
		d.LoginButton = true
	}

	if _, err := page.Element(`iframe[src*="captcha"], div[id*="captcha"]`); err == nil {
		d.Captcha = true
	}

	url := page.MustInfo().URL
	if strings.Contains(url, "checkpoint") || strings.Contains(url, "challenge") {
		d.Checkpoint = true
	}

	return d, nil
}
