document.addEventListener('DOMContentLoaded', function () {
    const codeInput = document.getElementById('code');
    document.getElementById('uploadCodeBtn').addEventListener('click', () => codeInput.click());
  
    codeInput.addEventListener('change', function () {
        if (!this.files[0]) return;
        
        const formData = new FormData();
        formData.append('code', this.files[0]);
  
        fetch('/upload/code', {
            method: 'POST',
            body: formData,
        })
        .then(res => res.text())
        .then(data => {
            console.log("Файл загружен:", data);
        })
        .catch(err => console.error("Ошибка:", err));
    });
  
    const eventSource = new EventSource('/upload/code');
  
    eventSource.addEventListener('message', function(e) {
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
  
    eventSource.onerror = function(err) {
        console.error("SSE ошибка:", err);
    };
  });
  