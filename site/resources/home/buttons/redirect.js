document.addEventListener('DOMContentLoaded', function() {
    // Function to handle the redirection
    function handleButtonClick(event) {
        var target = event.target;
        if (target.id === 'login') {
            window.location.href = '/login';
        } else if (target.id === 'register') {
            window.location.href = '/register';
        }
    }

    // Add event listeners to the buttons
    document.getElementById('login').addEventListener('click', handleButtonClick);
    document.getElementById('register').addEventListener('click', handleButtonClick);
});
