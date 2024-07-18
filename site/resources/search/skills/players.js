let playersData = [];
let filteredPlayersData = [];
let currentPlayersPage = 1;
let totalPlayersPages = 0;
let skillDetails = { skills: [] }; // Initialize with an empty skills array

// Default filter values
let combatMin = 0, combatMax = 126;
let totalLevelMin = 0, totalLevelMax = 2500;
let totalExpMin = 0, totalExpMax = 10000000000;


async function fetchPlayersData(skillId) {
    try {
        const response = await fetch(`/api/v2/public/skill-toplist-users?skill-id=${skillId}`);
        const jsonData = await response.json();
        playersData = jsonData;

        if (skillDetails.skills.length !== 0){
            sortPlayersDataInit(skillDetails.skills[0], false)
        }


        // Apply initial filters
        setupFilters();
        applyFilters();
        updateSortControls(skillDetails);


    } catch (error) {
        console.error('Error fetching players data:', error);
    }
}

function updateTable() {
    const tableBody = document.querySelector('#players-table tbody');
    tableBody.innerHTML = ''; // Clear existing rows
    const start = (currentPlayersPage - 1) * 10;
    const end = start + 10;
    const pageData = filteredPlayersData.slice(start, end); // Use filteredPlayersData
    pageData.forEach(item => {
        const row = createPlayersRow(item);
        tableBody.appendChild(row);
    });
}

// Sorts the players data and logs the top 3 and bottom 3 entries
function sortPlayersDataInit(skill, ascending) {
    // Ensure the data array is not empty and the skill is specified
    if (!playersData.length || !skill) {
        console.log("No data to sort or skill not specified");
        return;
    }

    // Sorting the players data based on skill gains
    playersData.sort((a, b) => {
        const valueA = (a.SkillGains && a.SkillGains[skill]) || 0; // Ensuring the SkillGains object exists
        const valueB = (b.SkillGains && b.SkillGains[skill]) || 0;
        return ascending ? valueA - valueB : valueB - valueA;
    });

    // Printing the sorted data to check the top 3 and bottom 3
    const top3 = playersData.slice(0, 3);
    const bottom3 = playersData.slice(-3);

    console.log(`Top 3 players for skill ${skill}:`);
    top3.forEach(player => console.log(`${player.Username || 'Unnamed'}: ${player.SkillGains[skill] || 0}`));

    console.log(`Bottom 3 players for skill ${skill}:`);
    bottom3.forEach(player => console.log(`${player.Username || 'Unnamed'}: ${player.SkillGains[skill] || 0}`));

}

// Sorts the players data and logs the top 3 and bottom 3 entries
function sortPlayersData(skill, ascending) {
    // Ensure the data array is not empty and the skill is specified
    if (!filteredPlayersData.length || !skill) {
        console.log("No data to sort or skill not specified");
        return;
    }

    // Sorting the players data based on skill gains
    filteredPlayersData.sort((a, b) => {
        const valueA = (a.SkillGains && a.SkillGains[skill]) || 0; // Ensuring the SkillGains object exists
        const valueB = (b.SkillGains && b.SkillGains[skill]) || 0;
        return ascending ? valueA - valueB : valueB - valueA;
    });

    // Printing the sorted data to check the top 3 and bottom 3
    const top3 = filteredPlayersData.slice(0, 3);
    const bottom3 = filteredPlayersData.slice(-3);

    console.log(`Top 3 players for skill ${skill}:`);
    top3.forEach(player => console.log(`${player.Username || 'Unnamed'}: ${player.SkillGains[skill] || 0}`));

    console.log(`Bottom 3 players for skill ${skill}:`);
    bottom3.forEach(player => console.log(`${player.Username || 'Unnamed'}: ${player.SkillGains[skill] || 0}`));
    updatePaginationControls();  // Ensure this is called to reflect changes
    updateTable();
}





