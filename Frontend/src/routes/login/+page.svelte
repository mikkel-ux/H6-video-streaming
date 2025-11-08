<script lang="ts">
    import { enhance } from '$app/forms';
    import type { PageProps } from './$types';
    let errorMsg = $state<string>("");
    let { data, form }: PageProps = $props();
</script>
<div class="flex flex-col items-center justify-center min-h-screen">

    <h1>login</h1>
        <form method="POST" action="?/login" use:enhance
         class="flex flex-col gap-4">
         {#if form?.missing}<p class="text-red-500">Please fill in all fields.</p>{/if}
            {#if form?.incorrect}<p class="text-red-500">Invalid email or password.</p>{/if}

            <label for="email">Email:</label>
            <input type="email" id="email" name="email" class="text-black" />

            <label for="password">Password:</label>
            <input type="password" id="password" name="password" class="text-black" min="8" />

            <button type="submit" class="bg-blue-500 text-white py-2 px-4 rounded">
                Login
            </button>
            {#if errorMsg}
                <p class="text-red-500">{errorMsg}</p>
            {/if}
        </form>
        <p class="text-sm">don't have an account? <a href="/register">Register</a></p>
    </div>