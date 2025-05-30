
  document.addEventListener('DOMContentLoaded', function () {
    const modelSelect = document.getElementById('dataset');

    fetch('/api/dataset')
      .then(response => response.json())
      .then(models => {
        if (models.length === 0) {
          const option = document.createElement('option');
          option.value = '';
          option.text = 'Нет загруженных датасетовй';
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
        console.error("Ошибка загрузки датасетов:", err);
        const option = document.createElement('option');
        option.value = '';
        option.text = 'Ошибка загрузки';
        option.disabled = true;
        option.selected = true;
        modelSelect.appendChild(option);
      });
  });
