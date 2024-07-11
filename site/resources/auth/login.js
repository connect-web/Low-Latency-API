document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("loginForm");

    form.addEventListener("submit", function(event) {
        event.preventDefault();

        const csrfToken = fetchCsrfTokenIfNeeded();

        const formData = new FormData(form);
        const jsonData = {};

        formData.forEach((value, key) => {
            jsonData[key] = value;
        });

        sendRequest('/api/login', jsonData, csrfToken)
            .then(data => {
                console.log(data);
                // Handle the response data
            })
            .catch(error => {
                console.error("Error:", error);
                // Handle the error
            });
    });
});
