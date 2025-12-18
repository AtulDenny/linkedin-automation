package main

import (
	"time"

	"linkedin-automation-poc/internal/auth"
	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/connect"
	"linkedin-automation-poc/internal/logger"
	"linkedin-automation-poc/internal/messaging"
	"linkedin-automation-poc/internal/search"
	"linkedin-automation-poc/internal/storage"
)

func main() {
	log := logger.New()
	log.Info("Launching stealth browser")

	// ----------------------------
	// STEP 1: SEARCH MODULE DEMO
	// ----------------------------
	searchSvc := search.New(log)

	criteria := search.Criteria{
		JobTitle: "Software Engineer",
		Company:  "DemoCorp",
		Location: "Remote",
		Page:     1,
	}

	profiles, err := searchSvc.Search(criteria)
	if err != nil {
		log.Error(err.Error())
		return
	}

	profiles = search.Deduplicate(profiles)

	log.Info("Search results:")
	for _, p := range profiles {
		log.Info(p.Name + " | " + p.URL)
	}

	// ----------------------------
	// STEP 2: CONNECTION REQUESTS
	// ----------------------------
	connectSvc := connect.New(log, 5) // daily limit = 5

	for _, p := range profiles {
		note := "Hi " + p.Name + ", I'd love to connect and exchange ideas."
		_ = connectSvc.Send(p, note)
	}
	// ----------------------------
	// STEP 3: MESSAGING SYSTEM
	// ----------------------------
	msgSvc := messaging.New(log)

	template := "Hi {{name}}, thanks for connecting! Looking forward to learning more about {{company}}."

	for _, p := range profiles {
		_ = msgSvc.SendFollowUp(p, template)
	}

	// ----------------------------
	// STEP 3: BROWSER + AUTH FLOW
	// ----------------------------
	br, err := browser.Launch()
	if err != nil {
		log.Error(err.Error())
		return
	}

	page := br.MustPage()

	// Load cookies if available
	_ = storage.LoadCookies(page)

	log.Info("Opening LinkedIn login page")
	if err := auth.OpenLoginPage(page); err != nil {
		log.Error(err.Error())
		return
	}

	state, err := auth.DetectLoginState(page)
	if err != nil {
		log.Error(err.Error())
		return
	}

	// CAPTCHA GUARD (IMPORTANT)
	if state.Captcha {
		log.Warn("CAPTCHA detected. Aborting automated login.")
		time.Sleep(10 * time.Second)
		_ = storage.SaveCookies(page)
		return
	}

	log.Info("Login page state:")
	log.Info("Email field present: " + boolStr(state.EmailField))
	log.Info("Password field present: " + boolStr(state.PasswordField))
	log.Info("Login button present: " + boolStr(state.LoginButton))
	log.Info("Checkpoint detected: " + boolStr(state.Checkpoint))

	// Save cookies before exit
	_ = storage.SaveCookies(page)

	log.Info("Demo running. Press CTRL+C to exit.")
	select {}
}

func boolStr(b bool) string {
	if b {
		return "YES"
	}
	return "NO"
}
