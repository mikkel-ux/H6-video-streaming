import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, fetch }) => {
	const authToken = cookies.get('auth_token');
	const refreshToken = cookies.get('refresh_token');
	console.log('something');

	if (!authToken || !refreshToken) {
		throw redirect(302, '/login');
	}

	return {};
};
