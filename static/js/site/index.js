// document.addEventListener('DOMContentLoaded', function () {
//     // Обработчик событий для кликов по всему body
//     document.body.addEventListener('click', function (event) {
//         // Делегирование событий для разных элементов
//         if (event.target.classList.contains('cassette-item')) {
//             event.stopPropagation();
//             handleCassetteClick(event);
//         }

//         if (event.target.id === 'orderButton') {
//             event.preventDefault();
//             event.stopPropagation();
//             handleOrderClick(event);
//         }

//         if (event.target.id === 'show_reservations_button') {
//             event.stopPropagation();
//             event.stopImmediatePropagation();
//             showReservations(event);
//         }

//         if (event.target.id === 'close_reservations_view_button') {
//             event.stopPropagation();
//             event.stopImmediatePropagation();
//             closeViewReservations(event);
//         }

//         if (event.target.id === 'show_form_add_cassette_button') {
//             event.stopPropagation();
//             showFormAddCassette();
//         }

//         if (event.target.id === 'reserveButton') {
//             event.preventDefault();
//             event.stopPropagation();
//             handleReserveClick(event);
//         }

//         if (event.target.id === 'deleteButton') {
//             event.preventDefault();
//             event.stopPropagation();
//             handleDeleteCassette(event);
//         }

//         if (event.target.id === 'deleteReservationButton') {
//             event.preventDefault();
//             event.stopPropagation();
//             handleDeleteReservationClick(event);
//         }
//     });
// });

// function handleOrderClick(event) {
//     const button = event.target;
    
//     // Проверка, был ли уже отправлен запрос
//     if (button.hasAttribute('data-request-sent') && button.getAttribute('data-request-sent') === 'true') {
//         console.log('Запрос уже отправлен');
//         return; // Если запрос был отправлен, ничего не делаем
//     }

//     // Устанавливаем флаг, что запрос отправляется
//     button.setAttribute('data-request-sent', 'true');
    
//     const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10);
//     const element = document.getElementById("cassetteid");
//     if (!element) {
//         console.error("Элемент с id 'cassetteid' не найден.");
//         return;
//     }

//     const cassetteID = element.textContent;
//     const storeID = document.querySelector('meta[name="store_id"]').content;
//     const userID = document.querySelector('meta[name="user_id"]').content;

//     const dataOrder = {
//         userId: parseInt(userID, 10),
//         cassetteId: parseInt(cassetteID, 10),
//         storeId: parseInt(storeID, 10)
//     };

//     fetch(`http://localhost:8081/orders`, {
//         method: 'POST',
//         headers: {
//             'Authorization': `Bearer ${getCookieByKey('access_token')}`,
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify(dataOrder),
//     })
//     .then(response => {
//         if (response.ok) {
//             updateCounters();

//             button.disabled = true;
//             console.log('Заказ успешно оформлен');
//         } else {
//             throw new Error('Ошибка при оформлении заказа.');
//         }
//     })
//     .catch(error => {
//         console.error('Ошибка:', error);
//         alert('Не удалось оформить заказ. Попробуйте снова.');
//     })
//     .finally(() => {
//         // Снимаем флаг после завершения запроса, чтобы позволить повторное нажатие, если необходимо
//         button.removeAttribute('data-request-sent');
//     });
// }

// function updateCounters() {
//     const totalCountElement = document.getElementById('total_count_value');
//     const rentedCountElement = document.getElementById('rented_count_value');

//     console.log(totalCountElement, rentedCountElement)
    
//     // Проверим, что элементы существуют
//     if (!totalCountElement || !rentedCountElement) {
//         console.error("Элементы для счетчиков не найдены");
//         return;
//     }

//     // Получаем текущие значения счетчиков
//     let totalCount = parseInt(totalCountElement.textContent.replace('Total count: ', ''), 10);
//     let rentedCount = parseInt(rentedCountElement.textContent.replace('Rented count: ', ''), 10);
    
