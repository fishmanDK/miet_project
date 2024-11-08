document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('cassetteForm').addEventListener('submit', function (event) {
        event.preventDefault(); // Предотвращаем отправку формы по умолчанию

        // Собираем данные из формы
        const name = document.getElementById('cassetteName').value;
        const genre = document.getElementById('cassetteGenre').value;
        const year = document.getElementById('yearOfRelease').value;
        const storeID = document.querySelector('meta[name="store_id"]').content; // получаем значение из мета-тега
        const totalCount = document.getElementById('totalCount').value;

        // Создаем объект данных
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
                'Authorization': `Bearer ${accessToken = getCookieByKey('access_token')}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok.');
        })
        .then(data => {
            console.log('Received new cassette:', data);
            const { id } = data;
            // Добавляем новую строку в таблицу
            const cassetteList = document.querySelector('table');
            if (!cassetteList) {
                console.error("Не найдена таблица для добавления строки");
                return;
            }

            const newRow = document.createElement('tr');
            
            // Добавляем ячейки с полученными данными
            newRow.innerHTML = `
                <td class="cassette-item" data-id="{{$index}}" data-type="id">${id}</td>
                <td class="cassette-item" data-id="{{$index}}" data-type="name">${name}</td>
                <td class="cassette-item" data-id="{{$index}}" data-type="genre">${genre}</td>
                <td class="cassette-item" data-id="{{$index}}" data-type="year">${year}</td>
            `;
            
            cassetteList.appendChild(newRow);

            // Очищаем поля формы после добавления кассеты
            document.getElementById('cassetteForm').reset();
        })
        .catch((error) => {
            console.error('Ошибка:', error);
        });
    });
});
