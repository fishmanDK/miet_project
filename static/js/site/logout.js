document.addEventListener('DOMContentLoaded', function () {
    const logoutLink = document.querySelector('.nav-link[href="#"]'); // Находим элемент Logout

    logoutLink.addEventListener('click', function (event) {
        event.preventDefault(); // Предотвращаем переход по ссылке

        // Удаляем куки access_token и refresh_token
        document.cookie = "access_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
        document.cookie = "refresh_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";

        // Перенаправляем на страницу авторизации
        window.location.href = '/auth/sign-in';
    });
});