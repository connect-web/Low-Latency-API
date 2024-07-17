function getCsrfToken() {
    const cookies = document.cookie.split(";");
    for (let i = 0; i < cookies.length; i++) {
        let cookiePair = cookies[i].split('=');
        if (cookiePair[0].trim() === "csrf_") {
            return cookiePair[1];
        }
    }
    return "";
}

function fetchCsrfTokenIfNeeded() {
    let csrfToken = getCsrfToken();
    if (csrfToken === "") {
        fetch('/api/auth/csrf', {
            method: 'GET',
        }).then(() => {
            csrfToken = getCsrfToken();
        });
    }
    return csrfToken;
}

function sendRequest(url, jsonData, csrfToken, method = "POST") {
    return fetch(url, {
        method: method,
        headers: {
            "Content-Type": "application/json",
            "X-CSRF-Token": csrfToken
        },
        body: JSON.stringify(jsonData)
    })
        .then(response => {
            if (response.status === 403) {
                return response.json().then(data => {
                    if (data.csrfToken) {
                        return sendRequest(url, jsonData, data.csrfToken, method); // Retry with new CSRF token
                    } else {
                        console.error("Error:", data.error);
                        /*
                        if (data.error === "CSRF token is invalid or missing"){
                            console.log("Csrf validation retrying!")
                            return sendRequest(url, jsonData, fetchCsrfTokenIfNeeded(), method)
                        }
                         */
                        throw new Error(data.error);
                    }
                });
            } else {
                return response.json();
            }
        });
}
