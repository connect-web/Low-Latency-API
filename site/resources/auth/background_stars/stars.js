document.addEventListener('DOMContentLoaded', function () {
    const numStars = 100; // Adjust the number of stars as needed

    function generateStars(starContainerId) {
        const starContainer = document.getElementById(starContainerId);
        const width = window.innerWidth;
        const height = window.innerHeight;
        const starSize = 2; // Adjust size of the stars

        for (let i = 0; i < numStars; i++) {
            const star = document.createElement('div');
            star.style.position = 'absolute';
            star.style.width = `${starSize}px`;
            star.style.height = `${starSize}px`;
            star.style.backgroundColor = '#FFF';
            star.style.boxShadow = '0 0 2px #FFF';
            star.style.borderRadius = '50%';
            star.style.top = `${Math.random() * height}px`;
            star.style.left = `${Math.random() * width}px`;
            starContainer.appendChild(star);
        }
    }

    generateStars('stars');
    generateStars('stars2');
    generateStars('stars3');

    window.addEventListener('resize', () => {
        document.getElementById('stars').innerHTML = '';
        document.getElementById('stars2').innerHTML = '';
        document.getElementById('stars3').innerHTML = '';
        generateStars('stars');
        generateStars('stars2');
        generateStars('stars3');
    });
});
