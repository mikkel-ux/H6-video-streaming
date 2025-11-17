package dto

type GetChannelResponse struct {
	ChannelID   int64          `json:"channelId"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	User        UserSummary    `json:"user"`
	Videos      []VideoPreview `json:"videos"`
	IsOwner     bool           `json:"isOwner"`
}

type ChannelSummary struct {
	ChannelID int64  `json:"channelId"`
	Name      string `json:"name"`
}
