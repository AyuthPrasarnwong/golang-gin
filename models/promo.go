package models

type Promo struct {
	CampaignId string `json:"campaign_id"`
	Amount     string `json:"amount"`
	PromoCode  string `json:"promo_code"`
	MemberId   string `json:"member_id"`
}
