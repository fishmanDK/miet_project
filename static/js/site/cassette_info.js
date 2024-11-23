document.addEventListener('DOMContentLoaded', function () {
    document.body.addEventListener('click', function (event) {
        if (event.target.classList.contains('cassette-item')) {
            event.stopPropagation();
            handleCassetteClick(event);
        }

        if (event.target.id === 'orderButton') {
            event.preventDefault();
            event.stopPropagation();
            handleOrderClick(event);
        }

        if (event.target.id === 'show_reservations_button') {
            event.stopPropagation();
            event.stopImmediatePropagation();
            showOrders(event);
        }

        if (event.target.id === 'close_reservations_view_button') {
            event.stopPropagation();
            event.stopImmediatePropagation();
            closeViewReservations(event);
        }

        if (event.target.id === 'show_form_add_cassette_button') {
            event.stopPropagation();
            showFormAddCassette();
        }

        if (event.target.id === 'reserveButton') {
            event.preventDefault();
            event.stopPropagation();
            handleReserveClick(event);
        }

        if (event.target.id === 'deleteButton') {
            event.preventDefault();
            event.stopPropagation();
            handleDeleteCassette(event);
        }

        if (event.target.id === 'deleteReservationButton') {
            event.preventDefault();
            event.stopPropagation();
            handleDeleteReservationClick(event);
        }
    });
});

function saveChangesButton(event) {
    const button = event.target;

    let elems = document.getElementsByClassName("info-row")

    for (let elem of elems) {
        console.log(elem)
    }

    console.log(1)
}

