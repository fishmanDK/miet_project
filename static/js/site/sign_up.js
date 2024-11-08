document.getElementById("signupForm").addEventListener("submit", function(event) {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirmPassword").value;
    const errorMessage = document.getElementById("errorMessage");

    // Сбрасываем сообщение об ошибке
    errorMessage.textContent = "";

    // Проверяем, что поле email не пустое
    if (email.trim() === "") {
        errorMessage.textContent = "Email field cannot be empty.";
        event.preventDefault(); // Предотвращаем отправку формы
        return;
    }

    // Проверяем длину пароля
    if (password.length < 8) {
        errorMessage.textContent = "Password must be at least 8 characters long.";
        event.preventDefault(); // Предотвращаем отправку формы
        return;
    }

    // Проверяем, что пароль не пустой
    if (password.trim() === "") {
        errorMessage.textContent = "Password field cannot be empty.";
        event.preventDefault(); // Предотвращаем отправку формы
        return;
    }

    // Проверяем совпадение паролей
    if (password !== confirmPassword) {
        errorMessage.textContent = "Passwords do not match.";
        event.preventDefault(); // Предотвращаем отправку формы
        return;
    }
});

document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('signupForm').addEventListener('submit', function (event) {
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
        fetch('http://localhost:8081/auth/sign-up', {
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
            console.log('2:', data);
            const { id } = data;
            
            // Очищаем поля формы после добавления кассеты
            document.getElementById('signupForm').reset();
            window.location.href = '/auth/sign-in';
        })
        .catch((error) => {
            console.error('Ошибка:', error);
        });
    });
});
