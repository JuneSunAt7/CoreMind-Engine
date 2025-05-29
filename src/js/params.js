
  document.addEventListener('DOMContentLoaded', function () {
    const modelSelect = document.getElementById('modelSelect');

    fetch('/api/models')
      .then(response => response.json())
      .then(models => {
        if (models.length === 0) {
          const option = document.createElement('option');
          option.value = '';
          option.text = 'Нет загруженных моделей';
          option.disabled = true;
          option.selected = true;
          modelSelect.appendChild(option);
        } else {
          models.forEach(model => {
            const option = document.createElement('option');
            option.value = model;
            option.text = model;
            modelSelect.appendChild(option);
          });
        }
      })
      .catch(err => {
        console.error("Ошибка загрузки моделей:", err);
        const option = document.createElement('option');
        option.value = '';
        option.text = 'Ошибка загрузки моделей';
        option.disabled = true;
        option.selected = true;
        modelSelect.appendChild(option);
      });
  });
