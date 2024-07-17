document.addEventListener('DOMContentLoaded', function() {
    const apiURL = 'http://127.0.0.1:4050/api/v2/public/skill-toplist';
    const rowsPerPage = 10; // Number of rows per page
    const maxPageButtons = 5; // Max number of page buttons to display
    let data = [];
    let currentPage = 1;
    let totalPages = 0;

    // Fetch data from the API
    async function fetchData() {
        try {
            const response = await fetch(apiURL);
            const jsonData = await response.json();
            data = jsonData;
            totalPages = Math.ceil(data.length / rowsPerPage);
            updateTable();
            updatePagination();
            updateTopNav();
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    }

    // Create table row
    function createTableRow(item) {
        const row = document.createElement('tr');
        row.dataset.attrId = item.id;

        // Skills cell
        const skillsCell = document.createElement('td');
        skillsCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        const skillsContainer = document.createElement('div');
        skillsContainer.className = 'skills-container';
        item.skills.forEach(skill => {
            const img = document.createElement('img');
            img.src = `./resources/imgs/hiscore/skill_icons/${skill.toLowerCase()}.png`;
            img.alt = `${skill} Icon`;
            img.className = 'img-fluid inline-block h-6 w-6 mr-2 mb-2';
            skillsContainer.appendChild(img);
        });
        skillsCell.appendChild(skillsContainer);
        row.appendChild(skillsCell);

        // Player count cell
        const countCell = document.createElement('td');
        countCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        countCell.textContent = item.count;
        row.appendChild(countCell);

        // View cell (additional cell if required)
        const viewCell = document.createElement('td');
        viewCell.className = 'font-normal px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800';
        viewCell.textContent = 'View'; // Replace with actual content if needed
        row.appendChild(viewCell);

        return row;
    }

    // Update table with rows for the current page
    function updateTable() {
        const tableBody = document.querySelector('tbody');
        tableBody.innerHTML = ''; // Clear current table

        const start = (currentPage - 1) * rowsPerPage;
        const end = start + rowsPerPage;
        const pageData = data.slice(start, end);

        pageData.forEach(item => {
            const row = createTableRow(item);
            tableBody.appendChild(row);
        });
    }

    // Update pagination buttons
    function updatePagination() {
        const paginationContainer = document.querySelector('.flex.w-full.mt-5.space-x-2.justify-end');
        paginationContainer.innerHTML = ''; // Clear current pagination

        // Previous button
        const prevButton = document.createElement('button');
        prevButton.className = 'inline-flex items-center h-8 w-8 justify-center text-gray-400 rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none';
        prevButton.innerHTML = '<svg class="w-4" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>';
        prevButton.addEventListener('click', () => {
            if (currentPage > 1) {
                currentPage--;
                updateTable();
                updatePagination();
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
                updateTable();
                updatePagination();
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
                updateTable();
                updatePagination();
                updateTopNav();
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
                updateTable();
                updatePagination();
                updateTopNav();
            }
        });

        newTopNavNextButton.addEventListener('click', () => {
            if (currentPage < totalPages) {
                currentPage++;
                updateTable();
                updatePagination();
                updateTopNav();
            }
        });
    }

    // Initial fetch and table setup
    fetchData();
});
