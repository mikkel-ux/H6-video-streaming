<script lang="ts">
	import { page } from '$app/state';
	import { onMount } from 'svelte';
    import type { PageData } from './$types';

    let { data }: { data: PageData } = $props();
	if (!data.data) {
		throw new Error('No video data available');
	}
	let urlParams = new URLSearchParams(page.url.search);
	let startTime = parseInt(urlParams.get('t') || '0', 10);

	let videoEl = $state<HTMLVideoElement | null>(null);
	let videoLikes = $state<number>(data.data.likes ?? 0);
	let videoDislikes = $state<number>(data.data.dislikes ?? 0);
	let videoIsLiked = $state<boolean>(data.data.isLiked ?? false);
	let videoIsDisliked = $state<boolean>(data.data.isDisliked ?? false);

	onMount(() => {
		setStartTime();
	});

	async function like() {
		if (!data.data) {
			return;
		}
		try {
			const response = await fetch(`/api/video/like`, {
				method: 'POST',
				credentials: 'include',
				body: JSON.stringify({ videoId: data.data.videoId })
			});
			if (response.ok) {
				if (videoIsLiked) {
					videoLikes -= 1;
					videoIsLiked = false;
				} else{
					videoLikes += 1;
					videoIsLiked = true;
				}
			} else {
				console.error('Failed to like the video');
			}
		} catch (error) {
			console.error('Error liking the video:', error);
		}		
	}

	async function dislike() {
		if (!data.data) {
			return;
		}
		try {
			const response = await fetch(`/api/video/dislike`, {
				method: 'POST',
				credentials: 'include',
				body: JSON.stringify({ videoId: data.data.videoId })
			});
			if (response.ok) {
				if (videoIsDisliked) {
					videoDislikes -= 1;
					videoIsDisliked = false;
				} else{
					videoDislikes += 1;
					videoIsDisliked = true;
				}
			} else {
				console.error('Failed to dislike the video');
			}
		} catch (error) {
			console.error('Error disliking the video:', error);
		}		
	}

	function setStartTime() {
		if (videoEl) {
			videoEl.currentTime = startTime;
		}
	}

</script>

<div class="max-w-6xl mx-auto px-4 py-8">
	<h1 class="text-3xl md:text-4xl font-bold text-white mb-2 tracking-tight">{data.data.title}</h1>

	
	<div class="flex items-center space-x-3 mb-6 align-middle justify-between">
        <div>

            <span class="text-sm text-gray-400">Uploaded by</span>
            <span class="font-medium text-yellow-300 hover:underline cursor-pointer">
                <a href={`/channel/${data.data.channel?.channelId}`}>{data.data.channel?.name || "Unknown"}</a>
            </span>
            <span class="text-gray-500 text-sm">•</span>
            <span class="text-sm text-gray-400">
                {new Date(data.data.uploaded).toLocaleDateString()}
            </span>
        </div>

        {#if data.data.likes !== undefined || data.data.dislikes !== undefined}
		    <div class="flex items-center gap-4 mt-6 text-sm text-gray-400">
			    <!-- <span onclick={like}>👍 {videolikes}</span>
			    <span onclick={dislike}>👎 {videodislikes}</span> -->
					<button onclick={like} class="hover:text-white cursor-pointer {videoIsLiked ? 'text-green-500' : ''}">👍 {videoLikes}</button>
					<button onclick={dislike} class="hover:text-white cursor-pointer {videoIsDisliked ? 'text-red-500' : ''}">👎 {videoDislikes}</button>
		    </div>
	    {/if}
	</div>

	<div class="relative rounded-xl overflow-hidden shadow-2xl bg-black mb-6">
		<video
			bind:this={videoEl}
			controls
			autoplay
			class="w-full aspect-video"
			poster={`http://localhost:8080/api/images/${data.data.thumbnail}`}
			onloadedmetadata={setStartTime}
		>
			<source src={`http://localhost:8080/api/videos/stream/${data.data.url}`} type="video/mp4" />
			<track
				kind="captions"
				label="English Captions"
				srcLang="en"
				src={`http://localhost:8080/api/videos/stream/${data.data.url}.vtt`}
				default
			/>
			Your browser does not support the video tag.
		</video>
	</div>

	<div class="mt-8 p-6 bg-gray-900/50 backdrop-blur rounded-lg border border-gray-700">
		<h2 class="text-xl font-semibold text-white mb-3">Description</h2>
		<p class="text-gray-300 whitespace-pre-wrap leading-relaxed">
			{data.data.description || "No description provided."}
		</p>
	</div>
</div>