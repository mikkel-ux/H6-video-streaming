<script lang="ts">
    import { enhance } from '$app/forms';
    import type { PageProps } from './$types';
    let { data, form }: PageProps = $props();
</script>
<div class="flex flex-col items-center justify-center min-h-screen">
    <h1 class="mb-6 text-3xl font-bold">Login</h1>
    <form method="POST" action="?/login" use:enhance class="w-full max-w-md space-y-4">
        {#if form?.missing}
            <p class="text-red-500">Please fill in all fields.</p>
        {/if}
        {#if form?.incorrect}
            <p class="text-red-500">Invalid email or password.</p>
        {/if}

        <div class="space-y-2">
            <label for="email" class="block text-sm font-medium">Email:</label>
            <input 
                type="email" 
                id="email" 
                name="email" 
                class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
        </div>

        <div class="space-y-2">
            <label for="password" class="block text-sm font-medium">Password:</label>
            <input 
                type="password" 
                id="password" 
                name="password" 
                minlength="8"
                class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
        </div>

        <button 
            type="submit" 
            class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
            Login
        </button>
    </form>
    <p class="mt-4 text-sm">
        Don't have an account? 
        <a href="/register" class="font-medium text-blue-600 hover:text-blue-500">
            Register
        </a>
    </p>
</div>