//     // Проверка на корректность значений
//     if (isNaN(totalCount) || isNaN(rentedCount)) {
//         console.error("Некорректные значения счетчиков");
//         return;
//     }

//     console.log(`Старые значения - Total count: ${totalCount}, Rented count: ${rentedCount}`);

//     // Обновляем счетчики
//     totalCount -= 1;
//     rentedCount += 1;

//     // Отображаем обновленные значения
//     totalCountElement.textContent = `${totalCount}`;
//     rentedCountElement.textContent = `${rentedCount}`;

//     console.log(`Обновленные значения - Total count: ${totalCount}, Rented count: ${rentedCount}`);
// }

// function closeViewReservations(event) {
//     const button = event.target;
//     const ViewAllOrdersButton = `
//         <button id="show_reservations_button" type="button" class="btn btn-primary btn-sm">Показать все заказы</button>
//     `;
//     const buttonContainer = button.parentElement;
//     button.remove();
//     buttonContainer.insertAdjacentHTML('beforeend', ViewAllOrdersButton);

//     const cassetteInfo = document.getElementById('reservationsTableContainer');
//     cassetteInfo.remove();
// }

// function showFormAddCassette() {
//     const cassetteInfo = document.getElementById('formContainer');
//     if (document.getElementById('cassetteForm')) {
//         console.log("Форма уже существует, повторное создание не требуется.");
//         return;
//     }

//     const newHtml = `
//     <h3>Добавление новой кассеты</h3>
//     <form id="cassetteForm">
//         <div class="mb-3">
//             <label for="cassetteName" class="form-label">Name</label>
//             <input type="text" class="form-control" id="cassetteName">
//         </div>
//         <div class="mb-3">
//             <label for="cassetteGenre" class="form-label">Genre</label>
//             <input type="text" class="form-control" id="cassetteGenre">
//         </div>
//         <div class="mb-3">
//             <label for="yearOfRelease" class="form-label">Year of release</label>
//             <input type="number" class="form-control" id="yearOfRelease" min="1900" max="2100" step="1" placeholder="Enter year (e.g. 2024)">
//         </div>
//         <div class="mb-3">
//             <label for="totalCount" class="form-label">Total count</label>
//             <input type="number" class="form-control" id="totalCount" step="1">
//         </div>
//         <button type="submit" class="btn btn-primary">Добавить</button>
//     </form>
//     `;
//     cassetteInfo.innerHTML = newHtml;

//     document.getElementById('cassetteForm').addEventListener('submit', function (event) {
//         event.preventDefault();
//         const name = document.getElementById('cassetteName').value;
//         const genre = document.getElementById('cassetteGenre').value;
//         const year = document.getElementById('yearOfRelease').value;
//         const storeID = document.querySelector('meta[name="store_id"]').content;
//         const totalCount = document.getElementById('totalCount').value;

//         const data = {
//             name: name,
//             genre: genre,
//             year: parseInt(year, 10),
//             storeId: parseInt(storeID, 10),
//             totalCount: parseInt(totalCount, 10),
//         };

//         fetch('http://localhost:8081/cassette', {
//             method: 'POST',
//             headers: {
//                 'Authorization': `Bearer ${getCookieByKey('access_token')}`,
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify(data)
//         })
//         .then(response => {
//             if (!response.ok) {
//                 throw new Error('Network response was not ok.');
//             }
//             return response.json();
//         })
//         .then(data => {
//             const { id } = data;
//             const cassetteList = document.querySelector('table');
//             if (!cassetteList) {
//                 console.error("Не найдена таблица для добавления строки");
//                 return;
//             }

//             const newRow = document.createElement('tr');
//             newRow.innerHTML = `
//                 <td class="cassette-item" data-id="${id}" data-type="id">${id}</td>
//                 <td class="cassette-item" data-id="${id}" data-type="name">${name}</td>
//                 <td class="cassette-item" data-id="${id}" data-type="genre">${genre}</td>
//                 <td class="cassette-item" data-id="${id}" data-type="year">${year}</td>
//             `;
//             cassetteList.appendChild(newRow);
//             document.getElementById('cassetteForm').reset();
//         })
//         .catch((error) => {
//             console.error('Ошибка:', error);
//         });
//     });
// }

