const container = document.getElementById('imageGrid');
const totalImages = 24; // Total number of images
const skills = [
    'attack',
    'strength',
    'defence',
    'ranged',
    'magic',
    'prayer',
    'hitpoints',
    'cooking',
    'woodcutting',
    'fletching',
    'fishing',
    'firemaking',
    'crafting',
    'smithing',
    'mining',
    'herblore',
    'agility',
    'thieving',
    'slayer',
    'farming',
    'runecraft',
    'hunter',
    'construction',
]



function getSkillPath(skill){
    return `./resources/imgs/hiscore/skill_icons/${skill}.png`
}

function titleCase(str) {
    return str.toLowerCase().replace(/(^|\s)\S/g, (letter) => letter.toUpperCase());
}

// Example usage:


function setupTableListeners() {
    const tableBody = document.getElementById('skillsTable');
    tableBody.removeEventListener("click", toggleTick); // Remove listener to prevent duplicates if setup is called multiple times
    tableBody.addEventListener("click", toggleTick);
}

function toggleTick(event) {
    const target = event.target.closest('.tick-box'); // Ensures that clicks on the tick-box toggle the class
    if (target) {
        target.classList.toggle('ticked');
    }
}

// Function to create and populate the table
function populateTable() {
    const tableBody = document.getElementById('skillsTable');
    skills.forEach(skill => {
        const row = tableBody.insertRow();
        const iconCell = row.insertCell(0);
        const nameCell = row.insertCell(1);
        const dailyXpCell = row.insertCell(2);
        const minLevelCell = row.insertCell(3);
        const maxLevelCell = row.insertCell(4);
        const tickCell = row.insertCell(5);

        const img = document.createElement('img');
        img.src = getSkillPath(skill);
        img.className = 'img-fluid';
        img.alt = `${skill} Icon`;
        iconCell.appendChild(img);

        nameCell.innerText = titleCase(skill);

        dailyXpCell.innerHTML = `<input type='text' class='text-input' value='0' />`;
        minLevelCell.innerHTML = `<input type='text' class='text-input' value='1' />`;
        maxLevelCell.innerHTML = `<input type='text' class='text-input' value='126' />`;
        tickCell.innerHTML = `
        <div class="thirdTick" style="margin:auto;">
    <div class="tick-box" >
      <svg class="tick-mark" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 40 40"><path fill="none" d="M8 20.2l7.1 7.2 16.7-16.8"/></svg>
    </div>
  </div>`;
    });

    let thirdTicks = document.getElementsByClassName("thirdTick");

// Convert HTMLCollection to an array to use forEach
    Array.from(thirdTicks).forEach(ticker => {
        ticker.addEventListener("click", () => {
            // This needs to be specific to each ticker, not a global state
            let isTicked = ticker.classList.contains("ticked");
            if (!isTicked) {
                ticker.classList.add("ticked");
            } else {
                ticker.classList.remove("ticked");
            }
        });
    });

}

populateTable();



/*

Hide / Show filters table
 */

document.addEventListener('DOMContentLoaded', () => {
    const selectButtons = document.querySelectorAll('.button-filter-toggle');
    const filterTable = document.getElementById('filterTable')

    selectButtons.forEach(button => {
        if (button && filterTable) {

            button.addEventListener('click', function() {
                filterTable.classList.toggle('hidden');
            });
        }
    })
});