<script lang="ts">
	import { onMount } from 'svelte';
    import type { PageData } from './$types';
    
    export let data: PageData;

    onMount(() => {
        console.log('Home page mounted');
    });
</script>

<h1 class="text-3xl font-bold underline">
    Welcome to the Home Page!
</h1>

<pre>{JSON.stringify(data, null, 2)}</pre>