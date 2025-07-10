package dto

type Category string

const (
	CategorySteam      Category = "steam"
	CategoryFortnite   Category = "fortnite"
	CategoryMiHoYo     Category = "mihoyo"
	CategoryRiot       Category = "riot"
	CategoryTelegram   Category = "telegram"
	CategorySupercell  Category = "supercell"
	CategoryOrigin     Category = "origin"
	CategoryWoT        Category = "world-of-tanks"
	CategoryWoTBlitz   Category = "wot-blitz"
	CategoryEpicGames  Category = "epicgames"
	CategoryGifts      Category = "gifts"
	CategoryEFT        Category = "escape-from-tarkov"
	CategorySocialClub Category = "socialclub"
	CategoryUplay      Category = "uplay"
	CategoryWarThunder Category = "war-thunder"
	CategoryDiscord    Category = "discord"
	CategoryTikTok     Category = "tiktok"
	CategoryInstagram  Category = "instagram"
	CategoryBattleNet  Category = "battlenet"
	CategoryChatGPT    Category = "chatgpt"
	CategoryVPN        Category = "vpn"
	CategoryRoblox     Category = "roblox"
	CategoryWarface    Category = "warface"
	CategoryMinecraft  Category = "minecraft"
	CategoryEmpty      Category = ""
)

type CategoryID int

const (
	CategoryIDSteam      CategoryID = 1
	CategoryIDOrigin     CategoryID = 3
	CategoryIDWarface    CategoryID = 4
	CategoryIDUplay      CategoryID = 5
	CategoryIDChatGPT    CategoryID = 6
	CategoryIDSocialClub CategoryID = 7
	CategoryIDFortnite   CategoryID = 9
	CategoryIDInstagram  CategoryID = 10
	CategoryIDBattleNet  CategoryID = 11
	CategoryIDEpicGames  CategoryID = 12
	CategoryIDRiot       CategoryID = 13
	CategoryIDWoT        CategoryID = 14
	CategoryIDSupercell  CategoryID = 15
	CategoryIDWoTBlitz   CategoryID = 16
	CategoryIDMiHoYo     CategoryID = 17
	CategoryIDEFT        CategoryID = 18
	CategoryIDVPN        CategoryID = 19
	CategoryIDTikTok     CategoryID = 20
	CategoryIDGifts      CategoryID = 30
	CategoryIDMinecraft  CategoryID = 28
	CategoryIDRoblox     CategoryID = 31
)

type OperationType string

const (
	OpIncome            OperationType = "income"
	OpCost              OperationType = "cost"
	OpPaidItem          OperationType = "paid_item"
	OpSoldItem          OperationType = "sold_item"
	OpWithdrawalBalance OperationType = "withdrawal_balance"
	OpRefilledBalance   OperationType = "refilled_balance"
	OpInternalPurchase  OperationType = "internal_purchase"
	OpMoneyTransfer     OperationType = "money_transfer"
	OpClaimHold         OperationType = "claim_hold"
	OpInsuranceDeposit  OperationType = "insurance_deposit"
	OpPaidMail          OperationType = "paid_mail"
	OpContest           OperationType = "contest"
)

type HoldPeriod string

const (
	HoldHour  HoldPeriod = "hour"
	HoldDay   HoldPeriod = "day"
	HoldWeek  HoldPeriod = "week"
	HoldMonth HoldPeriod = "month"
)

type Currency string

const (
	CurrencyRUB Currency = "rub"
	CurrencyUAH Currency = "uah"
	CurrencyKZT Currency = "kzt"
	CurrencyBYN Currency = "byn"
	CurrencyUSD Currency = "usd"
	CurrencyEUR Currency = "eur"
	CurrencyGBP Currency = "gbp"
	CurrencyCNY Currency = "cny"
	CurrencyTRY Currency = "try"
	CurrencyJPY Currency = "jpy"
	CurrencyBRL Currency = "brl"
)

type ItemOrigin string

const (
	OriginBrute            ItemOrigin = "brute"
	OriginStealer          ItemOrigin = "stealer"
	OriginPhishing         ItemOrigin = "phishing"
	OriginAutoreg          ItemOrigin = "autoreg"
	OriginPersonal         ItemOrigin = "personal"
	OriginResale           ItemOrigin = "resale"
	OriginDummy            ItemOrigin = "dummy"
	OriginSelfRegistration ItemOrigin = "self_registration"
)

