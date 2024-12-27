document.addEventListener("DOMContentLoaded", () => {
    const userTableBody = document.querySelector("#userTable tbody");
    const addUserForm = document.querySelector("#addUserForm");

    // Fungsi untuk mengambil dan menampilkan data user
    const fetchUsers = async () => {
        try {
            const response = await fetch("/users");
            if (!response.ok) throw new Error("Failed to fetch users");
            const users = await response.json();

            userTableBody.innerHTML = ""; // Hapus konten tabel sebelumnya
            users.forEach(user => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${user.id}</td>
                    <td>${user.name}</td>
                    <td>${user.email}</td>
                    <td>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${user.id}">Delete</button>
                    </td>
                `;
                userTableBody.appendChild(row);
            });

            // Tambahkan event listener untuk tombol hapus
            document.querySelectorAll(".delete-btn").forEach(button => {
                button.addEventListener("click", async (e) => {
                    const userId = e.target.getAttribute("data-id");
                    await deleteUser(userId);
                });
            });
        } catch (error) {
            console.error(error);
            alert("Failed to fetch users. Please try again later.");
        }
    };

    // Fungsi untuk menambahkan user baru
    addUserForm.addEventListener("submit", async (e) => {
        e.preventDefault();
        const formData = new FormData(addUserForm);
        const user = {
            name: formData.get("name"),
            email: formData.get("email"),
        };

        try {
            const response = await fetch("/users", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(user),
            });
            if (!response.ok) throw new Error("Failed to add user");
            addUserForm.reset(); // Reset form
            fetchUsers(); // Refresh tabel
        } catch (error) {
            console.error(error);
            alert("Failed to add user. Please try again later.");
        }
    });

    // Fungsi untuk menghapus user
    const deleteUser = async (id) => {
        try {
            const response = await fetch(`/users/${id}`, {
                method: "DELETE",
            });
            if (!response.ok) throw new Error("Failed to delete user");
            fetchUsers(); // Refresh tabel
        } catch (error) {
            console.error(error);
            alert("Failed to delete user. Please try again later.");
        }
    };

    // Panggil fetchUsers saat halaman dimuat
    fetchUsers();
});
