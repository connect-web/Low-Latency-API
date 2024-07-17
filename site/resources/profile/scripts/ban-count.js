
document.addEventListener('DOMContentLoaded', async function() {
    const data = await authenticatedGetRequest('/api/v2/user/ban-count');

    if (data) {
        updateBanCountData(data);
    }
});

function updateBanCountData(data) {
    document.querySelector('.TotalBans').textContent = data.TotalBans;
}