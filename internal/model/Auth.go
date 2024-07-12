package model

type LoginType struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hcaptcha string `json:"h-captcha-response"`
	Gcaptcha string `json:"g-recaptcha-response"`
}

type RegisterType struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hcaptcha string `json:"h-captcha-response"`
	Gcaptcha string `json:"g-recaptcha-response"`
}
