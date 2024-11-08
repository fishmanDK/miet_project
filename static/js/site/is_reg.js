document.addEventListener('DOMContentLoaded', function() {
    console.log("2122")
    fetch('http://localhost:8081/check', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${accessToken = getCookieByKey('access_token')}`,
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            if (response.ok) {
                return 
            }
            throw new Error('Network response was not ok.');
        })
        .then(data => {
            console.log("don't reg")
        })
        .catch(error => {
            console.error('Error:', error);
        });
});

function getCookieByKey(name) {
    const value = `; ${document.cookie}`;
    
    const parts = value.split(`; ${name}=`);
    
    if (parts.length === 2) {
        return parts.pop().split(';').shift();
    }
    
    return null;
}
