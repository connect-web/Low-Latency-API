document.addEventListener('DOMContentLoaded', function() {
    const toplistApiURL = '/api/v2/public/boss-minigame-toplist';
    let toplistData = [];
    let currentPage = 1;
    let totalPages = 0;

    // Fetch data from the Toplist API
    async function fetchToplistData() {
        try {
            const response = await fetch(toplistApiURL);
            const jsonData = await response.json();
            toplistData = jsonData;
            totalPages = Math.ceil(toplistData.length / 10); // Assuming 10 rows per page for pagination
            updateToplistTable();
            updateToplistPagination();
        } catch (error) {
            console.error('Error fetching toplist data:', error);
        }
    }

    // Create table row for toplist
    function createToplistRow(item) {
        const row = document.createElement('tr');
        row.dataset.attrId = item.minigame;

        // Minigame name cell
        const minigameCell = document.createElement('td');
        minigameCell.className = 'px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800 minigame-toplist-image';
        minigameCell.innerHTML = `${getMinigameImage(item.minigame)}`;;
        row.appendChild(minigameCell);

        // Count cell
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
            skillDetails = {
                skills: [item.minigame],
                count: item.count,
            }
            fetchPlayersData(item.minigame);
        });
        viewCell.appendChild(viewButton);
        row.appendChild(viewCell);

        return row;
    }

    // Update toplist table with rows for the current page
    function updateToplistTable() {
        const tableBody = document.querySelector('#toplist-table tbody');
        tableBody.innerHTML = '';

        const start = (currentPage - 1) * 10;
        const end = Math.min(start + 10, toplistData.length);
        const pageData = toplistData.slice(start, end);

        pageData.forEach(item => {
            tableBody.appendChild(createToplistRow(item));
        });
    }

    function updateToplistPagination() {
        const paginationContainer = document.querySelector('.flex.w-full.mt-5.space-x-2.justify-end');
        paginationContainer.innerHTML = '';

        const maxPageButtons = 5;
        const halfPageVisible = Math.floor(maxPageButtons / 2);

        const startPage = Math.max(1, currentPage - halfPageVisible);
        const endPage = Math.min(startPage + maxPageButtons - 1, totalPages);

        // Previous button
        const prevButton = document.createElement('button');
        prevButton.className = 'inline-flex items-center h-8 w-8 justify-center text-gray-400 rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none';
        prevButton.innerHTML = '<svg class="w-4" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>';
        prevButton.disabled = currentPage === 1;

        prevButton.addEventListener('click', () => {
            currentPage = Math.max(1, currentPage - 1);
            updateToplistTable();
            updateToplistPagination();
            updateTopNav();
        });
        paginationContainer.appendChild(prevButton);

        for (let i = startPage; i <= endPage; i++) {
            const pageButton = document.createElement('button');
            pageButton.textContent = i;
            pageButton.className = `inline-flex items-center h-8 w-8 justify-center rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none ${i === currentPage ? 'bg-gray-100 dark:bg-gray-800 dark:text-white' : 'text-gray-500'}`;
            pageButton.onclick = () => {
                currentPage = i;
                updateToplistTable();
                updateToplistPagination();
                updateTopNav();
            };
            paginationContainer.appendChild(pageButton);
        }

        // Next button
        const nextButton = document.createElement('button');
        nextButton.className = 'inline-flex items-center h-8 w-8 justify-center text-gray-400 rounded-md shadow border border-gray-200 dark:border-gray-800 leading-none';
        nextButton.innerHTML = '<svg class="w-4" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"></polyline></svg>';
        nextButton.disabled = currentPage === totalPages;
        nextButton.addEventListener('click', () => {
            currentPage = Math.min(totalPages, currentPage + 1);
            updateToplistTable();
            updateToplistPagination();
            updateTopNav();
        });
        paginationContainer.appendChild(nextButton);
    }

    function updateTopNav() {
        const topNavSpan = document.querySelector('.Toplist .pagination-top');
        topNavSpan.textContent = `Page ${currentPage} of ${totalPages}`;

        const topNavPrevButton = document.querySelector('.Toplist .pagination-left');
        const topNavNextButton = document.querySelector('.Toplist .pagination-right');

        // Remove existing event listeners to prevent multiple triggers
        topNavPrevButton.replaceWith(topNavPrevButton.cloneNode(true));
        topNavNextButton.replaceWith(topNavNextButton.cloneNode(true));

        const newTopNavPrevButton = document.querySelector('.Toplist .pagination-left');
        const newTopNavNextButton = document.querySelector('.Toplist .pagination-right');


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

    fetchToplistData();
});
