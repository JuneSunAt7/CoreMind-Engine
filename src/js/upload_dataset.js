document.addEventListener('DOMContentLoaded', function () {
const datasetInput = document.getElementById('dataset');
    document.getElementById('uploadDatasetBtn').addEventListener('click', () => datasetInput.click());
    console.log("Кнопка загрузки датасета нажата");
    datasetInput.addEventListener('change', function () {
        if (!this.files[0]){
            console.log("Файл не выбран");
            return;
        } 
  
        const formData = new FormData();
        formData.append('dataset', this.files[0]);
        
  
        fetch('/upload/dataset', {
            method: 'POST',
            body: formData,
        })
        .then(res => res.text())
        .then(data => {
            console.log("Файл загружен:", data);
        })
        .catch(err => console.error("Ошибка:", err));
    });
  
    const eventDataSource = new EventSource('/upload/dataset');
  
    eventDataSource.addEventListener('message', function(e) {
        const percent = parseInt(e.data);
        const progressBar = document.getElementById('globalProgressFill');
        const progressText = document.getElementById('progressPercent');
  
        if (!isNaN(percent) && progressBar && progressText) {
            progressBar.style.width = percent + '%';
            progressText.textContent = percent + '%';
  
            if (percent === 100) {
                setTimeout(() => {
                    alert('Загрузка завершена!');
                    progressBar.style.width = '0%';
                    progressText.textContent = '0%';
                }, 500);
            }
        }
    });
  
    eventDataSource.onerror = function(err) {
        console.error("SSE ошибка:", err);
    };
});