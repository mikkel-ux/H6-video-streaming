import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get('token');
	try {
		const response = await fetch('http://localhost:8080/api/videos/random30', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		});

		if (!response.ok) {
			throw new Error(`Failed to fetch videos: ${response.statusText}`);
		}

		const data = await response.json();
		return { videos: data };
	} catch (error) {
		console.error('Error fetching videos:', error);
		return { videos: [] };
	}
};
