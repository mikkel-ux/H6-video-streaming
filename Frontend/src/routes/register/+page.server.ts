import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, fetch }) => {
	/* const authToken = cookies.get('auth_token');
	const refreshToken = cookies.get('refresh_token'); */
	console.log('register');

	/* if (authToken || refreshToken) {
		throw redirect(302, '/home');
	} */

	return {};
};
