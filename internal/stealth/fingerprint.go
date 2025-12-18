package stealth

import "github.com/go-rod/rod"

func DisableWebDriver(page *rod.Page) error {
	_, err := page.Eval(`
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});
	`)
	return err
}
