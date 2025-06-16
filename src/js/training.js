document.addEventListener('DOMContentLoaded', function () {
    const trainingStarted = localStorage.getItem('trainingStarted');

    if (trainingStarted === 'true') {
        const progressBar = document.getElementById('trainingProgress');
        const progressText = document.getElementById('trainingProgressPercent');

        if (!progressBar || !progressText) {
            console.warn("Элементы для отображения прогресса обучения не найдены");
            return;
        }

        const eventSource = new EventSource('/stream/train');

        eventSource.addEventListener('message', function(e) {
            const percent = parseInt(e.data);

            if (!isNaN(percent) && progressBar && progressText) {
                progressBar.style.width = percent + '%';
                progressText.textContent = percent + '%';

                if (percent === 95) {
                    setTimeout(() => {
                        progressBar.style.width = '0%';
                        progressText.textContent = '0%';
                        localStorage.removeItem('trainingStarted'); // Сбрасываем флаг
                    }, 500);
                }
            }
        });

        eventSource.onerror = function(err) {
            console.error("Ошибка SSE при обучении:", err);
        };
    }
});