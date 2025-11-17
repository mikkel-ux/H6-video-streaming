<script lang="ts">
    import type { PageData } from './$types';
    import { onMount } from 'svelte';

    let { data }: { data: PageData } = $props();
    onMount(() => {
        console.log(data.data);
    });
</script>

<div class="p-4">
    <h1 class="text-2xl font-bold mb-4">{data.data.title}</h1>
    <video controls autoplay class="w-full max-w-3xl mb-4">
			<source src={`http://localhost:8080/api/videos/stream/${data.data.url}`} type="video/mp4" />
			<track
                kind="captions"
                label="English captions"
                srcLang="en"
                src={`http://localhost:8080/api/videos/stream/${data.data.url}.vtt`}
                default
			/>
		</video>
    <p class="mb-2">{data.data.description}</p>
</div>