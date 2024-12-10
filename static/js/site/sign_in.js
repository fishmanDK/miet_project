document.addEventListener('DOMContentLoaded', function () {
    // Находим форму и добавляем обработчик события "submit"
    const form = document.getElementById('signinForm');
    if (form) {
        form.addEventListener('submit', function (event) {
            event.preventDefault(); // Предотвращаем отправку формы
            sign_in(); // Вызываем функцию авторизации
        });
    } else {
        console.error('Form with id "signinForm" not found.');
    }
});

function sign_in() {
    console.log("Sign-in process started");

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const errorMessage = document.getElementById('errorMessage');

    errorMessage.textContent = ""; // Сбрасываем ошибку

    if (!email.trim()) {
        errorMessage.textContent = "Email field cannot be empty.";
        return;
    }

    if (!password.trim()) {
        errorMessage.textContent = "Password field cannot be empty.";
        return;
    }

    if (password.length < 8) {
        errorMessage.textContent = "Password must be at least 8 characters long.";
        return;
    }

    const data = { email, password };

    fetch('http://localhost:8081/auth/sign-in', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
        .then((response) => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok.');
        })
        .then((data) => {
            const { access_token, refresh_token } = data;
            document.cookie = `access_token=${access_token}; path=/;`;
            document.cookie = `refresh_token=${refresh_token}; path=/;`;
            document.getElementById('signinForm').reset();
            window.location.href = '/store';
        })
        .catch((error) => {
            console.error('Error:', error);
            errorMessage.textContent = 'Failed to sign in. Please try again.';
        });
}
