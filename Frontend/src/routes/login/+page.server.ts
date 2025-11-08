import { redirect, fail } from '@sveltejs/kit';
import type { Actions } from './$types';

export const load = async ({ cookies, fetch }) => {
	const authToken = cookies.get('auth_token');
	const refreshToken = cookies.get('refresh_token');

	if (authToken || refreshToken) {
		throw redirect(302, '/home');
	}

	return {};
};

export const actions: Actions = {
	login: async ({ request }) => {
		const formData = await request.formData();
		const email = formData.get('email') as string;
		const password = formData.get('password') as string;

		if (!email || !password) {
			return fail(400, { email, missing: true });
		}

		const res = await fetch('http://localhost:8080/api/login', {
			method: 'POST',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ email, password })
		});

		if (!res.ok) {
			return fail(401, { email, incorrect: true });
		}
		throw redirect(302, '/register');
	}
};
