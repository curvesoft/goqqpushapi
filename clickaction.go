package goqqpushapi

const (
	TYPE_ACTIVITY = 1
	TYPE_URL      = 2
	TYPE_INTENT   = 3
)

type MiniBrower struct {
	Url          string `json:"url"`
	ConfirmOnUrl int    `json:"confirm"`
}

type ClickAction struct {
	ActionType int        `json:"action_type"`
	Brower     MiniBrower `json:"brower"`
	Activity   string     `json:"activity"`
	Intent     string     `json:"intent"`
}

func (this *ClickAction) IsValid() bool {
	if this.ActionType < TYPE_ACTIVITY || this.ActionType > TYPE_INTENT {
		return false
	}

	if this.ActionType == TYPE_URL {
		if this.Brower.Url == "" || this.Brower.ConfirmOnUrl < 0 || this.Brower.ConfirmOnUrl > 1 {
			return false
		}
		return true
	}
	if this.ActionType == TYPE_INTENT {
		if this.Intent == "" {
			return false
		}
		return true
	}
	return true
}
