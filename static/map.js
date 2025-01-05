const API_BASE_URL = "http://localhost:8080"

let map = L.map('map').setView([0, 0], 2);

// set up the OpenStreetMap layer
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors'
}).addTo(map);

// Event listener for map clicks
map.on('click', function(e) {
    let latitude = e.latlng.lat;
    let longitude = e.latlng.lng;

    // Place a marker on the map
    L.marker([latitude, longitude]).addTo(map);

    // Send geolocation data to the backend
    fetch(`${API_BASE_URL}/api/stops`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ latitude: latitude, longitude: longitude })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }
        return response.json();
    })
    .then(data => console.log('Success:', data))
    .catch(error => console.error('Error:', error));
});
