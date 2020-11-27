function NewUserController() {

    async function createUser(body) {

        const options = {
            method: "POST",
            body: JSON.stringify(body),
        };

        return fetch("http://localhost:3000/api/v1/users", options)
            .then(function (response) {
                if (response.status === 500) {
                    return response.json()
                        .then((json) => {
                            throw json;
                        });
                } 
                return response.json();
            })
    }

    function createNewUser() {
        let user = parseUserForm();
        let err = Validade(user);
        if (err != null) {
            alert(err);
            return;
        }
        createUser(user)
            .then(() => alert("Cadastrado com sucesso!"))
            .catch((err) => alert(err));
    }

    function Validade(user) {
        let errors = [];

        if (user.first_name.length <= 3) {
            errors.push("Nome dever conter mais de 3 caracteres!");
        }

        if (user.email.length <= 3) {
            errors.push("Email é obrigatório!");
        }

        if (user.password <= 3) {
            errors.push("Senha é obrigatória!");
        }

        if (user.password != user.repeat_password) {
            errors.push("Senhas devem ser iguais");
        }

        if (errors.length > 0) {
            return errors.join("\n");
        }
        return null;
    }

    function parseUserForm() {
        let name = document.getElementById("input-register-name").value;
        let email = document.getElementById("input-register-email").value;
        let password = document.getElementById("input-register-password").value;
        let repeatPassword = document.getElementById(
            "input-register-repeat-password"
        ).value;

        return {
            first_name: name,
            email: email,
            password: password,
            repeat_password: repeatPassword,
        };
    }

    function listeners() {
        let btnSubmit = document.getElementById("btn-resgister-user");
        btnSubmit.addEventListener("click", createNewUser);
    }

    return {
        createUser: createUser,
        parseUserForm: parseUserForm,
        listeners: listeners,
    };
}