// function handleCassetteClick(event) {
//     const cassetteItems = document.querySelectorAll(`[data-id="${event.target.getAttribute('data-id')}"]`);
//     let id, name, genre, year;

//     cassetteItems.forEach(item => {
//         const type = item.getAttribute('data-type');
//         if (type === 'id') id = item.textContent;
//         if (type === 'name') name = item.textContent;
//         if (type === 'genre') genre = item.textContent;
//         if (type === 'year') year = item.textContent;
//     });

//     const cassetteInfo = document.getElementById('cassetteInfoContainer');
//     cassetteInfo.innerHTML = `
//         <h3>${name}</h3>
//         <p>Жанр: ${genre}</p>
//         <p>Год выпуска: ${year}</p>
//         <button id="reserveButton" class="btn btn-primary">Зарезервировать</button>
//         <button id="deleteButton" class="btn btn-danger">Удалить</button>
//     `;
// }

// function handleDeleteCassette(event) {
//     const cassetteId = parseInt(event.target.getAttribute('data-id'), 10);

//     fetch(`http://localhost:8081/cassette/${cassetteId}`, {
//         method: 'DELETE',
//         headers: {
//             'Authorization': `Bearer ${getCookieByKey('access_token')}`,
//         },
//     })
//     .then(response => {
//         if (response.ok) {
//             alert('Кассета удалена');
//             location.reload();
//         } else {
//             throw new Error('Ошибка при удалении кассеты');
//         }
//     })
//     .catch(error => {
//         console.error('Ошибка:', error);
//     });
// }

// function handleDeleteReservationClick(event) {
//     const reservationId = parseInt(event.target.getAttribute('data-id'), 10);

//     fetch(`http://localhost:8081/reservations/${reservationId}`, {
//         method: 'DELETE',
//         headers: {
//             'Authorization': `Bearer ${getCookieByKey('access_token')}`,
//         },
//     })
//     .then(response => {
//         if (response.ok) {
//             alert('Резервация удалена');
//             location.reload();
//         } else {
//             throw new Error('Ошибка при удалении резервации');
//         }
//     })
//     .catch(error => {
//         console.error('Ошибка:', error);
//     });
// }





// document.addEventListener('DOMContentLoaded', function() {
//     // Обработчик событий для кликов по всему body
//     document.body.addEventListener('click', function(event) {
//         // Делегирование событий для разных элементов
//         if (event.target.classList.contains('cassette-item')) {
//             event.stopPropagation();
//             handleCassetteClick(event);
//         }

//         if (event.target.id === 'orderButton') {
//             event.preventDefault();
//             event.stopPropagation();
//             handleOrderClick(event);
//         }

//         if (event.target.id === 'show_reservations_button') {
//             event.stopPropagation();
//             event.stopImmediatePropagation();
//             showReservations(event);
//         }

//         if (event.target.id === 'close_reservations_view_button') {
//             event.stopPropagation();
//             event.stopImmediatePropagation();
//             closeViewReservations(event);
//         }

//         if (event.target.id === 'show_form_add_cassette_button') {
//             event.stopPropagation();
//             showFormAddCassette();
//         }

//         if (event.target.id === 'reserveButton') {
//             event.preventDefault();
//             event.stopPropagation();
//             handleReserveClick(event);
//         }

//         if (event.target.id === 'deleteReservationButton') {
//             event.preventDefault();
//             event.stopPropagation();
//             handleDeleteReservationClick(event);
//         }
//     });
// });

// // function handleOrderClick(event) {
// //     const button = event.target;
// //     const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10);
// //     const element = document.getElementById("cassetteid");
// //     if (!element) {
// //         console.error("Элемент с id 'cassetteid' не найден.");
// //         return;
// //     }

