import { redirect, fail } from '@sveltejs/kit';
import type { Actions } from './$types';
import type { GetChannelResponse } from '../../../../lib/myTypes';

export const load = async ({ params, cookies }: { params: { id: string }; cookies: any }) => {
	const { id } = params;
	const token = cookies.get('token');

	try {
		const response = await fetch(`http://localhost:8080/api/channels/${id}`, {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});
		if (!response.ok) {
			throw new Error('Failed to fetch channel data');
		}
		const data: GetChannelResponse = await response.json();

		return { data, error: null };
	} catch (error) {
		console.error(error);
		return { data: null, error: error };
	}
};

export const actions: Actions = {
	uploadVideo: async ({ request, cookies, fetch, params }) => {
		const token = cookies.get('token');
		const formData = await request.formData();
		const name = formData.get('name') as string;
		const description = formData.get('description') as string;
		const videoFile = formData.get('videoFile') as File;
		const channelId = params.id;
		if (!name || !description || !videoFile) {
			console.log('missing data');

			return fail(400, { missing: true });
		}

		const uploadData = new FormData();
		uploadData.append('name', name);
		uploadData.append('description', description);
		uploadData.append('videoFile', videoFile);
		uploadData.append('channelId', channelId);

		const res = await fetch(`http://localhost:8080/api/videos`, {
			method: 'POST',
			headers: {
				Authorization: `Bearer ${token}`
			},
			body: uploadData
		});
		if (!res.ok) {
			console.log('Upload failed');

			return fail(500, { serverError: true });
		}
		return { success: true };
	}
};
