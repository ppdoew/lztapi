package dto

import (
	"time"
)

type AllUserAccountsOptional struct {
	UserID     int        `url:"user_id,omitempty"`
	CategoryID CategoryID `url:"category_id,omitempty"`

	Page         int        `url:"page,omitempty"`
	Show         ItemStatus `url:"show,omitempty"`
	DeleteReason string     `url:"delete_reason,omitempty"`
	Title        string     `url:"title,omitempty"`
	Pmin         int        `url:"pmin,omitempty"`
	Pmax         int        `url:"pmax,omitempty"`
	Login        string     `url:"login,omitempty"`

	Origin    []ItemOrigin `url:"origin[],omitempty"`
	NotOrigin []ItemOrigin `url:"not_origin[],omitempty"`
	OrderBy   ItemOrder    `url:"order_by,omitempty"`

	SoldBefore        bool   `url:"sb,omitempty"`
	SoldByMeBefore    bool   `url:"sb_by_me,omitempty"`
	NotSoldBefore     bool   `url:"nsb,omitempty"`
	NotSoldByMeBefore bool   `url:"nsb_by_me,omitempty"`
	Username          string `url:"username,omitempty"`

	PublishedStartDate    time.Time `url:"published_startDate,omitempty"`
	PublishedEndDate      time.Time `url:"published_endDate,omitempty"`
	FilterByPublishedDate bool      `url:"filter_by_published_date,omitempty"`

	BuyerOperationDateStart    time.Time `url:"buyer_operation_dateStart,omitempty"`
	BuyerOperationDateEnd      time.Time `url:"buyer_operation_dateEnd,omitempty"`
	FilterByBuyerOperationDate bool      `url:"filter_by_buyer_operation_date,omitempty"`

	DeleteDateStart    time.Time `url:"delete_dateStart,omitempty"`
	DeleteDateEnd      time.Time `url:"delete_dateEnd,omitempty"`
	FilterByDeleteDate bool      `url:"filter_by_delete_date,omitempty"`
}

type AllUserPurchasedOptional struct {
	UserID     int        `url:"user_id,omitempty"`
	CategoryID CategoryID `url:"category_id,omitempty"`

	Page  int        `url:"page,omitempty"`
	Show  ItemStatus `url:"show,omitempty"`
	Title string     `url:"title,omitempty"`
	Pmin  int        `url:"pmin,omitempty"`
	Pmax  int        `url:"pmax,omitempty"`
	Login string     `url:"login,omitempty"`

	Origin    []ItemOrigin `url:"origin[],omitempty"`
	NotOrigin []ItemOrigin `url:"not_origin[],omitempty"`
	OrderBy   ItemOrder    `url:"order_by,omitempty"`

	SB      bool `url:"sb,omitempty"`
	SBByMe  bool `url:"sb_by_me,omitempty"`
	NSB     bool `url:"nsb,omitempty"`
	NSBByMe bool `url:"nsb_by_me,omitempty"`
}

type AllFavouritesAccountsOptional struct {
	Page      int          `url:"page,omitempty"`
	Show      ItemStatus   `url:"show,omitempty"`
	Title     string       `url:"title,omitempty"`
	Pmin      int          `url:"pmin,omitempty"`
	Pmax      int          `url:"pmax,omitempty"`
	Origin    []ItemOrigin `url:"origin[],omitempty"`
	NotOrigin []ItemOrigin `url:"not_origin[],omitempty"`
	OrderBy   ItemOrder    `url:"order_by,omitempty"`
	SB        bool         `url:"sb,omitempty"`
	SBByMe    bool         `url:"sb_by_me,omitempty"`
	NSB       bool         `url:"nsb,omitempty"`
	NSBByMe   bool         `url:"nsb_by_me,omitempty"`
}

type AllViewedAccountsOptional struct {
	Page      int          `url:"page,omitempty"`
	Show      ItemStatus   `url:"show,omitempty"`
	Title     string       `url:"title,omitempty"`
	Pmin      int          `url:"pmin,omitempty"`
	Pmax      int          `url:"pmax,omitempty"`
	Origin    []ItemOrigin `url:"origin[],omitempty"`
	NotOrigin []ItemOrigin `url:"not_origin[],omitempty"`
	OrderBy   ItemOrder    `url:"order_by,omitempty"`
	SB        bool         `url:"sb,omitempty"`
	SBByMe    bool         `url:"sb_by_me,omitempty"`
	NSB       bool         `url:"nsb,omitempty"`
	NSBByMe   bool         `url:"nsb_by_me,omitempty"`
}

type AddProxyOptional struct {
	ProxyIP   string `url:"proxy_ip,omitempty"`
	ProxyPort int    `url:"proxy_port,omitempty"`
	ProxyUser string `url:"proxy_user,omitempty"`
	ProxyPass string `url:"proxy_pass,omitempty"`
	ProxyRow  string `url:"proxy_row,omitempty"`
}

type DeleteProxyOptional struct {
	ProxyID   int  `url:"proxy_id,omitempty"`
	DeleteAll bool `url:"delete_all,omitempty"`
}

