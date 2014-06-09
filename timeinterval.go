package goqqpushapi

type MiniTime struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
}

func (this *MiniTime) IsValid() bool {
	return this.Hour >= 0 && this.Hour <= 23 && this.Min >= 0 && this.Min <= 59
}

type TimeInterval struct {
	Start MiniTime `json:"start"`
	End   MiniTime `json:"end"`
}

func (this *TimeInterval) IsValid() bool {
	return this.Start.IsValid() && this.End.IsValid()
}
