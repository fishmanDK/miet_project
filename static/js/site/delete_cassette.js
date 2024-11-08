// document.getElementById('deleteButton').addEventListener('click', function () {
//     // Получаем выбранный элемент (например, строку таблицы или идентификатор кассеты)
//     const selectedCassetteRow = /* логика для получения выбранной строки или ID кассеты */

//     // Пример: если вы используете data-id для идентификации
//     if (selectedCassetteRow) {
//         // Удалите элемент из таблицы
//         selectedCassetteRow.remove();

//         // Дополнительно: вы можете сделать запрос на сервер для удаления кассеты из базы данных
//         // fetch(`http://localhost:8081/cassette/${selectedCassetteId}`, { method: 'DELETE' })
//         //     .then(response => {
//         //         if (response.ok) {
//         //             console.log('Кассета удалена');
//         //         } else {
//         //             console.error('Ошибка при удалении кассеты');
//         //         }
//         //     });
//     } else {
//         console.log('Выберите кассету для удаления');
//     }
// });