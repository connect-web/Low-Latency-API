function ShowLoading() {
    const loadingContainer = document.querySelector('.loading-container');
    if (loadingContainer) {
        loadingContainer.classList.remove('hidden');
    }
}

function HideLoading() {
    const loadingContainer = document.querySelector('.loading-container');
    if (loadingContainer) {
        loadingContainer.classList.add('hidden');
    }
}

document.addEventListener("DOMContentLoaded", function() {
    HideLoading();
});