// //     const cassetteID = element.textContent;
// //     const storeID = document.querySelector('meta[name="store_id"]').content;
// //     const userID = document.querySelector('meta[name="user_id"]').content;

// //     const dataOrder = {
// //         userId: parseInt(userID, 10),
// //         cassetteId: parseInt(cassetteID, 10),
// //         storeId: parseInt(storeID, 10)
// //     };

// //     fetch(`http://localhost:8081/orders`, {
// //         method: 'POST',
// //         headers: {
// //             'Authorization': `Bearer ${getCookieByKey('access_token')}`,
// //             'Content-Type': 'application/json',
// //         },
// //         body: JSON.stringify(dataOrder),
// //     })
// //     .then(response => {
// //         if (response.ok) {
// //             button.disabled = true;
// //         } else {
// //             throw new Error('Ошибка при оформлении заказа.');
// //         }
// //     })
// //     .catch(error => {
// //         console.error('Ошибка:', error);
// //         alert('Не удалось оформить заказ. Попробуйте снова.');
// //     });
// // }


// function handleOrderClick(event) {
//     const button = event.target;
    
//     // Проверка, был ли уже отправлен запрос
//     if (button.hasAttribute('data-request-sent') && button.getAttribute('data-request-sent') === 'true') {
//         console.log('Запрос уже отправлен');
//         return; // Если запрос был отправлен, ничего не делаем
//     }

//     // Устанавливаем флаг, что запрос отправляется
//     button.setAttribute('data-request-sent', 'true');
    
//     const cassetteId = parseInt(button.getAttribute('data-cassette-id'), 10);
//     const element = document.getElementById("cassetteid");
//     if (!element) {
//         console.error("Элемент с id 'cassetteid' не найден.");
//         return;
//     }

//     const cassetteID = element.textContent;
//     const storeID = document.querySelector('meta[name="store_id"]').content;
//     const userID = document.querySelector('meta[name="user_id"]').content;

//     const dataOrder = {
//         userId: parseInt(userID, 10),
//         cassetteId: parseInt(cassetteID, 10),
//         storeId: parseInt(storeID, 10)
//     };

//     fetch(`http://localhost:8081/orders`, {
//         method: 'POST',
//         headers: {
//             'Authorization': `Bearer ${getCookieByKey('access_token')}`,
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify(dataOrder),
//     })
//     .then(response => {
//         if (response.ok) {
//             updateCounters();

//             button.disabled = true;
//             console.log('Заказ успешно оформлен');
//         } else {
//             throw new Error('Ошибка при оформлении заказа.');
//         }
//     })
//     .catch(error => {
//         console.error('Ошибка:', error);
//         alert('Не удалось оформить заказ. Попробуйте снова.');
//     })
//     .finally(() => {
//         // Снимаем флаг после завершения запроса, чтобы позволить повторное нажатие, если необходимо
//         button.removeAttribute('data-request-sent');
//     });
// }

// function updateCounters() {
//     const totalCountElement = document.getElementById('totalCount');
//     const rentedCountElement = document.getElementById('rentedCount');
    
//     // Проверим, что элементы существуют
//     if (!totalCountElement || !rentedCountElement) {
//         console.error("Элементы для счетчиков не найдены");
//         return;
//     }

//     // Получаем текущие значения счетчиков
//     let totalCount = parseInt(totalCountElement.textContent.replace('Total count: ', ''), 10);
//     let rentedCount = parseInt(rentedCountElement.textContent.replace('Rented count: ', ''), 10);
    
//     // Проверка на корректность значений
//     if (isNaN(totalCount) || isNaN(rentedCount)) {
//         console.error("Некорректные значения счетчиков");
//         return;
//     }

//     console.log(`Старые значения - Total count: ${totalCount}, Rented count: ${rentedCount}`);

//     // Обновляем счетчики
//     totalCount -= 1;
//     rentedCount += 1;

