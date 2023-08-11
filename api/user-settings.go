package api

import "fmt"

type UserSettingKey string

const (
	UserSettingLocaleKey   UserSettingKey = "locale"
	UserSettingThemeKey    UserSettingKey = "theme"
	UserSettingMoodViewKey UserSettingKey = "mood_view"
)

var (
	UserSettingLocaleValue   = []string{"en", "hi"}
	UserSettingThemeValue    = []string{"system", "light", "dark"}
	UserSettingMoodViewValue = []string{"list", "grid"}
)

func (k UserSettingKey) String() string {
	switch k {
	case UserSettingLocaleKey:
		return "locale"
	case UserSettingThemeKey:
		return "theme"
	case UserSettingMoodViewKey:
		return "mood_view"
	}
	return ""
}

type UserSetting struct {
	UserId       int    `json:"user_id"`
	SettingKey   string `json:"setting_key"`
	SettingValue string `json:"setting_value"`
}

type UpsertUserSetting struct {
	UserId       int            `json:"_"`
	SettingKey   UserSettingKey `json:"setting_key"`
	SettingValue string         `json:"setting_value"`
}

func (u *UpsertUserSetting) Validate() error {
	switch u.SettingKey {
	case UserSettingThemeKey:
		if ok := contains(UserSettingThemeValue, u.SettingValue); !ok {
			return fmt.Errorf("invalid theme value")
		}
	case UserSettingLocaleKey:
		if ok := contains(UserSettingThemeValue, u.SettingValue); !ok {
			return fmt.Errorf("invalid locale value")
		}
	case UserSettingMoodViewKey:
		if ok := contains(UserSettingMoodViewValue, u.SettingValue); !ok {
			return fmt.Errorf("invalid mood_view value")
		}
	}
	if u.SettingKey.String() == "" {
		return fmt.Errorf("invalid user setting key")
	}
	return nil
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