type InvoiceListOptional struct {
	Page       int      `url:"page,omitempty"`
	Currency   Currency `url:"currency,omitempty"`
	Status     string   `url:"status,omitempty"`
	Amount     float64  `url:"amount,omitempty"`
	MerchantID int      `url:"merchant_id,omitempty"`
}

type InvoiceOptional struct {
	InvoiceId int    `url:"invoice_id,omitempty"`
	PaymentID string `url:"payment_id,omitempty"`
}

type CreateInvoiceOptional struct {
	Currency       Currency `url:"currency"`
	Amount         float64  `url:"amount"`
	PaymentID      string   `url:"payment_id"`
	Comment        string   `url:"comment"`
	URLSuccess     string   `url:"url_success"`
	URLCallback    string   `url:"url_callback,omitempty"`
	MerchantID     int      `url:"merchant_id"`
	Lifetime       int      `url:"lifetime,omitempty"`
	AdditionalData string   `url:"additional_data,omitempty"`
	IsTest         bool     `url:"is_test,omitempty"`
}

type EditMarketSettingsOptional struct {
	Currency                     Currency `url:"currency"`
	UserAllowAskDiscount         bool     `url:"user_allow_ask_discount,omitempty"`
	MaxDiscountPercent           int      `url:"max_discount_percent,omitempty"`
	DisableSteamGuard            bool     `url:"disable_steam_guard,omitempty"`
	DeauthorizeSteam             bool     `url:"deauthorize_steam,omitempty"`
	ChangePasswordOnPurchase     bool     `url:"change_password_on_purchase,omitempty"`
	HideFavourites               bool     `url:"hide_favourites,omitempty"`
	ShowTooLowPriceChangeWarning bool     `url:"show_too_low_price_change_warning,omitempty"`
	AllowAcceptAccounts          []string `url:"allow_accept_accounts,comma,omitempty"`
	MarketCustomTitle            string   `url:"market_custom_title,omitempty"`
	TelegramAPIID                string   `url:"telegram_api_id,omitempty"`
	TelegramAPIHash              string   `url:"telegram_api_hash,omitempty"`
	TelegramDeviceModel          string   `url:"telegram_device_model,omitempty"`
	TelegramSystemVersion        string   `url:"telegram_system_version,omitempty"`
	TelegramAppVersion           string   `url:"telegram_app_version,omitempty"`
	TelegramLangPack             string   `url:"telegram_lang_pack,omitempty"`
	TelegramLangCode             string   `url:"telegram_lang_code,omitempty"`
	TelegramSystemLangCode       string   `url:"telegram_system_lang_code,omitempty"`
}

type CreatePayoutOptional struct {
	PaymentSystem string                 `json:"payment_system"`
	Wallet        string                 `json:"wallet"`
	Amount        float64                `json:"amount"`
	Currency      Currency               `json:"currency"`
	IncludeFee    *bool                  `json:"include_fee,omitempty"`
	Extra         map[string]interface{} `json:"extra,omitempty"`
}

type TransferOptional struct {
	UserID           int        `url:"user_id,omitempty"`
	Username         string     `url:"username,omitempty"`
	Amount           int        `url:"amount"`
	Currency         Currency   `url:"currency"`
	Comment          string     `url:"comment,omitempty"`
	TransferHold     bool       `url:"transfer_hold,omitempty"`
	HoldLengthValue  int        `url:"hold_length_value,omitempty"`
	HoldLengthOption HoldPeriod `url:"hold_length_option,omitempty"`
}

type PaymentHistoryOptional struct {
	Type             OperationType `url:"type,omitempty"`
	Pmin             int           `url:"pmin,omitempty"`
	Pmax             int           `url:"pmax,omitempty"`
	Page             int           `url:"page,omitempty"`
	OperationIDLt    int           `url:"operation_id_lt,omitempty"`
	Receiver         string        `url:"receiver,omitempty"`
	Sender           string        `url:"sender,omitempty"`
	StartDate        time.Time     `url:"startDate,omitempty"`
	EndDate          time.Time     `url:"endDate,omitempty"`
	Wallet           string        `url:"wallet,omitempty"`
	Comment          string        `url:"comment,omitempty"`
	IsHold           bool          `url:"is_hold,omitempty"`
	ShowPaymentStats bool          `url:"show_payment_stats,omitempty"`
}

type CreateAutoPaymentsOptional struct {
	SecretAnswer     string   `url:"secret_answer,omitempty"`
	UsernameReceiver string   `url:"username_receiver"`
	Day              int      `url:"day"`
	Amount           float64  `url:"amount"`
	Currency         Currency `url:"currency,omitempty"`
	Description      string   `url:"description,omitempty"`
}

type FastSellOptional struct {
	Title             string     `url:"title,omitempty"`
	TitleEn           string     `url:"title_en,omitempty"`
	Price             float64    `url:"price"`
	CategoryID        CategoryID `url:"category_id"`
	Currency          Currency   `url:"currency"`
	ItemOrigin        ItemOrigin `url:"item_origin"`
	ExtendedGuarantee Guarantee  `url:"extended_guarantee,omitempty"`
	AllowAskDiscount  bool       `url:"allow_ask_discount,omitempty"`
	ProxyID           int        `url:"proxy_id,omitempty"`
	RandomProxy       bool       `url:"random_proxy,omitempty"`
}

