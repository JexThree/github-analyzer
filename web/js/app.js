let chart = null;

async function searchUser() {

    const username = document.getElementById("username").value;

    if (!username) {
        alert("Ingresa un usuario");
        return;
    }

    /*
      MÁS ADELANTE:
      const response = await fetch(`/github/${username}`);
      const data = await response.json();
    */

    const data = {
        name: "Linus Torvalds",
        login: username,
        repos: 8,
        followers: 250000,
        language: "C",
        popularRepo: "linux",

        languages: {
            C: 70,
            Python: 20,
            Go: 10
        }
    };

    document.getElementById("result").classList.remove("hidden");

    document.getElementById("name").textContent = data.name;
    document.getElementById("login").textContent = data.login;
    document.getElementById("repos").textContent = data.repos;
    document.getElementById("followers").textContent = data.followers;
    document.getElementById("language").textContent = data.language;
    document.getElementById("popularRepo").textContent = data.popularRepo;

    renderChart(data.languages);
}

function renderChart(languages) {

    const ctx = document.getElementById("languageChart");

    if (chart) {
        chart.destroy();
    }

    chart = new Chart(ctx, {
        type: "pie",
        data: {
            labels: Object.keys(languages),
            datasets: [{
                data: Object.values(languages)
            }]
        }
    });
}