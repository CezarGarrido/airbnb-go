
function NewApartmentController() {
    function findLocations() {
        return fetch('http://localhost:3000/api/v1/apartments');
    }
    return {
        findLocations: findLocations
    }
}

