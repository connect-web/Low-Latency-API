document.addEventListener("DOMContentLoaded", function() {
    if (!document.cookie.split('; ').find(row => row.startsWith('csrf_='))) {
        fetch('/api/auth/csrf', {
            method: 'GET',
        })
            .then(response => response.json())
            .then(data => {
                console.log(data);
            })
            .catch(error => console.error('Error:', error));
    }

});
