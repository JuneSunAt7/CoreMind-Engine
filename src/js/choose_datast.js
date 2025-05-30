document.addEventListener('DOMContentLoaded', function () {
    const datasetSelect = document.getElementById('datasetSelect');

    fetch('/api/datasets') 
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка сети при загрузке датасетов');
            }
            return response.json();
        })
        .then(datasets => {
            if (datasets.length === 0) {
                const option = document.createElement('option');
                option.value = '';
                option.text = 'Нет загруженных датасетов';
                option.disabled = true;
                option.selected = true;
                datasetSelect.appendChild(option);
            } else {
                datasets.forEach(dataset => {
                    const option = document.createElement('option');
                    option.value = dataset;
                    option.text = dataset;
                    datasetSelect.appendChild(option);
                });
            }
        })
        .catch(err => {
            console.error("Ошибка загрузки датасетов:", err);
            const option = document.createElement('option');
            option.value = '';
            option.text = 'Ошибка загрузки';
            option.disabled = true;
            option.selected = true;
            datasetSelect.appendChild(option);
        });
});