package helper

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail     = fmt.Errorf("invalid_email")
	ErrInvalidPhone     = fmt.Errorf("invalid_phone")
	ErrInvalidImageType = fmt.Errorf("invalid_image_type")
	ErrInvalidVideoType = fmt.Errorf("invalid_video_type")
	ErrInvalidIconType  = fmt.Errorf("invalid_icon_type")
)

func IsValidPhone(phone string) bool {
	if containsPlus := strings.Contains(phone, "+"); containsPlus {
		return false
	}

	// phone = strings.Replace(phone, "+", "", -1)
	r := regexp.MustCompile(`^998[0-9]{2}[0-9]{7}$`)
	return r.MatchString(phone)
}

func IsValidEmail(email string) bool {
	r := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return r.MatchString(email)
}

func IsValidLogin(login string) bool {
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,29}$`)
	return r.MatchString(login)
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func IsValidPassword(s string) bool {
	hasLetter := regexp.MustCompile(`[a-zA-Z]`)
	hasSpecial := regexp.MustCompile(`[:\-+_=%#@!^&<*,.]`)
	hasLength := len(s) >= 6
	return hasLength && hasSpecial.MatchString(s) && hasLetter.MatchString(s)
}
