package dto

func IsValidCategory(c Category) bool {
	switch c {
	case
		CategoryEmpty,
		CategorySteam,
		CategoryFortnite,
		CategoryMiHoYo,
		CategoryRiot,
		CategoryTelegram,
		CategorySupercell,
		CategoryOrigin,
		CategoryWoT,
		CategoryWoTBlitz,
		CategoryEpicGames,
		CategoryGifts,
		CategoryEFT,
		CategorySocialClub,
		CategoryUplay,
		CategoryWarThunder,
		CategoryDiscord,
		CategoryTikTok,
		CategoryInstagram,
		CategoryBattleNet,
		CategoryChatGPT,
		CategoryVPN,
		CategoryRoblox,
		CategoryWarface,
		CategoryMinecraft:
		return true
	default:
		return false
	}
}