type Guarantee int

const (
	GuaranteeHalfDay   Guarantee = -1
	GuaranteeOneDay    Guarantee = 0
	GuaranteeThreeDays Guarantee = 1
)

type ItemStatus string

const (
	StatusActive          ItemStatus = "active"
	StatusPaid            ItemStatus = "paid"
	StatusDeleted         ItemStatus = "deleted"
	StatusAwaiting        ItemStatus = "awaiting"
	StatusClosed          ItemStatus = "closed"
	StatusDiscountRequest ItemStatus = "discount_request"
	StatusStickied        ItemStatus = "stickied"
	StatusPreActive       ItemStatus = "pre_active"
)

type ItemOrder string

const (
	OrderPriceAsc        ItemOrder = "price_to_up"
	OrderPriceDesc       ItemOrder = "price_to_down"
	OrderPostedDateDesc  ItemOrder = "pdate_to_down"
	OrderPostedDateAsc   ItemOrder = "pdate_to_up"
	OrderUploadDateDesc  ItemOrder = "pdate_to_down_upload"
	OrderUploadDateAsc   ItemOrder = "pdate_to_up_upload"
	OrderDeletedDateAsc  ItemOrder = "ddate_to_down"
	OrderDeletedDateDesc ItemOrder = "ddate_to_up"
	OrderEditedDateAsc   ItemOrder = "edate_to_down"
	OrderEditedDateDesc  ItemOrder = "edate_to_up"
)

type Extra string

const (
	ExtraProxy            Extra = "proxy"
	ExtraCloseItem        Extra = "close_item"
	ExtraRegion           Extra = "region"
	ExtraService          Extra = "service"
	ExtraSystem           Extra = "system"
	ExtraConfirmationCode Extra = "confirmationCode"
	ExtraCookies          Extra = "cookies"
	ExtraLoginNoCookies   Extra = "login_without_cookies"
	ExtraCookieLogin      Extra = "cookie_login"
	ExtraMFAFile          Extra = "mfa_file"
	ExtraDota2MMR         Extra = "dota2_mmr"
	ExtraEAGames          Extra = "ea_games"
	ExtraUplayGames       Extra = "uplay_games"
	ExtraTheQuarry        Extra = "the_quarry"
	ExtraWarframe         Extra = "warframe"
	ExtraARK              Extra = "ark"
	ExtraARKAscended      Extra = "ark_ascended"
	ExtraGenshinCurrency  Extra = "genshin_currency"
	ExtraHonkaiCurrency   Extra = "honkai_currency"
	ExtraZenlessCurrency  Extra = "zenless_currency"
	ExtraTelegramClient   Extra = "telegramClient"
	ExtraTelegramJSON     Extra = "telegramJson"
	ExtraCheckChannels    Extra = "checkChannels"
	ExtraCheckSpam        Extra = "checkSpam"
	ExtraCheckHypixelBan  Extra = "checkHypixelBan"
)

type AppID int

const (
	AppIDCS2      AppID = 730
	AppIDPUBG     AppID = 578080
	AppIDSteam    AppID = 753
	AppIDDota     AppID = 570
	AppIDTF2      AppID = 440
	AppIDRust     AppID = 252490
	AppIDUnturned AppID = 304930
	AppIDKF2      AppID = 232090
	AppIDDST      AppID = 322330
)

type EmailType string

const (
	EmailTypeNative  EmailType = "native"
	EmailTypeAutoreg EmailType = "autoreg"
)

type TypeHtml string

const (
	TypeProfiles TypeHtml = "profiles"
	TypeGames    TypeHtml = "games"
)

const (
	AccountImageTypeSkins    = "skins"
	AccountImageTypePickaxes = "pickaxes"
	AccountImageTypeDances   = "dances"
	AccountImageTypeGliders  = "gliders"
	AccountImageTypeWeapons  = "weapons"
	AccountImageTypeAgents   = "agents"
	AccountImageTypeBuddies  = "buddies"
)

type MethodRequest string

const (
	MethodGet    MethodRequest = "GET"
	MethodPost   MethodRequest = "POST"
	MethodPut    MethodRequest = "PUT"
	MethodDelete MethodRequest = "DELETE"
)
