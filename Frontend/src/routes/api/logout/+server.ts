import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ cookies }) => {
	const token = cookies.get('token');

	const response = await fetch('http://localhost:8080/api/logout', {
		method: 'POST',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		}
	});

	if (response.ok) {
		cookies.delete('token', { path: '/' });
		cookies.delete('refresh_token', { path: '/' });
		return json({ success: true });
	}
	return json({ success: false }, { status: 500 });
};
