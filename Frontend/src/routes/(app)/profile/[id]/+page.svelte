<script lang="ts">
    import { enhance } from '$app/forms';
    import type { PageProps } from './$types';
    type userProfile = {
        userId: string;
        userName: string;
        email: string;
        age?: number;
        channelId: string;
    }

    type profile = {
        user: userProfile;
        userId: string;
    }

    let { data, form }: { data: profile, form: PageProps } = $props();
    let changePasswordModalOpen = $state<boolean>(false);
</script>

<div class="p-4 max-w-md mx-auto bg-white rounded-xl shadow-md space-y-4 text-black">
    <h1 class="text-2xl font-bold mb-4 text-center">Profile Page</h1>
    <p class="mb-2">User ID: {data.user.userId}</p>
    <p class="mb-2">Username: {data.user.userName}</p>
    <p class="mb-2">Email: {data.user.email}</p>
    {#if data.user.age}
    <p class="mb-2">Age: {data.user.age}</p>
    {/if}
    <a href={`/channel/${data.user.channelId}`} class="text-blue-500 hover:underline">go to channel</a>
    <br>
    <br>
    <button onclick={() => changePasswordModalOpen = true}
        class="bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >change password</button>
</div>


{#if changePasswordModalOpen}
    <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-gray-900 p-6 rounded-lg w-96">
            <h2 class="text-2xl mb-4">Change Password</h2>
            <form method="POST" action="?/changePassword" use:enhance>
                <label for="currentPassword">Current Password:</label>
                <input type="password" name="currentPassword" class="w-full mb-4 p-2 border border-gray-300 rounded text-black"/>
                <label for="newPassword">New Password:</label>
                <input type="password" name="newPassword" class="w-full mb-4 p-2 border border-gray-300 rounded text-black"/>
                <button type="submit" class="bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Change Password</button>
            </form>
            <button class="mt-4" onclick={() => changePasswordModalOpen = false}>Close</button>
        </div>
    </div>
{/if}
