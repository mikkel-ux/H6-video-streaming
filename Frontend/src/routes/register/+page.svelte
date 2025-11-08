<script lang="ts">
    import { enhance } from '$app/forms';
    import type { PageProps } from './$types';
    let { data, form }: PageProps = $props();
</script>

<div class="flex flex-col items-center justify-center min-h-screen">
    <h1 class="mb-6 text-3xl font-bold">Register</h1>
    <form method="POST" action="?/register" class="w-full max-w-md space-y-6">
        {#if form?.missing}
            <p class="text-red-500">Please fill in all fields.</p>
        {/if}
        {#if form?.exists}
            <p class="text-red-500">Email already exists.</p>
        {/if}
        {#if form?.serverError}
            <p class="text-red-500">Something went wrong. Please try again</p>
        {/if}
        
        <div class="grid grid-cols-1 gap-6 sm:grid-cols-2">
            <div class="space-y-2">
                <label for="firstName" class="block text-sm font-medium">First name:</label>
                <input 
                    type="text" 
                    id="firstName" 
                    name="firstName" 
                    class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
            </div>

            <div class="space-y-2">
                <label for="lastName" class="block text-sm font-medium">Last name:</label>
                <input 
                    type="text" 
                    id="lastName" 
                    name="lastName" 
                    class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
            </div>
        </div>

        <div class="space-y-2">
            <label for="userName" class="block text-sm font-medium">Username:</label>
            <input 
                type="text" 
                id="userName" 
                name="userName" 
                class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
        </div>

        <div class="space-y-2">
            <label for="age" class="block text-sm font-medium">Age:</label>
            <input 
                type="number" 
                id="age" 
                name="age" 
                min="15"
                class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
        </div>

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

        <div class="space-y-2">
            <label for="channelName" class="block text-sm font-medium">Channel name:</label>
            <input 
                type="text" 
                id="channelName" 
                name="channelName" 
                class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
        </div>

        <div class="space-y-2">
            <label for="channelDescription" class="block text-sm font-medium">Bio</label>
            <textarea 
                id="channelDescription" 
                name="channelDescription" 
                rows="3"
                class="w-full px-3 py-2 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            ></textarea>
        </div>

        <button 
            type="submit" 
            class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
            Register
        </button>
    </form>
    <p class="mt-4 text-sm">
        Already have an account? 
        <a href="/login" class="font-medium text-blue-600 hover:text-blue-500">
            Login
        </a>
    </p>
</div>

