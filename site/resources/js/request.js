document.addEventListener('DOMContentLoaded', function () {
    const findUsersButton = document.querySelector('.button-findUsers');
    findUsersButton.addEventListener('click', function () {
        const tableRows = document.querySelectorAll('#skillsTable tr');
        const params = new URLSearchParams();

        tableRows.forEach(row => {
            const skillName = row.cells[1].innerText.trim().toLowerCase(); // Assuming the skill name is in the second column
            const isActive = row.querySelector('.thirdTick').classList.contains('ticked');

            if (isActive) {
                const dailyXP = row.cells[2].querySelector('input').value;
                const minLevel = row.cells[3].querySelector('input').value;
                const maxLevel = row.cells[4].querySelector('input').value;

                params.append(`${skillName}`, 'true');
                params.append(`${skillName}_daily`, dailyXP);
                params.append(`${skillName}_min_lvl`, minLevel);
                params.append(`${skillName}_max_lvl`, maxLevel);
            }
        });

        fetch(`/api/find-skill-bots?${params.toString()}`)
            .then(response => response.json())
            .then(data => tableFiller(data))
            .catch(error => console.error('Error:', error));
    });
});

example_user = {
    "Username": "Mobilerr",
    "SkillGains": {"Cooking": 124315, "Overall": 124315},
    "SkillLevels": {
        "Agility": 85,
        "Attack": 99,
        "Construction": 99,
        "Cooking": 125,
        "Crafting": 99,
        "Defence": 96,
        "Farming": 92,
        "Firemaking": 99,
        "Fishing": 99,
        "Fletching": 99,
        "Herblore": 99,
        "Hitpoints": 103,
        "Hunter": 86,
        "Magic": 90,
        "Mining": 95,
        "Prayer": 99,
        "Ranged": 104,
        "Runecraft": 79,
        "Slayer": 93,
        "Smithing": 99,
        "Strength": 99,
        "Thieving": 86,
        "Woodcutting": 100
    },
    "SkillRatios": null,
}

function titleCase(str) {
    return str.toLowerCase().replace(/(^|\s)\S/g, (letter) => letter.toUpperCase());
}
function viewLargeNumbers(value) {
    if (value > 10_000_000) {
        return `${(value / 1_000_000).toFixed(2)}M`;
    } else if (value > 100_000) {
        return `${(value / 1_000).toFixed(2)}K`;
    }
    return value.toString();
}

function getSkillImage(skill) {
    skill = skill.toLowerCase();
    return `<img src="./resources/imgs/hiscore/skill_icons/${skill}.png" alt="Icon ${skill}" class="img-fluid">`;
}
function tableFiller(users) {
    const tableBody = document.getElementById('tableBody');
    // Clear existing table rows
    while (tableBody.firstChild) {
        tableBody.removeChild(tableBody.firstChild);
    }

    var skillsOrder = [
        "overall", "attack", "defence", "strength", "hitpoints",
        "ranged", "prayer", "magic", "cooking", "woodcutting",
        "fletching", "fishing", "firemaking", "crafting",
        "smithing", "mining", "herblore", "agility", "thieving",
        "slayer", "farming", "runecraft", "hunter", "construction"
    ];

    skillsOrder = skillsOrder.map(skill => titleCase(skill));

    users.forEach(user => {
        const row = tableBody.insertRow();

        // Insert Username
        const nameCell = row.insertCell();
        nameCell.textContent = user.Username;

        // Insert XP/Day information
        const xpDayCell = row.insertCell();
        let xpContent = document.createElement('div');
        xpContent.className = 'skill-container col-xp-day';

        skillsOrder.forEach(skill => {
            if (user.SkillGains[skill] !== undefined) {
                let skillInfo = document.createElement('div');
                skillInfo.className = 'skillTable';
                skillInfo.innerHTML = `${getSkillImage(skill)}<p>${viewLargeNumbers(user.SkillGains[skill])}</p>`;
                xpContent.appendChild(skillInfo);
            }
        });
        xpDayCell.appendChild(xpContent);

        // Insert Skill Levels
        const skillLevelCell = row.insertCell();
        let levelContent = document.createElement('div');
        levelContent.className = 'skill-container';

        skillsOrder.forEach(skill => {
            if (user.SkillLevels[skill] !== undefined) {
                let skillInfo = document.createElement('div');
                skillInfo.className = 'skillTable';
                skillInfo.innerHTML = `${getSkillImage(skill)}<p>${user.SkillLevels[skill]}</p>`;
                levelContent.appendChild(skillInfo);
            }
        });
        skillLevelCell.appendChild(levelContent);
    });
}


// Example usage
const users = [{
    "Username": "Mobilerr",
    "SkillGains": { "Cooking": 124315, "Overall": 124315 },
    "SkillLevels": {
        "Agility": 85,
        "Attack": 99,
        "Construction": 99,
        "Cooking": 125,
        "Crafting": 99,
        "Defence": 96,
        "Farming": 92,
        "Firemaking": 99,
        "Fishing": 99,
        "Fletching": 99,
        "Herblore": 99,
        "Hitpoints": 103,
        "Hunter": 86,
        "Magic": 90,
        "Mining": 95,
        "Prayer": 99,
        "Ranged": 104,
        "Runecraft": 79,
        "Slayer": 93,
        "Smithing": 99,
        "Strength": 99,
        "Thieving": 86,
        "Woodcutting": 100
    }
}];

// tableFiller(users);
