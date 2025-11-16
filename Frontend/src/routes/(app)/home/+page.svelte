<script lang="ts">
	import { onMount } from 'svelte';
    import type { PageData } from './$types';

    type VideoPreview = {
        id: number;
        title: string;
        thumbnail: string;
    }
    
    /* export let data: PageData; */
    let { data }: { data: PageData } = $props();

    let videos = $state<VideoPreview[]>(data.videos);
</script>

<h1 class="text-3xl font-bold underline">
    Welcome to the Home Page!
</h1>

<div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4 mt-4">
    {#each videos as video}
        <div class="border rounded-lg overflow-hidden shadow-lg">
            <img src={`http://localhost:8080/api/images/${video.thumbnail}`} alt={video.title} class="w-full h-48 object-cover" />
            <div class="p-4">
                <h2 class="text-xl font-semibold">{video.title}</h2>
            </div>
        </div>
    {/each}
</div>