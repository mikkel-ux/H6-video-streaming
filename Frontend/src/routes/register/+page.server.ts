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
	register: async ({ request, cookies, fetch }) => {
		const formData = await request.formData();
		const firstName = formData.get('firstName') as string;
		const lastName = formData.get('lastName') as string;
		const email = formData.get('email') as string;
		const password = formData.get('password') as string;
		const userName = formData.get('userName') as string;
		const age = formData.get('age') as string;
		const channelName = formData.get('channelName') as string;
		const channelDescription = formData.get('channelDescription') as string;
		if (!email || !password || !userName || !age || !channelName || !channelDescription) {
			return fail(400, { email, missing: true });
		}

		const res = await fetch('http://localhost:8080/api/users', {
			method: 'POST',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				firstName,
				lastName,
				email,
				password,
				userName,
				age: Number(age),
				channelName,
				channelDescription
			})
		});

		if (!res.ok) {
			if (res.status === 409) {
				return fail(409, { email, exists: true });
			}

			return fail(500, { email, serverError: true });
		}
		throw redirect(302, '/login');
	}
};
