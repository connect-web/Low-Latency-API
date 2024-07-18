document.addEventListener('DOMContentLoaded', function() {
    const toplistApiURL = 'http://127.0.0.1:4050/api/v2/public/skill-toplist';
    const playersApiURL = 'http://127.0.0.1:4050/api/v2/public/skill-toplist-users?skill-id=';
    const rowsPerPage = 10; // Number of rows per page
    const maxPageButtons = 5; // Max number of page buttons to display
    let toplistData = [];
    let playersData = [];
    let currentPage = 1;
    let currentPlayersPage = 1;
    let totalPages = 0;
    let totalPlayersPages = 0;

    // Fetch data from the Toplist API
    async function fetchToplistData() {
        try {
            const response = await fetch(toplistApiURL);
            const jsonData = await response.json();
            toplistData = jsonData;
            totalPages = Math.ceil(toplistData.length / rowsPerPage);
            updateToplistTable();
            updateToplistPagination();
            updateTopNav();
        } catch (error) {
            console.error('Error fetching toplist data:', error);
        }
    }

    // Fetch data from the Players API
    async function fetchPlayersData(skillId) {
        try {
            const response = await fetch(playersApiURL + skillId);
            const jsonData = await response.json();
            playersData = jsonData;
            totalPlayersPages = Math.ceil(playersData.length / rowsPerPage);
            currentPlayersPage = 1; // Reset to first page on new data fetch
            updatePlayersTable();
            updatePlayersPagination();
        } catch (error) {
            console.error('Error fetching players data:', error);
        }
    }

    // Create table row for toplist
    function createToplistRow(item) {
        const row = document.createElement('tr');
        row.dataset.attrId = item.id;

        // Skills cell
        const skillsCell = document.createElement('td');
        skillsCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        const skillsContainer = document.createElement('div');
        skillsContainer.className = 'skills-container';
        item.skills.forEach(skill => {
            skillsContainer.innerHTML += `${getSkillImage(skill)}`;
        });
        skillsCell.appendChild(skillsContainer);
        row.appendChild(skillsCell);

        // Player count cell
        const countCell = document.createElement('td');
        countCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        countCell.textContent = item.count;
        row.appendChild(countCell);

        // View cell
        const viewCell = document.createElement('td');
        viewCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        const viewButton = document.createElement('button');
        viewButton.className = 'text-blue-500 hover:underline';
        viewButton.textContent = 'View';
        viewButton.addEventListener('click', () => {
            fetchPlayersData(item.id);
        });
        viewCell.appendChild(viewButton);
        row.appendChild(viewCell);

        return row;
    }

    // Create table row for players
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
        for (let skill in item.SkillLevels) {
            skillLevelsContainer.innerHTML += `${getSkillImage(skill)} ${viewLargeNumbers(item.SkillLevels[skill])}`;
        }
        skillLevelsCell.appendChild(skillLevelsContainer);
        row.appendChild(skillLevelsCell);

        // Minigames cell
        const minigamesCell = document.createElement('td');
        minigamesCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        const minigamesContainer = document.createElement('div');
        minigamesContainer.className = 'skills-container wrapped';
        for (let minigame in item.Minigames) {
            minigamesContainer.innerHTML += `${getMinigameImage(minigame)} ${viewLargeNumbers(item.Minigames[minigame])}`;
        }
        minigamesCell.appendChild(minigamesContainer);
        row.appendChild(minigamesCell);

        // Skill Gains cell
        const skillGainsCell = document.createElement('td');
        skillGainsCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        const skillGainsContainer = document.createElement('div');
        skillGainsContainer.className = 'skills-container wrapped';
        for (let skill in item.SkillGains) {
            skillGainsContainer.innerHTML += `${getSkillImage(skill)} ${viewLargeNumbers(item.SkillGains[skill])}`;
        }
        skillGainsCell.appendChild(skillGainsContainer);
        row.appendChild(skillGainsCell);

        // Minigame Gains cell
        const minigameGainsCell = document.createElement('td');
        minigameGainsCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        const minigameGainsContainer = document.createElement('div');
        minigameGainsContainer.className = 'skills-container';
        for (let minigame in item.MinigameGains) {
            minigameGainsContainer.innerHTML += `${getMinigameImage(minigame)} ${viewLargeNumbers(item.MinigameGains[minigame])}`;
        }
        minigameGainsCell.appendChild(minigameGainsContainer);
        row.appendChild(minigameGainsCell);

        return row;
    }

    // Update toplist table with rows for the current page
    function updateToplistTable() {
        const tableBody = document.querySelector('#toplist-table tbody');
        tableBody.innerHTML = ''; // Clear current table

        const start = (currentPage - 1) * rowsPerPage;
        const end = start + rowsPerPage;
        const pageData = toplistData.slice(start, end);

        pageData.forEach(item => {
            const row = createToplistRow(item);
            tableBody.appendChild(row);
        });
    }

    // Update players table with rows for the current page
    function updatePlayersTable() {
        const tableBody = document.querySelector('#players-table tbody');
        tableBody.innerHTML = ''; // Clear current table

        const start = (currentPlayersPage - 1) * rowsPerPage;
        const end = start + rowsPerPage;
        const pageData = playersData.slice(start, end);

        pageData.forEach(item => {
            const row = createPlayersRow(item);
            tableBody.appendChild(row);
        });
    }

    // Update toplist pagination buttons
    function updateToplistPagination() {
        const paginationContainer = document.querySelector('.flex.w-full.mt-5.space-x-2.justify-end');
        paginationContainer.innerHTML = ''; // Clear current pagination

        // Previous button
        const prevButton = document.createElement('button');
        prevButton.className = 'inline-flex items-center h-8 w-8 justify-center text-gray-400 rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none';
        prevButton.innerHTML = '<svg class="w-4" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>';
        prevButton.addEventListener('click', () => {
            if (currentPage > 1) {
                currentPage--;
                updateToplistTable();
                updateToplistPagination();
                updateTopNav();
            }
        });
        paginationContainer.appendChild(prevButton);

        // Page buttons
        const startPage = Math.max(1, currentPage - Math.floor(maxPageButtons / 2));
        const endPage = Math.min(totalPages, startPage + maxPageButtons - 1);

        for (let i = startPage; i <= endPage; i++) {
            const pageButton = document.createElement('button');
            pageButton.className = `inline-flex items-center h-8 w-8 justify-center rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none ${i === currentPage ? 'bg-gray-100 dark:bg-gray-800 dark:text-white' : 'text-gray-500'}`;
            pageButton.textContent = i;
            pageButton.addEventListener('click', () => {
                currentPage = i;
                updateToplistTable();
                updateToplistPagination();
                updateTopNav();
            });
            paginationContainer.appendChild(pageButton);
        }

        // Next button
        const nextButton = document.createElement('button');
        nextButton.className = 'inline-flex items-center h-8 w-8 justify-center text-gray-400 rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none';
        nextButton.innerHTML = '<svg class="w-4" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"></polyline></svg>';
        nextButton.addEventListener('click', () => {
            if (currentPage < totalPages) {
                currentPage++;
                updateToplistTable();
                updateToplistPagination();
                updateTopNav();
            }
        });
        paginationContainer.appendChild(nextButton);
    }

    // Update players pagination buttons
    function updatePlayersPagination() {
        const paginationContainer = document.querySelector('#players-pagination');
        paginationContainer.innerHTML = ''; // Clear current pagination

        // Previous button
        const prevButton = document.createElement('button');
        prevButton.className = 'inline-flex items-center h-8 w-8 justify-center text-gray-400 rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none';
        prevButton.innerHTML = '<svg class="w-4" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>';
        prevButton.addEventListener('click', () => {
            if (currentPlayersPage > 1) {
                currentPlayersPage--;
                updatePlayersTable();
                updatePlayersPagination();
            }
        });
        paginationContainer.appendChild(prevButton);

        // Page buttons
        const startPage = Math.max(1, currentPlayersPage - Math.floor(maxPageButtons / 2));
        const endPage = Math.min(totalPlayersPages, startPage + maxPageButtons - 1);

        for (let i = startPage; i <= endPage; i++) {
            const pageButton = document.createElement('button');
            pageButton.className = `inline-flex items-center h-8 w-8 justify-center rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none ${i === currentPlayersPage ? 'bg-gray-100 dark:bg-gray-800 dark:text-white' : 'text-gray-500'}`;
            pageButton.textContent = i;
            pageButton.addEventListener('click', () => {
                currentPlayersPage = i;
                updatePlayersTable();
                updatePlayersPagination();
            });
            paginationContainer.appendChild(pageButton);
        }

        // Next button
        const nextButton = document.createElement('button');
        nextButton.className = 'inline-flex items-center h-8 w-8 justify-center text-gray-400 rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none';
        nextButton.innerHTML = '<svg class="w-4" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"></polyline></svg>';
        nextButton.addEventListener('click', () => {
            if (currentPlayersPage < totalPlayersPages) {
                currentPlayersPage++;
                updatePlayersTable();
                updatePlayersPagination();
            }
        });
        paginationContainer.appendChild(nextButton);
    }

    // Update top navigation
    function updateTopNav() {
        const topNavSpan = document.querySelector('.pagination-top');
        topNavSpan.textContent = `Page ${currentPage} of ${totalPages}`;

        const topNavPrevButton = document.querySelector('.pagination-left');
        const topNavNextButton = document.querySelector('.pagination-right');

        // Remove existing event listeners to prevent multiple triggers
        topNavPrevButton.replaceWith(topNavPrevButton.cloneNode(true));
        topNavNextButton.replaceWith(topNavNextButton.cloneNode(true));

        const newTopNavPrevButton = document.querySelector('.pagination-left');
        const newTopNavNextButton = document.querySelector('.pagination-right');

        newTopNavPrevButton.addEventListener('click', () => {
            if (currentPage > 1) {
                currentPage--;
                updateToplistTable();
                updateToplistPagination();
                updateTopNav();
            }
        });

        newTopNavNextButton.addEventListener('click', () => {
            if (currentPage < totalPages) {
                currentPage++;
                updateToplistTable();
                updateToplistPagination();
                updateTopNav();
            }
        });
    }

    // Initial fetch and table setup
    fetchToplistData();
});
