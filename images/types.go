package images

type colorConfig struct {
	Bg            string   `json:"bg"`
	PlayerColor   []string `json:"playerColor"`
	KDAColor      []string `json:"kdaColor"`
	Text          string   `json:"text"`
	TextDire      string   `json:"textDire"`
	TextRadiant   string   `json:"textRadiant"`
	TextDark      string   `json:"textDark"`
	TextGold      string   `json:"textGold"`
	TextDamage    string   `json:"textDamage"`
	TextHealing   string   `json:"textHealing"`
	TextXP        string   `json:"textXP"`
	TextTF        string   `json:"textTF"`
	TextTFDire    string   `json:"textTFDire"`
	TextTFRadiant string   `json:"textTFRadiant"`
	BackDire      string   `json:"backDire"`
	BackRadiant   string   `json:"backRadiant"`
	BackLost      string   `json:"backLost"`
	TextSize8     float64  `json:"textSize8"`
	TextSize12    float64  `json:"textSize12"`
	TextSize15    float64  `json:"textSize15"`
	TextSize20    float64  `json:"textSize20"`
	TextSize28    float64  `json:"textSize28"`
	TextSize35    float64  `json:"textSize35"`
	TextSize48    float64  `json:"textSize48"`
}
