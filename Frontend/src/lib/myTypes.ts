export type VideoPreview = {
	videoId: string;
	title: string;
	thumbnail: string;
};

export type UserSummary = {
	userId: string;
	userName: string;
};

export type GetChannelResponse = {
	channelId: string;
	name: string;
	description: string;
	user: UserSummary;
	videos: VideoPreview[];
	isOwner: boolean;
};
