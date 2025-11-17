<script lang="ts">
	import { onMount } from 'svelte';
    import type { PageData } from './$types';
	import { on } from 'svelte/events';

    type VideoPreview = {
        videoId: number;
        title: string;
        thumbnail: string;
    }
    
    /* export let data: PageData; */
    let { data }: { data: PageData } = $props();

    let videos = $state<VideoPreview[]>(data.videos);

    onMount(() => {
        console.log(videos);
    });
</script>

<h1 class="text-3xl font-bold underline text-center mt-6">
    Welcome to the Home Page!
</h1>

<div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4 mt-4 p-4">
    {#each videos as video}
        <a href={`/video/${video.videoId}`} class="block group text-center">
            <div class="w-full h-48 bg-gray-800 rounded-lg overflow-hidden mb-4">
                <img src={`http://localhost:8080/api/images/${video.thumbnail}`} alt={video.title} class="w-full h-full object-cover transform group-hover:scale-105 transition-transform duration-300" />
            </div>
            <h2 class="text-lg font-semibold text-white group-hover:text-yellow-300 transition-colors duration-300">{video.title}</h2>
        </a>
    {/each}
</div>