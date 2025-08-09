document.getElementById("appForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const data = {
        company: document.getElementById("company").value,
        position: document.getElementById("position").value,
        date: document.getElementById("date").value,
        link: document.getElementById("link").value,
        notes: document.getElementById("notes").value,
        status: document.getElementById("status").value
    };
    await fetch("/applications", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data)
    });
    loadApplications();
});

async function loadApplications() {
    const res = await fetch("/applications");
    const apps = await res.json();
    const tbody = document.querySelector("#applicationsTable tbody");
    tbody.innerHTML = "";
    apps.forEach(app => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${app.company}</td>
            <td>${app.position}</td>
            <td>${app.date}</td>
            <td>${app.status}</td>
            <td>
                <button onclick="deleteApplication(${app.id})">Delete</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

async function deleteApplication(id) {
    await fetch(`/applications/${id}`, { method: "DELETE" });
    loadApplications();
}

loadApplications();