function updateSortControls(skillDetails) {
    const sortControls = document.getElementById('sort-controls');
    sortControls.innerHTML = ''; // Clear previous controls

    // Create and style the dropdown for skills
    const sortDropdown = document.createElement('select');
    sortDropdown.className = 'dropdown-select rounded-md';
    skillDetails.skills.forEach(skill => {
        const option = document.createElement('option');
        option.value = skill;
        option.textContent = skill; // Enhancement to include skill icons could be added here
        sortDropdown.appendChild(option);
    });


    // Create the sort direction toggle button using the styled format
    const sortDirectionButton = document.createElement('button');
    sortDirectionButton.className = 'inline-flex mr-3 items-center h-8 pl-2.5 pr-2 rounded-md shadow text-gray-700 dark:text-gray-400 dark:border-gray-800 border border-gray-200 leading-none py-0';
    let isAscending = false; // Default sort order is now highest to lowest



    // run on the init

    // Set the initial arrow icon to indicate descending order
    sortDirectionButton.innerHTML = `
        <svg viewBox="0 0 24 24" class="w-4 mr-2 text-gray-400 dark:text-gray-600" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="6 15 12 9 18 15"></polyline> <!-- Arrow points down for descending -->
        </svg>
        Sort Direction
    `;

    sortDirectionButton.onclick = () => {
        isAscending = !isAscending;
        sortDirectionButton.querySelector('svg polyline').setAttribute('points', isAscending ? '6 9 12 15 18 9' : '6 15 12 9 18 15');
        currentPlayersPage = 1; // Reset to the first page
        sortPlayersData(sortDropdown.value, isAscending);
        updateTable(); // Add this line
        updatePaginationControls(); // Add this line
    };

    // Event listener for sorting based on skill selection
    sortDropdown.addEventListener('change', () => {
        currentPlayersPage = 1; // Reset to the first page
        sortPlayersData(sortDropdown.value, isAscending);
        updateTable(); // Add this line
        updatePaginationControls(); // Add this line
    });

    // Append controls to sort-controls div
    sortControls.appendChild(sortDropdown);
    sortControls.appendChild(sortDirectionButton);
}

function setupFilters() {
    const combatMinInput = document.getElementById('combat-min');
    const combatMaxInput = document.getElementById('combat-max');
    const totalLevelMinInput = document.getElementById('total-level-min');
    const totalLevelMaxInput = document.getElementById('total-level-max');
    const totalExpMinInput = document.getElementById('total-exp-min');
    const totalExpMaxInput = document.getElementById('total-exp-max');

    combatMinInput.addEventListener('change', () => {
        combatMin = combatMinInput.value !== '' ? parseInt(combatMinInput.value, 10) : 0;
        applyFilters();
    });

    combatMaxInput.addEventListener('change', () => {
        combatMax = combatMaxInput.value !== '' ? parseInt(combatMaxInput.value, 10) : 126;
        applyFilters();
    });

    totalLevelMinInput.addEventListener('change', () => {
        totalLevelMin = totalLevelMinInput.value !== '' ? parseInt(totalLevelMinInput.value, 10) : 0;
        applyFilters();
    });

    totalLevelMaxInput.addEventListener('change', () => {
        totalLevelMax = totalLevelMaxInput.value !== '' ? parseInt(totalLevelMaxInput.value, 10) : 2500;
        applyFilters();
    });

    totalExpMinInput.addEventListener('change', () => {
        totalExpMin = totalExpMinInput.value !== '' ? parseInt(totalExpMinInput.value, 10) : 0;
        applyFilters();
    });

    totalExpMaxInput.addEventListener('change', () => {
        totalExpMax = totalExpMaxInput.value !== '' ? parseInt(totalExpMaxInput.value, 10) : 10000000000;
        applyFilters();
    });
}

function applyFilters() {
    filteredPlayersData = playersData.filter(player =>
        player.CombatLevel >= combatMin && player.CombatLevel <= combatMax &&
        player.TotalLevel >= totalLevelMin && player.TotalLevel <= totalLevelMax &&
        player.TotalExperience >= totalExpMin && player.TotalExperience <= totalExpMax
    );
    totalPlayersPages = Math.ceil(filteredPlayersData.length / 10);
    currentPlayersPage = 1; // Reset to the first page
    updateTable();
    updatePaginationControls();
}



// NAVIGATION

function updatePaginationControls() {
    const pageInfos = document.querySelectorAll('.PlayersTable .player-nav-buttons .pagination-top');
    const prevButtons = document.querySelectorAll('.PlayersTable .player-nav-buttons .pagination-left');
    const nextButtons = document.querySelectorAll('.PlayersTable .player-nav-buttons .pagination-right');


    pageInfos.forEach(pageInfo => {
        pageInfo.textContent = `Page ${currentPlayersPage} of ${totalPlayersPages}`;
    });

    prevButtons.forEach(prevButton => {
        prevButton.disabled = currentPlayersPage === 1;
        prevButton.onclick = () => {
            if (currentPlayersPage > 1) {
                currentPlayersPage--;
                updateTable();
                updatePaginationControls();
            }
        };
    });

    nextButtons.forEach(nextButton => {
        nextButton.disabled = currentPlayersPage === totalPlayersPages;
        nextButton.onclick = () => {
            if (currentPlayersPage < totalPlayersPages) {
                currentPlayersPage++;
                updateTable();
                updatePaginationControls();
            }
        };
    });
}