function handleOrderClick(event) {
    const button = event.target;

    if (button.hasAttribute('data-request-sent') && button.getAttribute('data-request-sent') === 'true') {
        console.log('Запрос уже отправлен');
        return;
    }

    button.setAttribute('data-request-sent', 'true');

    const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10);
    const element = document.getElementById("cassetteid");
    if (!element) {
        console.error("Элемент с id 'cassetteid' не найден.");
        return;
    }

    const cassetteID = element.textContent;
    const storeID = document.querySelector('meta[name="store_id"]').content;
    const userID = document.querySelector('meta[name="user_id"]').content;

    const dataOrder = {
        userId: parseInt(userID, 10),
        cassetteId: parseInt(cassetteID, 10),
        storeId: parseInt(storeID, 10)
    };

    fetch(`http://localhost:8081/orders`, {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${getCookieByKey('access_token')}`,
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(dataOrder),
    })
        .then(response => {
            if (response.ok) {
                updateCounters();

                button.disabled = true;
                console.log('Заказ успешно оформлен');
            } else {
                throw new Error('Ошибка при оформлении заказа.');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Не удалось оформить заказ. Попробуйте снова.');
        })
        .finally(() => {
            button.removeAttribute('data-request-sent');
        });
}

function updateCounters() {
    const totalCountElement = document.getElementById('total_count_value');
    const rentedCountElement = document.getElementById('rented_count_value');

    if (!totalCountElement || !rentedCountElement) {
        console.error("Элементы для счетчиков не найдены");
        return;
    }

    let totalCount = parseInt(totalCountElement.textContent.replace('Total count: ', ''), 10);
    let rentedCount = parseInt(rentedCountElement.textContent.replace('Order count: ', ''), 10);

    if (isNaN(totalCount) || isNaN(rentedCount)) {
        console.error("Некорректные значения счетчиков");
        return;
    }

    totalCount -= 1;
    rentedCount += 1;

    totalCountElement.textContent = `${totalCount}`;
    rentedCountElement.textContent = `${rentedCount}`;


}

function closeViewReservations(event) {
    const button = event.target;
    const ViewAllOrdersButton = `
        <button id="show_reservations_button" type="button" class="btn btn-primary btn-sm">Показать все заказы</button>
    `;
    const buttonContainer = button.parentElement;
    button.remove();
    buttonContainer.insertAdjacentHTML('beforeend', ViewAllOrdersButton);

    const cassetteInfo = document.getElementById('reservationsTableContainer');
    cassetteInfo.remove();
}

function showFormAddCassette() {
    const cassetteInfo = document.getElementById('formContainer');
    if (document.getElementById('cassetteForm')) {
        console.log("Форма уже существует, повторное создание не требуется.");
        return;
    }

    const newHtml = `
    <h3>Добавление новой кассеты</h3>
    <form id="cassetteForm">
        <div class="mb-3">
            <label for="cassetteName" class="form-label">Name</label>
            <input type="text" class="form-control" id="cassetteName">
        </div>
        <div class="mb-3">
            <label for="cassetteGenre" class="form-label">Genre</label>
            <input type="text" class="form-control" id="cassetteGenre">
        </div>
        <div class="mb-3">
            <label for="yearOfRelease" class="form-label">Year of release</label>
            <input type="number" class="form-control" id="yearOfRelease" min="1900" max="2100" step="1" placeholder="Enter year (e.g. 2024)">
        </div>
        <div class="mb-3">
            <label for="totalCount" class="form-label">Total count</label>
            <input type="number" class="form-control" id="totalCount" step="1">
        </div>
        <button type="submit" class="btn btn-primary">Добавить</button>
    </form>
    `;
    cassetteInfo.innerHTML = newHtml;

    document.getElementById('cassetteForm').addEventListener('submit', function (event) {
        event.preventDefault();
        const name = document.getElementById('cassetteName').value;
        const genre = document.getElementById('cassetteGenre').value;
        const year = document.getElementById('yearOfRelease').value;
        const storeID = document.querySelector('meta[name="store_id"]').content;
        const totalCount = document.getElementById('totalCount').value;

        const data = {
            name: name,
            genre: genre,
            year: parseInt(year, 10),
            storeId: parseInt(storeID, 10),
            totalCount: parseInt(totalCount, 10),
        };

        fetch('http://localhost:8081/cassette', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${getCookieByKey('access_token')}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok.');
                }
                return response.json();
            })
            .then(data => {
                const { id } = data;
                const cassetteList = document.querySelector('table');
                if (!cassetteList) {
                    console.error("Не найдена таблица для добавления строки");
                    return;
                }

                const newRow = document.createElement('tr');
                newRow.innerHTML = `
                <td class="cassette-item" data-id="${id}" data-type="id">${id}</td>
                <td class="cassette-item" data-id="${id}" data-type="name">${name}</td>
                <td class="cassette-item" data-id="${id}" data-type="genre">${genre}</td>
                <td class="cassette-item" data-id="${id}" data-type="year">${year}</td>
            `;
                cassetteList.appendChild(newRow);
                document.getElementById('cassetteForm').reset();
            })
            .catch((error) => {
                console.error('Ошибка:', error);
            });
    });
}

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
            console.log(data);
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
    <button id="deleteReservationButton" type="button" class="btn btn-danger btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${id}">Отменить бронирование</button>
    `;
    const addReservationButton = `
    <button id="reserveButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${id}">Забронировать</button>
    `;

    // <span id="total_count_value"><span id="totalcount" class="cassette-text">${total_count}</span></span> 
    // <span id="rented_count_value"><span id="rentedcount" class="cassette-text">${rented_count}</span></span> 

    let detailsHtml = `
    <div class="cassette-info">
        <h4>Детали кассеты:</h4>
        <p id="cassette_id"><strong>ID: </strong><span id="cassetteid">${id}</span></p>
        <p class="cassette-inf info-row"><strong>Name:</strong> <span id="cassettename" class="cassette-text">${name}</span> <a href="#" class="edit-icon" onclick="editCassette(this, 'text', 'Name')">✏️</a></p>
        <p class="cassette-inf info-row"><strong>Genre:</strong> <span id="cassettegenre" class="cassette-text">${genre}</span> <a href="#" class="edit-icon" onclick="editCassette(this, 'text', 'Genre')">✏️</a></p>
        <p class="cassette-inf info-row" id="total_count">
            <strong>Total count:</strong>
            <span id="totalcount" class="cassette-text">${total_count}</span>
            <a href="#" class="edit-icon" onclick="editCassette(this, 'number', 'Total count')">✏️</a>
        </p>
        <p class="cassette-info info-row" id="remain" style="display: block; margin-right: 10px;">
            <strong>Remain:</strong>
            <span id="remaincount" class="cassette-text">${parseInt(total_count, 10) - parseInt(rented_count, 10)}</span>
        </p>
        <p class="cassette-info info-row" id="rented_count" style="display: inline-block; margin-right: 10px;">
            <strong>Order count:</strong>
            <span id="rentedcount" class="cassette-text">${rented_count}</span>
        </p>
    </div>
    `;

    if (isAdmin) {
        detailsHtml += `
        <button id="show_reservations_button" type="button" class="btn btn-primary btn-sm" style="inline-block: block;">Показать все заказы</button>
        <div id="reservationsTableContainer" style="display: none; margin-top: 20px;"></div>
        <div style="margin-top: 10px;">
            <button id="deleteButton" style="display: inline-block;" type="button" class="btn btn-danger btn-sm">Удалить</button>
            <button id="show_form_add_cassette_button" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;">Добавить кассету</button>
        </div>
        `;
    } else {
        if (isOrdered) {
            detailsHtml += `<p style="color: green;">Данная кассета уже заказана</p>`;
        } else if (total_count === 0) {
            if (isReservated) {
                detailsHtml += deleteReservationButton;
            } else {
                detailsHtml += addReservationButton;
            }
        } else {
            detailsHtml += `
            <button id="orderButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;">Заказать</button>
            `;
        }
    }

    cassetteInfo.innerHTML = detailsHtml;
}

function showOrders(event) {
    const button = event.target;
    const tableContainer = document.getElementById('reservationsTableContainer');
    tableContainer.style.display = 'block';

    const element = document.getElementById("cassetteid");
    if (!element) {
        console.error("Элемент с id 'cassetteid' не найден.");
        return;
    }

    const cassetteID = element.textContent;
    const storeID = document.querySelector('meta[name="store_id"]').content;

    fetch('http://localhost:8081/admin/orders?store_id=' + storeID + "&cassette_id=" + cassetteID, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${getCookieByKey('access_token')}`,
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Не удалось получить данные.');
            }
            return response.json();
        })
        .then(data => {
            console.log('Полученные данные бронирований:', data);

            tableContainer.innerHTML = '';

            const table = document.createElement('table');
            table.classList.add('table');

            const thead = document.createElement('thead');
            const headerRow = document.createElement('tr');
            headerRow.innerHTML = `
            <th>Cassette id</th>
            <th>Дата бронирования</th>
            <th>Email</th>
            <th>Action</th>
        `;
            thead.appendChild(headerRow);
            table.appendChild(thead);

            const tbody = document.createElement('tbody');

            if (!data || data.length === 0) {
                const noDataMessage = document.createElement('tr');
                noDataMessage.innerHTML = `<td colspan="2">Заказов еще не было</td>`;
                tbody.appendChild(noDataMessage);
            } else {
                data.forEach(reservation => {
                    const newRow = document.createElement('tr');
                    newRow.setAttribute('data-cassette-id', reservation.cassette_id);

                    newRow.innerHTML = `
                    <td>${reservation.cassette_id}</td>
                    <td>${reservation.reservation_date}</td>
                    <td>${reservation.email}</td>
                    <td>
                        <a href="#" class="delete-icon" title="Удалить заказ пользователя" onclick="deleteOrder(${reservation.cassette_id}, ${reservation.user_id})">🗑️</a>
                    </td>
                `;
                    tbody.appendChild(newRow);
                });
            }

            table.appendChild(tbody);
            tableContainer.appendChild(table);

            let closeReservationButton = document.getElementById('close_reservations_view_button');
            if (!closeReservationButton) {
                console.log('Кнопка "Скрыть заказы" ещё не существует, создаём её.');

                closeReservationButton = document.createElement('button');
                closeReservationButton.setAttribute('id', 'close_reservations_view_button');
                closeReservationButton.classList.add('btn', 'btn-primary', 'btn-sm');
                closeReservationButton.textContent = 'Скрыть заказы';

                closeReservationButton.style.display = 'inline-block';
                closeReservationButton.style.visibility = 'visible';
                closeReservationButton.style.position = 'relative';

                closeReservationButton.addEventListener('click', () => {
                    tableContainer.style.display = 'none';
                    closeReservationButton.remove();
                    button.style.display = 'inline-block';
                });

                button.insertAdjacentElement('afterend', closeReservationButton);
            } else {
                console.log('Кнопка "Скрыть заказы" уже существует.');
            }

            button.style.display = 'none';
        })
        .catch(error => {
            console.error('Ошибка при получении данных бронирования:', error);
            alert('Не удалось загрузить бронирования. Пожалуйста, попробуйте снова.');
        });
}

function deleteOrder(cassetteID, userID) {
    const totalCountRow = document.getElementById("totalcount")
    const rentedCountRow = document.getElementById("rentedcount")

    const totalCount = parseInt(totalCountRow.textContent, 10)
    const rentedCount = parseInt(rentedCountRow.textContent, 10)

    console.log(totalCount)
    console.log(rentedCount)

    const dataCreateReservation = {
        user_id: parseInt(userID, 10),
        cassette_id: parseInt(cassetteID, 10),
    };

    fetch(`http://localhost:8081/admin/orders`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${getCookieByKey('access_token')}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataCreateReservation)
    })
        .then(response => {
            if (response.ok) {
                const row = document.querySelector(`tr[data-cassette-id="${cassetteID}"]`);
                console.log(row)
                if (row) {
                    row.remove();
                }

                totalCountRow.textContent = (totalCount - 1).toString();
                rentedCountRow.textContent = (rentedCount - 1).toString();


            } else {
                throw new Error('Ошибка при бронировании.');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
        })

}

function editCassette(element, inputType, row) {
    if (!element || !element.classList.contains('edit-icon')) {
        console.error('Некорректный элемент для редактирования');
        return;
    }

    let parent = element.closest('.cassette-info');
    if (!parent) {
        console.error('Некорректный родительский элемент');
        return;
    }

    let textElement;

    switch (row) {
        case 'Name':
            textElement = document.getElementById('cassettename');
            break;
        case 'Genre':
            textElement = document.getElementById('cassettegenre');
            break;
        case 'Total count':
            textElement = document.getElementById('totalcount');
            break;
        // case 'Remain':
        //     textElement = document.getElementById('remaincount');
        //     break;
        // case 'Order count':
        //     textElement = document.getElementById('rentedcount');
        //     break;
    }

    console.log(textElement);
    if (!textElement) {
        console.error('Текстовый элемент не найден');
        return;
    }

    const input = document.createElement('input');
    input.type = inputType === 'number' ? 'number' : 'text';
    input.value = textElement.textContent.trim();
    input.id = "change-params";

    if (inputType === 'number') {
        input.min = '0';
    }

    try {
        textElement.replaceWith(input);
    } catch (error) {
        console.error('Ошибка при замене элемента:', error);
        return;
    }

    let saveButton = document.getElementById('saveButton');
    if (!saveButton) {
        saveButton = document.createElement('button');
        saveButton.id = 'saveButton';
        saveButton.textContent = 'Сохранить изменения';
        saveButton.classList.add('btn', 'btn-success', 'btn-sm');
        saveButton.style.marginTop = '10px';
        saveButton.style.marginLeft = '10px';
        saveButton.style.display = 'inline-block';

        saveButton.addEventListener('click', myFunction);

        const buttonsContainer = document.getElementById('reservationsTableContainer').parentNode;
        buttonsContainer.appendChild(saveButton);
    }

    input.focus();

    element.style.display = 'none';
}

function myFunction(event) {
    const button = event.target;

    let elems = document.getElementsByClassName("info-row");

    for (let elem of elems) {
        console.log(elem);

        let input = elem.querySelector('input');
        if (input) {
            let inputToSpan = document.createElement('span');
            inputToSpan.id = 'input-changes';
            inputToSpan.className = 'cassette-text';
            inputToSpan.textContent = input.value;

            input.parentNode.replaceChild(inputToSpan, input);
        }
    }

    let name = "";
    let ganre = "";
    let totalCount = 0;
    let rentedCount = 0;

    for (let elem of elems) {
        let strong = elem.querySelector('strong');
        let span = elem.querySelector('span');

        switch (strong.textContent) {
            case "Genre:":
                ganre = span.textContent;
                break;
            case "Name:":
                name = span.textContent;
                break;
            case "Total count:":
                totalCount = span.textContent;
                break;
            case "Remain:":
                remainCount = span.textContent;
                break;
        }
    }

    let id = document.getElementById('cassetteid')

    const data = {
        cassetteID: parseInt(id.textContent, 10),
        name: name,
        ganre: ganre,
        totalCount: parseInt(totalCount, 10),
        remain: parseInt(rentedCount, 10),
    };

    fetch('http://localhost:8081/cassette', {
        method: 'PUT',
        headers: {
            'Authorization': `Bearer ${getCookieByKey('access_token')}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok.');
            }
            return response.json();
        })
        .then(data => {
            window.location.reload();
        })
        .catch((error) => {
            console.error('Ошибка:', error);
        });

    button.remove();
}
й

function handleDeleteCassette(event) {
    const button = event.target;

    if (button.hasAttribute('data-request-sent') && button.getAttribute('data-request-sent') === 'true') {
        console.log('Запрос на удаление уже отправлен');
        return;
    }

    button.setAttribute('data-request-sent', 'true');

    const element = document.getElementById("cassetteid");
    if (!element) {
        console.error("Элемент с id 'cassetteid' не найден.");
        return;
    }

    const cassetteID = element.textContent;

    fetch(`http://localhost:8081/cassette/` + cassetteID, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${getCookieByKey('access_token')}`,
            'Content-Type': 'application/json'
        },
    })
        .then(response => {
            if (response.ok) {
                location.reload();
            } else {
                throw new Error('Ошибка при удалении кассеты.');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
        })
        .finally(() => {
            button.removeAttribute('data-request-sent');
        });
}

function handleDeleteReservationClick(event) {
    const button = event.target;
    const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10);

    const cassette_id = button.getAttribute('data-cassette-id');

    const dataCreateReservation = {
        cassette_id: parseInt(cassette_id, 10),
    };

    const AddReservationButton = `
    <button id="reserveButton" type="button" class="btn btn-primary btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${cassetteId}">Забронировать</button>
    `;

    fetch(`http://localhost:8081/reservations`, {
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
    const button = event.target;
    const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10);

    const user_id = document.querySelector('meta[name="user_id"]').content;
    const cassette_id = button.getAttribute('data-cassette-id');
    const dataCreateReservation = {
        user_id: parseInt(user_id, 10),
        cassette_id: parseInt(cassette_id, 10),
    };

    const deleteReservationButton = `
    <button id="deleteReservationButton" type="button" class="btn btn-danger btn-sm" style="display: inline-block; margin-left: 10px;" data-cassette-id="${cassetteId}">Отменить бронирование</button>
    `;

    if (button.hasAttribute('data-loading')) {
        return;
    }

    button.setAttribute('data-loading', true);
    button.disabled = true;

    fetch(`http://localhost:8081/reservations`, {
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
            button.removeAttribute('data-loading');
            button.disabled = false;
        });
}


function deleteReservation(cassetteID) {
    const dataCreateReservation = {
        cassette_id: parseInt(cassetteID, 10),
    };

    fetch(`http://localhost:8081/reservations`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${getCookieByKey('access_token')}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataCreateReservation)
    })
        .then(response => {
            if (response.ok) {
                location.reload();
            } else {
                throw new Error('Ошибка при бронировании.');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
        })
        .finally(() => {
            button.removeAttribute('data-loading');
            button.disabled = false;
        });
}