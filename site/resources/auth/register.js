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
                if (data.message === "User registered"){
                    window.location.replace("/profile");
                }
                if (data.message === "User registered, Login required."){
                    window.location.replace("/login");
                }
                // Handle the response data
            })
            .catch(error => {
                console.error("Error:", error);
                // Handle the error
            });
    });
});