type FastSellBodyOptional struct {
	Description       string      `json:"description,omitempty"`
	Information       string      `json:"information,omitempty"`
	Login             string      `json:"login,omitempty"`
	Password          string      `json:"password,omitempty"`
	LoginPassword     string      `json:"login_password,omitempty"`
	HasEmailLoginData bool        `json:"has_email_login_data,omitempty"`
	EmailLoginData    string      `json:"email_login_data,omitempty"`
	EmailType         string      `json:"email_type,omitempty"`
	Extra             interface{} `json:"extra,omitempty"`
}

type AddAccountOptional struct {
	Price      float64    `url:"price"`
	CategoryID CategoryID `url:"category_id"`
	Currency   Currency   `url:"currency"`
	ItemOrigin ItemOrigin `url:"item_origin"`

	Title             string    `url:"title,omitempty"`
	TitleEn           string    `url:"title_en,omitempty"`
	ExtendedGuarantee Guarantee `url:"extended_guarantee,omitempty"`
	ForceTempEmail    bool      `url:"forceTempEmail,omitempty"`
	ResellItemID      int       `url:"resell_item_id,omitempty"`
	HasEmailLoginData bool      `url:"has_email_login_data,omitempty"`
	EmailLoginData    string    `url:"email_login_data,omitempty"`
	EmailType         string    `url:"email_type,omitempty"`
	AllowAskDiscount  bool      `url:"allow_ask_discount,omitempty"`
	ProxyID           int       `url:"proxy_id,omitempty"`
	RandomProxy       bool      `url:"random_proxy,omitempty"`
	Description       string    `url:"description,omitempty"`
	Information       string    `url:"information,omitempty"`
}

type CheckAccountOptional struct {
	ResellItemID int  `url:"resell_item_id,omitempty"`
	RandomProxy  bool `url:"random_proxy,omitempty"`
}

type CheckAccountBodyOptional struct {
	Login             string                 `json:"login,omitempty"`
	Password          string                 `json:"password,omitempty"`
	LoginPassword     string                 `json:"login_password,omitempty"`
	HasEmailLoginData bool                   `json:"has_email_login_data,omitempty"`
	EmailLoginData    string                 `json:"email_login_data,omitempty"`
	EmailType         EmailType              `json:"email_type,omitempty"`
	Extra             map[string]interface{} `json:"extra,omitempty"`
}

type SteamInventoryValueOptional struct {
	AppID       AppID    `url:"app_id,omitempty"`
	Currency    Currency `url:"currency,omitempty"`
	IgnoreCache bool     `url:"ignore_cache,omitempty"`
}

type SteamValueOptional struct {
	Link        string   `url:"link,omitempty"`
	AppID       AppID    `url:"app_id,omitempty"`
	Currency    Currency `url:"currency,omitempty"`
	IgnoreCache bool     `url:"ignore_cache,omitempty"`
}

type UpdateSteamValueOptional struct {
	All       bool  `url:"all,omitempty"`
	AppID     AppID `url:"app_id,omitempty"`
	Authorize bool  `url:"authorize,omitempty"`
}

type ConfirmSDAOptional struct {
	ID    int `url:"id,omitempty"`
	Nonce int `url:"nonce,omitempty"`
}

type BulkGetAccountsOptional struct {
	ItemID           []int `json:"item_id,omitempty"`
	ParseSameItemIDs bool  `json:"parse_same_item_ids,omitempty"`
}

type EditAccountOptional struct {
	Title            string     `url:"title,omitempty"`
	TitleEn          string     `url:"title_en,omitempty"`
	Price            int        `url:"price,omitempty"`
	Currency         Currency   `url:"currency,omitempty"`
	ItemOrigin       ItemOrigin `url:"item_origin,omitempty"`
	EmailLoginData   string     `url:"email_login_data,omitempty"`
	EmailType        EmailType  `url:"email_type,omitempty"`
	AllowAskDiscount bool       `url:"allow_ask_discount,omitempty"`
	ProxyID          int        `url:"proxy_id,omitempty"`
}

type EditAccountBodyOptional struct {
	Description string `json:"description,omitempty"`
	Information string `json:"information,omitempty"`
}

type EmailLettersOptional struct {
	ItemID        int    `url:"item_id,omitempty"`
	EmailPassword string `url:"email_password,omitempty"`
	Email         string `url:"email,omitempty"`
	Password      string `url:"password,omitempty"`
	Limit         int    `url:"limit,omitempty"`
}

type EmailCodeOptional struct {
	ItemID int    `url:"item_id,omitempty"`
	Email  string `url:"email,omitempty"`
	Login  string `url:"login,omitempty"`
}

type BatchJob struct {
	Id     string        `json:"id,omitempty"`
	Uri    string        `json:"uri"`
	Method MethodRequest `json:"method,omitempty"`
	Params interface{}   `json:"params,omitempty"`
}
