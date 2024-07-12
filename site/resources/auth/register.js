document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("register-form");

    form.addEventListener("submit", function(event) {
        event.preventDefault();

        const formData = new FormData(form);
        const jsonData = {};

        formData.forEach((value, key) => {
            jsonData[key] = value;
        });

        const csrfToken = fetchCsrfTokenIfNeeded();

        sendRequest('/api/auth/register', jsonData, csrfToken)
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
