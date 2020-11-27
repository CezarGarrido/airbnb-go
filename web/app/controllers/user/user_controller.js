function NewUserController() {

    function createUser(opts) {

        return fetch('http://localhost:3000/api/v1/users', {
            method: 'post',
            body: JSON.stringify(opts)
        })
    }

    function createNewUser() {
        let user = parseUserForm();
        createUser(user).then(
            function (response) {
                if (response.status !== 200) {
                    console.log('Looks like there was a problem. Status Code: ' +
                        response.status);
                    return;
                }
                response.json().then(function (data) {
                    console.log("User create ok");
                    console.log(data);
                });
            }
        );
    }

    function parseUserForm() {

        let name = document.getElementById("input-register-name").value;
        let email = document.getElementById("input-register-email").value;
        let password = document.getElementById("input-register-password").value;
        let repeatPassword = document.getElementById("input-register-repeat-password").value;

        return {
            name: name,
            email: email,
            password: password,
            repeat_password: repeatPassword,
        }
    }

    function listeners() {
        let btnSubmit = document.getElementById("btn-resgister-user");
        btnSubmit.addEventListener("click", createNewUser);
    }

    return {
        createUser: createUser,
        parseUserForm: parseUserForm,
        listeners: listeners
    }
}