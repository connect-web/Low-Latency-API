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


skills.forEach(skill =>{
    const colDiv = document.createElement('div');
    colDiv.className = 'skillContainer pt-1'; // mb-4 for margin bottom

    const img = document.createElement('img');

    img.src = `./resources/imgs/hiscore/skill_icons/${skill}.png`;
    img.alt = `Icon ${skill}`;
    img.className = 'img-fluid skill'; // Bootstrap class for responsive images



    const controlContainer = document.createElement('div');
    controlContainer.className = 'control-container';

    const plusIcon = document.createElement('i')
    plusIcon.className = 'fa fa-plus fa-selectable fa-selected-plus'
    plusIcon.ariaHidden = 'true'

    const minusIcon = document.createElement('i')
    minusIcon.className = 'fa fa-minus fa-selectable fa-selected-minus'
    minusIcon.ariaHidden = 'true'

    const input = document.createElement('input');
    input.placeholder = "99";
    input.className = 'textInput';
    input.style.display = 'block'; // Change to 'block' to make it visible by default
    input.id = skill;

    controlContainer.appendChild(minusIcon)
    controlContainer.appendChild(plusIcon)

    colDiv.appendChild(img);
    colDiv.appendChild(controlContainer); // Add control container here
    colDiv.appendChild(input);
    container.appendChild(colDiv);
})


document.addEventListener('DOMContentLoaded', () => {
    const selectButtons = document.querySelectorAll('.skill');
    selectButtons.forEach(button => {
        if (button) {
            button.addEventListener('click', function() {
                this.classList.toggle('active-skill');
            });
        }
    })
});



function updatePlaceholder(newText) {
    const searchInput = document.getElementById('searchInput');
    searchInput.placeholder = newText;
}

function placeHolderTexts(buttonText){
    switch (buttonText){
        case "Ratio":
            return "Ratio threshold (sum)"
        case "Experience":
            return "Minimum Experience (per skill)"
        case "Levels":
            return "Minimum Level (per skill)"
    }
}

function InputPlaceHolderTexts(buttonText){
    switch (buttonText){
        case "Ratio":
            return ""
        case "Experience":
            return "13m"
        case "Levels":
            return "99"
    }
}




document.addEventListener('DOMContentLoaded', () => {
    var selectButtons = document.querySelector('.searchFilter').querySelectorAll('.card-body');
    var inputTexts = document.querySelectorAll('.textInput');
    const searchContainer = document.querySelector('.search-container');

    selectButtons.forEach(button => {
        button.addEventListener('click', () => {
            // Remove 'active-search-filter' from all card bodies
            selectButtons.forEach(otherButton => {
                otherButton.classList.remove('active-search-filter');
            });

            // Add 'active-search-filter' to the clicked card body
            button.classList.add('active-search-filter');
            updatePlaceholder(placeHolderTexts(button.innerText));

            const placeHolder = InputPlaceHolderTexts(button.innerText)
            if (button.innerText === "Ratio"){
                inputTexts.forEach(inputForm =>{
                    inputForm.style = "display:none;"
                })
                searchContainer.style = "display:flex;"
            }else{
                searchContainer.style = "display:none;"
                inputTexts.forEach(inputForm =>{
                    inputForm.placeholder = placeHolder
                    inputForm.style = "display:unset;"
                })
            }




        });
    });
});

