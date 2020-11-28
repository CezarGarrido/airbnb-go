
function NewApartmentController() {

    function findLocations() {
        return fetch('http://localhost:3000/api/v1/apartments');
    }

    function findByUUID(uuid) {
        return fetch('http://localhost:3000/api/v1/apartments/'+ uuid);
    }

    return {
        findByUUID,
        findLocations
    }
}

