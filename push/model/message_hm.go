package model

import (
	"github.com/fabian4/Fyi_sever/push/config"
)

type AndroidConfig struct {
	CollapseKey   int                  `json:"collapse_key,omitempty"`
	Urgency       string               `json:"urgency,omitempty"`
	Category      string               `json:"category,omitempty"`
	TTL           string               `json:"ttl,omitempty"`
	BiTag         string               `json:"bi_tag,omitempty"`
	FastAppTarget int                  `json:"fast_app_target,omitempty"`
	Data          string               `json:"data,omitempty"`
	Notification  *AndroidNotification `json:"notification,omitempty"`
}

type AndroidNotification struct {
	Title         string                 `json:"title,omitempty"`
	Body          string                 `json:"body,omitempty"`
	Icon          string                 `json:"icon,omitempty"`
	Color         string                 `json:"color,omitempty"`
	Sound         string                 `json:"sound,omitempty"`
	DefaultSound  bool                   `json:"default_sound,omitempty"`
	Tag           string                 `json:"tag,omitempty"`
	ClickAction   *ClickAction           `json:"click_action,omitempty"`
	BodyLocKey    string                 `json:"body_loc_key,omitempty"`
	BodyLocArgs   []string               `json:"body_loc_args,omitempty"`
	TitleLocKey   string                 `json:"title_loc_key,omitempty"`
	TitleLocArgs  []string               `json:"title_loc_args,omitempty"`
	MultiLangKey  map[string]interface{} `json:"multi_lang_key,omitempty"`
	ChannelId     string                 `json:"channel_id,omitempty"`
	NotifySummary string                 `json:"notify_summary,omitempty"`
	Image         string                 `json:"image,omitempty"`
	Style         int                    `json:"style,omitempty"`
	BigTitle      string                 `json:"big_title,omitempty"`
	BigBody       string                 `json:"big_body,omitempty"`

	AutoClear         int                `json:"auto_clear,omitempty"`
	NotifyId          int                `json:"notify_id,omitempty"`
	Group             string             `json:"group,omitempty"`
	Badge             *BadgeNotification `json:"badge,omitempty,omitempty"`
	Ticker            string             `json:"ticker,omitempty"`
	When              string             `json:"when,omitempty"`
	Importance        string             `json:"importance,omitempty"`
	UseDefaultVibrate bool               `json:"use_default_vibrate,omitempty"`
	UseDefaultLight   bool               `json:"use_default_light,omitempty"`
	VibrateConfig     []string           `json:"vibrate_config,omitempty"`
	Visibility        string             `json:"visibility,omitempty"`
	LightSettings     *LightSettings     `json:"light_settings,omitempty"`
	ForegroundShow    bool               `json:"foreground_show,omitempty"`
	Buttons           []Button           `json:"buttons,omitempty"`
}

type ClickAction struct {
	Type         int    `json:"type"` //when the type equals to 1, At least one of intent and action is not empty
	Intent       string `json:"intent,omitempty"`
	Action       string `json:"action,omitempty"`
	Url          string `json:"url,omitempty"`
	RichResource string `json:"rich_resource,omitempty"`
}

type BadgeNotification struct {
	AddNum int    `json:"add_num,omitempty"`
	SetNum int    `json:"set_num,omitempty"`
	Class  string `json:"class,omitempty"`
}

type LightSettings struct {
	Color            *Color `json:"color"`
	LightOnDuration  string `json:"light_on_duration,omitempty"`
	LightOffDuration string `json:"light_off_duration,omitempty"`
}

type Button struct {
	Name       string `json:"name"`
	ActionType int    `json:"action_type"`
	IntentType int    `json:"intent_type"`
	Intent     string `json:"intent"`
	Data       string `json:"data"`
}

type Color struct {
	Alpha int `json:"alpha"`
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

func GetDefaultAndroid() *AndroidConfig {
	android := &AndroidConfig{
		Urgency:      config.DeliveryPriorityNormal,
		TTL:          "86400s",
		CollapseKey:  -1,
		Notification: nil,
	}
	return android
}

func GetAndroidNotification(msg string, detail string, tag string) *AndroidNotification {
	notification := &AndroidNotification{
		DefaultSound:      true,
		Importance:        config.NotificationPriorityDefault,
		ClickAction:       getClickAction(),
		Body:              detail,
		Title:             msg,
		NotifySummary:     tag,
		UseDefaultVibrate: true,
		UseDefaultLight:   true,
		Visibility:        config.VisibilityPrivate,
		ForegroundShow:    true,
		Buttons:           getButtons(),
	}
	return notification
}

func getClickAction() *ClickAction {
	return &ClickAction{
		Type:   config.TypeIntentOrAction,
		Action: "Action",
	}
}

func getButtons() []Button {
	return []Button{
		Button{
			Name:       "详情",
			ActionType: 0,
		},
		Button{
			Name:       "已读",
			ActionType: 3,
		},
	}
}
