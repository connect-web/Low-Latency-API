
document.addEventListener('DOMContentLoaded', async function() {
    const data = await authenticatedGetRequest('/api/v2/user/global-stats');

    if (data) {
        updateBanCountData(data);
    }
});

function viewLargeNumbersProfilePage(value) {
    if (value > 1_000_000_000_000) {
        return `${(value / 1_000_000_000_000).toFixed(2)}T`;
    } else if (value > 1_000_000_000) {
        return `${(value / 1_000_000_000).toFixed(2)}B`;
    } else if (value > 10_000_000) {
        return `${(value / 1_000_000).toFixed(2)}M`;
    } else if (value > 100_000) {
        return `${(value / 1_000).toFixed(2)}K`;
    }
    return value.toString();
}

function updateBanCountData(data) {
    const totalBansElements = document.querySelectorAll('.TotalBans');
    if (totalBansElements.length > 0) {
        totalBansElements.forEach(element => {
            element.textContent = viewLargeNumbersProfilePage(data.Bans);
        });
    }

    const totalExperienceElement = document.querySelector('.TotalExperience');
    if (totalExperienceElement) {
        totalExperienceElement.textContent = viewLargeNumbersProfilePage(data.TotalExperience);
    }

    const suspiciousUsersElement = document.querySelector('.SuspiciousUsers');
    if (suspiciousUsersElement) {
        suspiciousUsersElement.textContent = viewLargeNumbersProfilePage(data.SuspiciousUsers);
    }

    const lastUpdatedElement = document.querySelector('.Last_updated');
    if (lastUpdatedElement) {
        const lastUpdatedString = data.Last_updated;
        const lastUpdatedDate = new Date(lastUpdatedString);
        const now = new Date();
        const diffInMs = now - lastUpdatedDate;
        const diffInHours = Math.max(Math.round(diffInMs / (1000 * 60 * 60)), 1); // Calculate hours and ensure a minimum of 1 hour

        lastUpdatedElement.textContent = `${diffInHours} hour${diffInHours > 1 ? 's' : ''} ago`;
    }




    const skillsContainer = document.querySelector('#skillsContainer .card-list');
    if (skillsContainer) {
        skillsContainer.innerHTML = ''; // Clear previous content
        // Convert the skills object to an array of [skill, amount] pairs and sort it by amount
        const sortedSkills = Object.entries(data.Skills).sort((a, b) => b[1] - a[1]);

        for (const [skill, amount] of sortedSkills) {
            const skillItem = document.createElement('div');
            skillItem.classList.add('card-item-skills');
            skillItem.innerHTML = `
                <div class="card lb">
                    ${getSkillImage(skill)}
                    <div class="title">${skill}</div>
                    <div class="value SuspiciousUsers">${viewLargeNumbersProfilePage(amount)}</div>
                </div>
            `;
            skillsContainer.appendChild(skillItem);
        }
    }

    // Update minigames
    const minigamesContainer = document.querySelector('#minigamesContainer .card-list');
    if (minigamesContainer) {
        minigamesContainer.innerHTML = ''; // Clear previous content

        // Convert the minigames object to an array of [minigame, amount] pairs and sort it by amount
        const sortedMinigames = Object.entries(data.Minigames).sort((a, b) => b[1] - a[1]);

        // Create and append minigame items in the sorted order
        for (const [minigame, amount] of sortedMinigames) {
            const minigameItem = document.createElement('div');
            minigameItem.classList.add('card-item-minigames');
            minigameItem.innerHTML = `
                <div class="card lb">
                    ${getMinigameImage(minigame)}
                    <div class="title">${minigame}</div>
                    <div class="value SuspiciousUsers">${viewLargeNumbersProfilePage(amount)}</div>
                </div>
            `;
            minigamesContainer.appendChild(minigameItem);
        }
    }
}

