/* test */
type channelSummary = {
	channelId: string;
	name: string;
};

type VideoData = {
	videoId: string;
	title: string;
	description: string;
	url: string;
	thumbnail: string;
	uploaded: string;
	likes: number;
	dislikes: number;
	channel: channelSummary;
	isLiked: boolean;
	isDisliked: boolean;
};

export const load = async ({ params, cookies }: { params: { id: string }; cookies: any }) => {
	const { id } = params;
	const token = cookies.get('token');

	try {
		const response = await fetch(`http://localhost:8080/api/videos/${id}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		});

		if (!response.ok) {
			throw new Error(`Failed to fetch video: ${response.statusText}`);
		}

		const data: VideoData = await response.json();
		return { data: data, error: null };
	} catch (error) {
		console.error('Error fetching video:', error);
		return { data: null, error: error };
	}
};
