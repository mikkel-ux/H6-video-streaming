<script lang="ts">
    let { children } = $props();

    async function logout() {
        const res = await fetch('/api/logout', {
            method: 'POST',
            credentials: 'include'
        })
        if (res.ok) {
            window.location.href = '/login';
        } else {
            console.error('Logout failed');
        }
    }
</script>

<nav class="p-4 text-white w-full">
    <button onclick={logout}>Logout</button>
</nav>
{@render children()}