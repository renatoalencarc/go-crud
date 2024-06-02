document.getElementById('createUserForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const userId = document.getElementById('userId').value;
    const userName = document.getElementById('userName').value;
    const userDescription = document.getElementById('userDescription').value;

    fetch('http://localhost:8081/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            id: userId,
            name: userName,
            description: userDescription
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log('User created:', data);
        alert('User created successfully');
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Error creating user');
    });
});

document.getElementById('deleteUserForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const userId = document.getElementById('deleteUserId').value;

    fetch(`http://localhost:8081/users/${userId}`, {
        method: 'DELETE'
    })
    .then(response => response.json())
    .then(data => {
        console.log('User deleted:', data);
        alert('User deleted successfully');
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Error deleting user');
    });
});

document.getElementById('listUsersButton').addEventListener('click', function() {
    fetch('http://localhost:8081/users', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => response.json())
    .then(users => {
        const userList = document.getElementById('userList');
        userList.innerHTML = '';
        users.forEach(user => {
            const listItem = document.createElement('li');
            listItem.textContent = `ID: ${user.id}, Name: ${user.name}, Description: ${user.description}`;
            userList.appendChild(listItem);
        });
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Error fetching users');
    });
});