// create row
function createPlayersRow(item) {
    const row = document.createElement('tr');

    // Username cell
    const usernameCell = document.createElement('td');
    usernameCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    usernameCell.textContent = item.Username;
    row.appendChild(usernameCell);

    // Combat Level cell
    const combatLevelCell = document.createElement('td');
    combatLevelCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    combatLevelCell.textContent = item.CombatLevel;
    row.appendChild(combatLevelCell);

    // Total Level cell
    const totalLevelCell = document.createElement('td');
    totalLevelCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    totalLevelCell.textContent = item.TotalLevel;
    row.appendChild(totalLevelCell);

    // Total Experience cell
    const totalExperienceCell = document.createElement('td');
    totalExperienceCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    totalExperienceCell.textContent = viewLargeNumbers(item.TotalExperience);
    row.appendChild(totalExperienceCell);

    // Skill Levels cell
    const skillLevelsCell = document.createElement('td');
    skillLevelsCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    const skillLevelsContainer = document.createElement('div');
    skillLevelsContainer.className = 'skills-container wrapped';
    const sortedSkillLevels = Object.entries(item.SkillLevels).sort((a, b) => b[1] - a[1]);
    sortedSkillLevels.forEach(([skill, value]) => {
        const skillItem = document.createElement('div');
        skillItem.className = 'skill-item';
        skillItem.innerHTML = `${getSkillImage(skill)} ${viewLargeNumbers(value)}`;
        skillLevelsContainer.appendChild(skillItem);
    });
    skillLevelsCell.appendChild(skillLevelsContainer);
    row.appendChild(skillLevelsCell);

    // Minigames cell
    const minigamesCell = document.createElement('td');
    minigamesCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    const minigamesContainer = document.createElement('div');
    minigamesContainer.className = 'skills-container wrapped';
    const sortedMinigames = Object.entries(item.Minigames).sort((a, b) => b[1] - a[1]);
    sortedMinigames.forEach(([minigame, value]) => {
        const minigameItem = document.createElement('div');
        minigameItem.className = 'skill-item';
        minigameItem.innerHTML = `${getMinigameImage(minigame)} ${viewLargeNumbers(value)}`;
        minigamesContainer.appendChild(minigameItem);
    });
    minigamesCell.appendChild(minigamesContainer);
    row.appendChild(minigamesCell);

    // Skill Gains cell
    const skillGainsCell = document.createElement('td');
    skillGainsCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    const skillGainsContainer = document.createElement('div');
    skillGainsContainer.className = 'skills-container wrapped';
    const sortedSkillGains = Object.entries(item.SkillGains).sort((a, b) => b[1] - a[1]);
    sortedSkillGains.forEach(([skill, value]) => {
        const skillGainItem = document.createElement('div');
        skillGainItem.className = 'skill-item';
        skillGainItem.innerHTML = `${getSkillImage(skill)} ${viewLargeNumbers(value)}`;
        skillGainsContainer.appendChild(skillGainItem);
    });
    skillGainsCell.appendChild(skillGainsContainer);
    row.appendChild(skillGainsCell);

    // Minigame Gains cell
    const minigameGainsCell = document.createElement('td');
    minigameGainsCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
    const minigameGainsContainer = document.createElement('div');
    minigameGainsContainer.className = 'skills-container';
    const sortedMinigameGains = Object.entries(item.MinigameGains).sort((a, b) => b[1] - a[1]);
    sortedMinigameGains.forEach(([minigame, value]) => {
        const minigameGainItem = document.createElement('div');
        minigameGainItem.className = 'skill-item';
        minigameGainItem.innerHTML = `${getMinigameImage(minigame)} ${viewLargeNumbers(value)}`;
        minigameGainsContainer.appendChild(minigameGainItem);
    });
    minigameGainsCell.appendChild(minigameGainsContainer);
    row.appendChild(minigameGainsCell);

    return row;
}