import { redirect, fail } from '@sveltejs/kit';
import type { Actions } from './$types';

export const load = async ({ params, cookies }: { params: { id: string }; cookies: any }) => {
	const { id } = params;
	const token = cookies.get('token');

	try {
		const response = await fetch(`http://localhost:8080/api/users/${id}`, {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});
		if (!response.ok) {
			throw new Error('Failed to fetch user data');
		}
		const user = await response.json();
		console.log(user);

		return { user };
	} catch (error) {
		console.error(error);
		return { user: null };
	}
};

export const actions: Actions = {
	changePassword: async ({ request, cookies, params }) => {
		const formData = await request.formData();
		const currentPassword = formData.get('currentPassword') as string;
		const newPassword = formData.get('newPassword') as string;
		const token = cookies.get('token');
		const { id } = params;
		if (!currentPassword || !newPassword) {
			return fail(400, { missing: true });
		}
		const response = await fetch(`http://localhost:8080/api/users/${id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			},
			body: JSON.stringify({ currentPassword, newPassword })
		});
		const data = await response.json();
		console.log(data);
		if (!response.ok) {
			console.log('Password change failed');

			return fail(400, { incorrect: true });
		}
		return { success: true };
	}
};
