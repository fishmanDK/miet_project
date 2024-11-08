document.addEventListener('DOMContentLoaded', function () {
    console.log('DOM полностью загружен');

    // Функция для обработки отправки формы
    function handleStoreFormSubmit(event) {
        event.preventDefault();

        console.log('Форма отправлена');

        const storeForm = event.target.closest('form');
        if (!storeForm) {
            console.error('Форма для добавления магазина не найдена');
            return;
        }

        const storeAddressInput = storeForm.querySelector('#storeAddress');
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

        // Отправляем данные магазина на сервер
        fetch('/store', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${accessToken = getCookieByKey('access_token')}`,
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

    // Поиск формы по более надежному селектору
    const forms = document.getElementsByTagName('form');
    let storeFormFound = false;
    for (let i = 0; i < forms.length && !storeFormFound; i++) {
        if (forms[i].querySelector('#storeAddress')) {
            storeFormFound = true;
            break;
        }
    }

    if (storeFormFound) {
        console.log('Форма для добавления магазина найдена');
        const form = forms[0];
        form.addEventListener('submit', handleStoreFormSubmit);
    } else {
        console.error('Форма для добавления магазина не найдена');
    }
});
