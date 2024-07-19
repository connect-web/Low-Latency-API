document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("register-form");
    const errorDialog = document.getElementById("Error");


    form.addEventListener("submit", function(event) {
        event.preventDefault();

        const formData = new FormData(form);
        const jsonData = {};

        formData.forEach((value, key) => {
            jsonData[key] = value;
        });

        const csrfToken = fetchCsrfTokenIfNeeded();
        ShowLoading()
        sendRequest('/api/auth/register', jsonData, csrfToken)
            .then(data => {
                if (data.message === "User registered"){
                    window.location.replace("/profile");
                    return;
                }
                if (data.message === "User registered, Login required."){
                    window.location.replace("/login");
                    return;
                }else{
                    handleError(data.error)
                }
                // Handle the response data
            })
            .catch(error => {
                console.error("Error:", error);
                if (error.message === 'CSRF token is invalid or missing') {
                    // Retry the request with a new CSRF token
                    const newCsrfToken = fetchCsrfTokenIfNeeded();
                    sendRequest('/api/auth/register', jsonData, csrfToken)
                        .then(data => {
                            if (data.message === "User registered"){
                                window.location.replace("/profile");
                            }
                            if (data.message === "User registered, Login required."){
                                window.location.replace("/login");
                            }
                            HideLoading();
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
                // Handle the error
            });

    }
    );


    // Add click event listener to the error dialog to hide it when clicked
    errorDialog.addEventListener("click", function() {
        errorDialog.classList.add("hidden");
    });
});




function handleError(err) {
    console.error("Handling err" +err)
    HideLoading();
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