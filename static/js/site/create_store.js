document.addEventListener('DOMContentLoaded', function () {
    console.log('DOM полностью загружен');

    // Функция для обработки отправки формы
    function handleStoreFormSubmit() {
        console.log("Обработчик сработал");

        const storeForm = document.getElementById('addStore');
    console.log(storeForm); // Убедитесь, что форма найдена
    
    const storeAddressInput = storeForm ? storeForm.querySelector('#storeAddress') : null;
    console.log(storeAddressInput); // Убедитесь, что поле с адресом найдено
    
    if (!storeAddressInput) {
        console.error('Поле ввода адреса магазина не найдено');
        return;
    }

        const storeAddress = storeAddressInput.value.trim();

        if (storeAddress === '') {
            alert('Пожалуйста, введите адрес магазина');
            return;
        }

        // Создаем объект данных магазина
        const storeData = {
            address: storeAddress,
        };

        console.log('Отправляем запрос на сервер с данными:', storeData);

        // Отправляем данные магазина на сервер
        fetch('/store', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${getCookieByKey('access_token')}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(storeData)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log('Получен новый магазин:', data);
            const { id } = data;

            // Находим таблицу магазинов и добавляем новый элемент
            const storeTable = document.querySelector('table');
            if (!storeTable) {
                console.error('Таблица магазинов не найдена');
                return;
            }

            // Создаем новый элемент строки таблицы
            const newRow = document.createElement('tr');
            newRow.className = 'store-item';

            // Добавляем данные в элемент строки таблицы
            newRow.innerHTML = `
                <td class="store-item" data-type="id">${id}</td>
                <td class="store-item" data-type="address">
                    <a href="store/${id}">${storeAddress}</a>
                </td>
            `;

            // Добавляем новую строку в конец таблицы
            storeTable.tBodies[0].appendChild(newRow);

            // Очищаем поля формы после добавления
            storeForm.reset();
        })
        .catch((error) => {
            console.error('Ошибка при добавлении магазина:', error);
            alert('Не удалось добавить магазин. Проверьте подключение к интернету и попробуйте снова.');
        });
    }

    // Привязываем событие к кнопке вместо формы
    const submitButton = document.getElementById('submitStoreForm');
    if (submitButton) {
        console.log('Кнопка для добавления магазина найдена');
        submitButton.addEventListener('click', handleStoreFormSubmit);
    } else {
        console.error('Кнопка для добавления магазина не найдена');
    }
});
