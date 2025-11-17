<script lang="ts">
    import type { GetChannelResponse } from '$lib/myTypes';
    import type { PageData, ActionData } from './$types';

    let { data, form }: { data: PageData, form: ActionData } = $props();
    const channelData = $state<GetChannelResponse | null>(data?.data);
    let uploadModalOpen = $state<boolean>(false);
</script>

<div class="text-white p-6 md:p-10">
	{#if !channelData}
		<div class="flex justify-center items-center h-[80vh]">
			<p class="text-xl font-medium animate-pulse">Loading channel...</p>
		</div>
	{:else}
		<header class="max-w-4xl mx-auto text-center mb-10">
			<h1 class="text-4xl md:text-5xl">
				{channelData.name}
			</h1>
			<p class="mt-4 text-lg max-w-2xl mx-auto">{channelData.description}</p>
		</header>
        {#if channelData.isOwner}
            <div class="max-w-4xl mx-auto text-center mb-8">
                <button onclick={() => uploadModalOpen = true} class="bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">upload video</button>
            </div>
        {/if}

		<main class="max-w-6xl mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            {#each channelData.videos as video}
                <a href={`/video/${video.videoId}`} class="block group text-center">
                    <div class="w-full h-48 bg-gray-800 rounded-lg overflow-hidden mb-4">
                        <img src={`http://localhost:8080/api/images/${video.thumbnail}`} alt={video.title} class="w-full h-full object-cover transform group-hover:scale-105 transition-transform duration-300" />
                    </div>
                    <h2 class="text-lg font-semibold text-white group-hover:text-yellow-300 transition-colors duration-300">{video.title}</h2>
                </a>
            {/each}
		</main>
	{/if}

    {#if uploadModalOpen}
        <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <div class="bg-gray-900 p-6 rounded-lg w-96">
                <h2 class="text-2xl mb-4">Upload Video</h2>
                <form method="POST" action="?/uploadVideo" enctype="multipart/form-data">
                    <label for="name">name:</label>
                    <input type="text" name="name" class="w-full mb-4 p-2 text-black"/>
                    <label for="description">description:</label>
                    <textarea name="description" class="w-full mb-4 p-2 text-black"></textarea>
                    <label for="videoFile">video file:</label>
                    <input type="file" name="videoFile" class="w-full mb-4 p-2 border border-gray-300 rounded"/>
                    <button type="submit" class="bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Upload</button>
                </form>
                <button class="mt-4" onclick={() => uploadModalOpen = false}>Close</button>
            </div>
        </div>
    {/if}
</div>