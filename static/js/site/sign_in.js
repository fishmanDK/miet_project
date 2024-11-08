document.getElementById("signinForm").addEventListener("submit", function(event) {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const errorMessage = document.getElementById("errorMessage");

    // Очищаем сообщение об ошибке
    errorMessage.textContent = "";

    // Проверяем, что поле email не пустое
    if (email.trim() === "") {
        errorMessage.textContent = "Email field cannot be empty.";
        event.preventDefault(); // Предотвращаем отправку формы
        return;
    }

    // Проверяем, что пароль не пустой и его длина больше 8 символов
    if (password.trim() === "") {
        errorMessage.textContent = "Password field cannot be empty.";
        event.preventDefault(); // Предотвращаем отправку формы
        return;
    }

    if (password.length < 8) {
        errorMessage.textContent = "Password must be at least 8 characters long.";
        event.preventDefault(); // Предотвращаем отправку формы
        return;
    }
});



document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('signinForm').addEventListener('submit', function (event) {
        event.preventDefault(); 

        // Собираем данные из формы
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        // Создаем объект данных
        const data = {
            email: email,
            password: password,
        };

        // Отправляем данные на сервер
        fetch('http://localhost:8081/auth/sign-in', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json' // Указываем, что отправляем JSON
            },
            body: JSON.stringify(data) // Преобразуем объект в JSON
        })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok.');
        })
        .then(data => {
            const { access_token, refresh_token } = data;
            
            document.cookie = `access_token=${access_token}; path=/;`;
            document.cookie = `refresh_token=${refresh_token}; path=/;`;

            // Очищаем поля формы после добавления кассеты
            document.getElementById('signinForm').reset();
            window.location.href = '/store';
        })
        .catch((error) => {
            console.error('Ошибка:', error);
        });
    });
});
