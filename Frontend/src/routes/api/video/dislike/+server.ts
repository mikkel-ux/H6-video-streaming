import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request, cookies }) => {
	const { videoId } = await request.json();
	const token = cookies.get('token');
	const response = await fetch(`http://localhost:8080/api/videos/${videoId}/dislike`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		}
	});
	if (response.ok) {
		return json({ success: true, error: null });
	} else {
		return json({ success: false, error: 'Failed to dislike video' }, { status: response.status });
	}
};
