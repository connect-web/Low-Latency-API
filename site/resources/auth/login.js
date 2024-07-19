document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("loginForm");
    const errorDialog = document.getElementById("Error");


    form.addEventListener("submit", function(event) {
        event.preventDefault();
        ShowLoading();
        const csrfToken = fetchCsrfTokenIfNeeded();

        const formData = new FormData(form);
        const jsonData = {};

        formData.forEach((value, key) => {
            jsonData[key] = value;
        });


        sendRequest('/api/auth/login', jsonData, csrfToken)
            .then(data => {
                console.log(data);
                if (data.message === "Logged in") {
                    window.location.replace("/profile");
                }else{
                    handleError(data.error)
                }
                // Handle the response data
            })
            .catch(error => {
                if (error.message === 'CSRF token is invalid or missing') {
                    // Retry the request with a new CSRF token
                    const newCsrfToken = fetchCsrfTokenIfNeeded();
                    sendRequest('/api/auth/login', jsonData, newCsrfToken)
                        .then(data => {
                            console.log(data);
                            HideLoading();
                            if (data.message === "Logged in") {
                                window.location.replace("/profile");
                            }

                            // Handle the response data
                        })
                        .catch(err => {
                            console.error("Retry error:", err);
                            HideLoading();

                            // Handle the retry error
                        });
                } else {
                    console.error("Error:", error);
                    HideLoading();
                    // Handle the error
                }
            });

    });

    // Add click event listener to the error dialog to hide it when clicked
    errorDialog.addEventListener("click", function() {
        errorDialog.classList.add("hidden");
    });
});



function handleError(err) {
    HideLoading();
    console.error("Handling err" +err)
    if (err.includes("Captcha failed")) {
        reloadCaptcha();
    }
    const errorDialog = document.getElementById("Error");
    const errorMessage = document.getElementById("errorMessage");
    errorMessage.textContent = err;
    errorDialog.classList.remove("hidden");
}

// Function to reload the CAPTCHA
function reloadCaptcha() {
    grecaptcha.reset();
}