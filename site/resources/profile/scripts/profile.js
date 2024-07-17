
document.addEventListener('DOMContentLoaded', async function() {
    const data = await authenticatedGetRequest('/api/v2/user/profile');

    if (data) {
        updateProfileData(data);
    }
});

function updateProfileData(profile) {
    // Update botsTracked
    const botsTrackedElement = document.querySelector('.value.botsTracked');
    if (botsTrackedElement) {
        botsTrackedElement.textContent = profile.BotsTracked;
    }

    // Update botsBanned
    const botsBannedElement = document.querySelector('.value.botsBanned');
    if (botsBannedElement) {
        botsBannedElement.textContent = profile.BotsBanned;
    }

    // Update BannedExperience
    const bannedExperienceElement = document.querySelector('.value.BannedExperience');
    if (bannedExperienceElement) {
        bannedExperienceElement.textContent = profile.BannedExperience;
    }

    // Update playersAdded
    const playersAddedElement = document.querySelector('.value.playersAdded');
    if (playersAddedElement) {
        playersAddedElement.textContent = profile.PlayersAdded;
    }

    // Update username in all elements with the class "Username"
    const usernameElements = document.querySelectorAll('.Username');
    if (usernameElements.length > 0) {
        usernameElements.forEach(element => {
            element.textContent = profile.Username;
        });
    }
}
