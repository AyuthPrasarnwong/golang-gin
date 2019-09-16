package models

type VerifyOtp struct {
	OtpCode  string `json:"otp_code"`
	OtpRef   string `json:"otp_ref"`
	AuthCode string `json:"auth_code"`
	Mobile   string `json:"mobile"`
}
