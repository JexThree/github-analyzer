const API_URL = "http://localhost:8080/users";

window.onload = () => {
    loadUsers();
};

async function loadUsers() {

    const response = await fetch(API_URL);

    const users = await response.json();

    const container =
        document.getElementById("usersContainer");

    container.innerHTML = "";

    users.forEach(user => {

        container.innerHTML += `
            <div class="user-card">

                <h3>
                    ${user.name}
                    ${user.lastname}
                </h3>

                <p>${user.githubEmail}</p>

                <p>${user.birthDate}</p>

                <button
                    onclick="deleteUser('${user.id}')"
                >
                    Eliminar
                </button>

            </div>
        `;
    });
}

async function createUser() {

    const user = {

        name:
            document.getElementById("name").value,

        lastname:
            document.getElementById("lastname").value,

        githubEmail:
            document.getElementById("email").value,

        birthDate:
            document.getElementById("birthDate").value
    };

    await fetch(API_URL, {

        method: "POST",

        headers: {
            "Content-Type": "application/json"
        },

        body: JSON.stringify(user)
    });

    loadUsers();
}
async function deleteUser(id) {

    await fetch(`${API_URL}/${id}`, {

        method: "DELETE"
    });

    loadUsers();
}