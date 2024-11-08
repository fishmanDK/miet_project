document.addEventListener('DOMContentLoaded', function() {
    // Делегирование события клика для элементов с классом cassette-item
    document.body.addEventListener('click', function(event) {
        if (event.target.classList.contains('cassette-item')) {
            event.stopPropagation();
            handleCassetteClick(event);
        }
    });

    document.body.addEventListener('click', function(event) {
        if (event.target.id === 'show_reservations_button') {
            event.stopPropagation();
            showReservations();
        }
    });

    // Делегирование события клика для кнопки "Добавить кассету"
    document.body.addEventListener('click', function(event) {
        if (event.target.id === 'show_form_add_cassette_button') {
            event.stopPropagation();
            showFormAddCassette();
        }
    });

    // Делегирование события клика для кнопки "Забронировать" и отмены бронирования
    document.body.addEventListener('click', function(event) {
        // Проверяем, что это именно кнопка "Забронировать"
        if (event.target.id === 'reserveButton') {
            event.preventDefault(); // предотвращаем стандартное поведение, если это форма или ссылка
            event.stopPropagation(); // останавливаем распространение события
            handleReserveClick(event);
        }
        // Проверяем для кнопки отмены бронирования
        if (event.target.id === 'deleteReservationButton') {
            console.log("click");
            event.preventDefault();
            event.stopPropagation();
            handleDeleteReservationClick(event);
        }
    });
});

function handleCassetteClick(event) {
    const cassetteItems = document.querySelectorAll(`[data-id="${event.target.getAttribute('data-id')}"]`);
    let id, name, genre, year;

    cassetteItems.forEach(item => {
        const type = item.getAttribute('data-type');
        if (type === 'id') id = item.textContent;
        if (type === 'name') name = item.textContent;
        if (type === 'genre') genre = item.textContent;
        if (type === 'year') year = item.textContent;
    });

    const isRequestSent = event.target.hasAttribute('data-request-sent');
    if (isRequestSent) {
        console.log('Запрос уже отправлен для этой кассеты');
        return;
    }

    event.target.setAttribute('data-request-sent', true);

    fetch('http://localhost:8081/cassette/details/' + id, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${getCookieByKey('access_token')}`,
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok.');
        })
        .then(data => {
            console.log(data)
            const { totalCount = 0, rentedCount = 0, isOrdered = false, isReservated = false } = data;
            showCassetteDetails(id, name, genre, year, totalCount, rentedCount, isOrdered, isReservated);
        })
        .catch(error => {
            console.error('Ошибка:', error);
        })
        .finally(() => {
            event.target.removeAttribute('data-request-sent');
        });
}

function showCassetteDetails(id, name, genre, year, total_count, rented_count, isOrdered, isReservated) {
    const cassetteInfo = document.getElementById('formContainer');
    const isAdmin = document.querySelector('meta[name="is_admin"]').content === "true";

    const deleteReservationButton = `
    <button id="deleteReservationButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${id}">Отменить бронирование</button>
    `;
    const AddReservationButton = `
    <button id="reserveButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${id}">Забронировать</button>
    `;

    let detailsHtml = `
    <h4>Детали кассеты:</h4>
    <p id="cassette_id"><strong>ID:</strong> ${id}</p>
    <p><strong>Name:</strong> ${name}</p>
    <p><strong>Genre:</strong> ${genre}</p>
    <p><strong>Year of release:</strong> ${year}</p>
    <p id="total_count">
        <strong>Total count:</strong>
        <span id="total_count_value">${total_count}</span>
    </p>
    <p id="rented_count" style="display: inline-block; margin-right: 10px;">
        <strong>Rented count:</strong>
        <span id="rented_count_value">${rented_count}</span>
    </p>
    `;

    if (isAdmin) {
        detailsHtml += `
        <button id="show_reservations_button" type="button" class="btn btn-primary btn-sm" style="inline-block: block;">Показать все брони</button>

        <div style="margin-top: 10px;">
            <button id="deleteButton" style="display: inline-block;" type="button" class="btn btn-danger btn-sm">Удалить</button>
            <button id="show_form_add_cassette_button" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;">Добавить кассету</button>
        </div>
        `;
    } else {
        if (isOrdered) {
            detailsHtml += `<p style="color: green;">Данная кассета уже заказана</p>`;
        } else if (total_count === 0) {
            if (isReservated){
                console.log(1)
                detailsHtml += deleteReservationButton;
            } else {
                console.log(2)
                detailsHtml += AddReservationButton;
            }
            
        } else {
            detailsHtml += `
            <button id="orderButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;">Заказать</button>
            `;
        }
    }

    cassetteInfo.innerHTML = detailsHtml;
}


function showReservations (event) {
    const button = event.target; // Получаем кнопку, на которую кликнули
    const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10); // Теперь переменная button уже доступна

    const cassette_id = button.getAttribute('data-cassette-id');

    const dataCreateReservation = {
       cassette_id: parseInt(cassette_id, 10),
    };
}

function handleDeleteReservationClick(event) {
    const button = event.target; // Получаем кнопку, на которую кликнули
    const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10); // Теперь переменная button уже доступна

    const cassette_id = button.getAttribute('data-cassette-id');

    const dataCreateReservation = {
       cassette_id: parseInt(cassette_id, 10),
    };

    const AddReservationButton = `
    <button id="reserveButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${cassetteId}">Забронировать</button>
    `;

    fetch(`http://localhost:8081/reservation`, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${getCookieByKey('access_token')}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dataCreateReservation)
        })
        .then(response => {
            if (response.ok) {
                const buttonContainer = button.parentElement;
                button.remove();
                buttonContainer.insertAdjacentHTML('beforeend', AddReservationButton);
            } else {
                throw new Error('Ошибка при отмене бронирования.');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
        });
}

function handleReserveClick(event) {
    const button = event.target; // Получаем кнопку, на которую кликнули
    const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10); // Теперь переменная button уже доступна

    const user_id = document.querySelector('meta[name="user_id"]').content;
    const cassette_id = button.getAttribute('data-cassette-id');
    const dataCreateReservation = {
        user_id: parseInt(user_id, 10),
        cassette_id: parseInt(cassette_id, 10),
    };

    const deleteReservationButton = `
    <button id="deleteReservationButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${cassetteId}">Отменить бронирование</button>
    `;

    // Если кнопка уже заблокирована, не выполнять запрос
    if (button.hasAttribute('data-loading')) {
        return;
    }

    // Устанавливаем атрибут, чтобы указать, что запрос в процессе
    button.setAttribute('data-loading', true);
    button.disabled = true; // Блокируем кнопку

    fetch(`http://localhost:8081/reservation`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${getCookieByKey('access_token')}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dataCreateReservation)
        })
        .then(response => {
            if (response.ok) {
                const buttonContainer = button.parentElement;
                button.remove();
                buttonContainer.insertAdjacentHTML('beforeend', deleteReservationButton);
            } else {
                throw new Error('Ошибка при бронировании.');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
        })
        .finally(() => {
            // Убираем атрибут и разблокируем кнопку, когда запрос завершен
            button.removeAttribute('data-loading');
            button.disabled = false;
        });
}