//     // Отображаем обновленные значения
//     totalCountElement.textContent = `Total count: ${totalCount}`;
//     rentedCountElement.textContent = `Rented count: ${rentedCount}`;

//     console.log(`Обновленные значения - Total count: ${totalCount}, Rented count: ${rentedCount}`);
// }




// function closeViewReservations(event){
//     const button = event.target;
//     const ViewAllOrdersButton = `
//         <button id="show_reservations_button" type="button" class="btn btn-primary btn-sm">Показать все заказы</button>
//     `;
//     const buttonContainer = button.parentElement;
//     button.remove();
//     buttonContainer.insertAdjacentHTML('beforeend', ViewAllOrdersButton);

//     const cassetteInfo = document.getElementById('reservationsTableContainer');
//     cassetteInfo.remove();
// }

// function showFormAddCassette() {
//     const cassetteInfo = document.getElementById('formContainer');
//     if (document.getElementById('cassetteForm')) {
//         console.log("Форма уже существует, повторное создание не требуется.");
//         return;
//     }

//     const newHtml = `
//     <h3>Добавление новой кассеты</h3>
//     <form id="cassetteForm">
//         <div class="mb-3">
//             <label for="cassetteName" class="form-label">Name</label>
//             <input type="text" class="form-control" id="cassetteName">
//         </div>
//         <div class="mb-3">
//             <label for="cassetteGenre" class="form-label">Genre</label>
//             <input type="text" class="form-control" id="cassetteGenre">
//         </div>
//         <div class="mb-3">
//             <label for="yearOfRelease" class="form-label">Year of release</label>
//             <input type="number" class="form-control" id="yearOfRelease" min="1900" max="2100" step="1" placeholder="Enter year (e.g. 2024)">
//         </div>
//         <div class="mb-3">
//             <label for="totalCount" class="form-label">Total count</label>
//             <input type="number" class="form-control" id="totalCount" step="1">
//         </div>
//         <button type="submit" class="btn btn-primary">Добавить</button>
//     </form>
//     `;
//     cassetteInfo.innerHTML = newHtml;

//     document.getElementById('cassetteForm').addEventListener('submit', function (event) {
//         event.preventDefault();
//         const name = document.getElementById('cassetteName').value;
//         const genre = document.getElementById('cassetteGenre').value;
//         const year = document.getElementById('yearOfRelease').value;
//         const storeID = document.querySelector('meta[name="store_id"]').content;
//         const totalCount = document.getElementById('totalCount').value;

//         const data = {
//             name: name,
//             genre: genre,
//             year: parseInt(year, 10),
//             storeId: parseInt(storeID, 10),
//             totalCount: parseInt(totalCount, 10),
//         };

//         fetch('http://localhost:8081/cassette', {
//             method: 'POST',
//             headers: {
//                 'Authorization': `Bearer ${getCookieByKey('access_token')}`,
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify(data)
//         })
//         .then(response => {
//             if (!response.ok) {
//                 throw new Error('Network response was not ok.');
//             }
//             return response.json();
//         })
//         .then(data => {
//             const { id } = data;
//             const cassetteList = document.querySelector('table');
//             if (!cassetteList) {
//                 console.error("Не найдена таблица для добавления строки");
//                 return;
//             }

//             const newRow = document.createElement('tr');
//             newRow.innerHTML = `
//                 <td class="cassette-item" data-id="${id}" data-type="id">${id}</td>
//                 <td class="cassette-item" data-id="${id}" data-type="name">${name}</td>
//                 <td class="cassette-item" data-id="${id}" data-type="genre">${genre}</td>
//                 <td class="cassette-item" data-id="${id}" data-type="year">${year}</td>
//             `;
//             cassetteList.appendChild(newRow);
//             document.getElementById('cassetteForm').reset();
//         })
//         .catch((error) => {
//             console.error('Ошибка:', error);
//         });
//     });
// }
