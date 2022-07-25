package telbiz

type telBizReqOAuth struct {
	ClientID  string `json:"ClientID" binding:"required"`
	Secret    string `json:"Secret" binding:"required"`
	GrantType string `json:"GrantType" binding:"required"`
	Scope     string `json:"Scope" binding:"required"`
}

type telBizResOAuth struct {
	AccessToken  string `json:"accessToken"`
	Expire       int    `json:"expire"`
	RefreshToken string `json:"refreshToken"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	Success      bool   `json:"success"`
	Detail       string `json:"detail"`
}

type TelBizReqSMS struct {
	Title   string `json:"Title" binding:"required"`
	Phone   string `json:"Phone" binding:"required"`
	Message string `json:"Message" binding:"required"`
}

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
}

type key struct {
	PartitionKey string `json:"partitionKey"`
	RangeKey     string `json:"rangeKey"`
}

type TelBizResSMS struct {
	Response response `json:"response"`
	Key      key      `json:"key"`
}

type TelBizReqTopUp struct {
	Phone  string  `json:"Phone" binding:"required"`
	Amount float64 `json:"Amount" binding:"required"`
}

type Title string

const (
	Default   Title = "Telbiz"
	News            = "News"
	Promotion       = "Promotion"
	OTP             = "OTP"
	Info            = "Info"
	Unknown         = "Unknown"
)
