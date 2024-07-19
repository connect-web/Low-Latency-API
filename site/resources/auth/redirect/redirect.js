document.addEventListener("DOMContentLoaded", async function() {
    // Check if the current path is /login or /register
    const currentPath = window.location.pathname;
    if (currentPath === '/login' || currentPath === '/register') {
        // Make an authenticated GET request to /api/v2/user/profile
        const data = await authenticatedGetRequest('/api/v2/user/profile');
        if (data === null){
            return;
        }

        // If the response is successful, redirect to /profile
        if ((currentPath === '/login' || currentPath === '/register') && data) {
            window.location.href = '/profile';
        }
        if ((currentPath === '/profile' || currentPath === '/search/skills' || currentPath === '/search/minigames' || currentPath === '/ml/skills' || currentPath === '/ml/minigames'  ) && !data) {
            window.location.href = '/login';
        }

    }
});

async function authenticatedGetRequest(url) {
    try {
        const response = await fetch(url, {
            method: 'GET',
            credentials: 'include'
        });

        if (response.status === 401) {
            window.location.href = '/login';
            return null;
        }

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        return await response.json();
    } catch (error) {
        console.log('Fetch error:', error);
        return null;
    }
}
