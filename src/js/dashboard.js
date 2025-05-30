document.addEventListener('DOMContentLoaded', function () {
    const modelInput = document.getElementById('model');
    const datasetInput = document.getElementById('dataset');
    const codeInput = document.getElementById('code');

    document.getElementById('uploadModelBtn').addEventListener('click', () => modelInput.click());
    document.getElementById('uploadDatasetBtn').addEventListener('click', () => datasetInput.click());
    document.getElementById('uploadCodeBtn').addEventListener('click', () => codeInput.click());

    function setupUploadHandler(inputElement, endpoint) {
        inputElement.addEventListener('change', function () {
            if (!this.files[0]) return;

            const formData = new FormData();
            formData.append(inputElement.name, this.files[0]);

            fetch('/upload/' + endpoint, {
                method: 'POST',
                body: formData,
            })
            .then(res => res.text())
            .then(data => {
                console.log("Файл загружен:", data);
            })
            .catch(err => console.error("Ошибка загрузки:", err));
        });
    }

    setupUploadHandler(modelInput, 'model');
    setupUploadHandler(datasetInput, 'dataset');
    setupUploadHandler(codeInput, 'code');
});