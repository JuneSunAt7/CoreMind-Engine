document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('trainForm');

    form.addEventListener('submit', function (e) {
        e.preventDefault();

        const formData = new FormData(form);
        const params = {
            model: formData.get('model'),
            dataset: formData.get('dataset'),
            epochs: parseInt(formData.get('epochs')),
            batchSize: parseInt(formData.get('batchSize')),
            optimizer: formData.get('optimizer'),
            learningRate: parseFloat(formData.get('learningRate')),
            useAugmentation: formData.get('useAugmentation') === 'on',
            shuffleData: formData.get('shuffleData') === 'on'
        };
        fetch('/start-training', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(params)
        })
        .then(res => {
            if (res.ok) {
                localStorage.setItem('trainingStarted', 'true');
                window.location.href = '/dashboard';
            } else {
                res.text().then(text => alert('Ошибка: ' + text));
            }
        })
        .catch(err => {
            console.error(err);
            alert('Не удалось запустить обучение');
        });
    });
});