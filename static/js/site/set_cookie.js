function getCookieByKey(name) {
    // Добавляем точку с запятой перед значением cookie, чтобы избежать частичных совпадений
    const value = `; ${document.cookie}`;
    
    // Разбиваем строку cookie по имени переданного ключа
    const parts = value.split(`; ${name}=`);
    
    // Проверяем, существует ли cookie с переданным именем
    if (parts.length === 2) {
        // Возвращаем значение cookie
        return parts.pop().split(';').shift();
    }
    
    // Возвращаем null, если cookie не найдено
    return null;
}

function fetchWithTokens(url, method, data) {
    const accessToken = getCookieByKey('access_token');
    const refreshToken = getCookieByKey('refresh_token');

    const headers = {
        'Content-Type': 'application/json', // Указываем, что отправляем JSON
    };

    // Добавляем токены в заголовки, если они существуют
    if (accessToken) {
        headers['Authorization'] = `Bearer ${accessToken}`;
    }

    if (refreshToken) {
        headers['Refresh-Token'] = refreshToken; // Используйте этот заголовок, если нужно
    }

    return fetch(url, {
        method: method,
        headers: headers,
        body: JSON.stringify(data),
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        }
        throw new Error('Network response was not ok.');
    });
}