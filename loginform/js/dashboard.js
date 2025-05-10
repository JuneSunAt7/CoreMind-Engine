$(document).ready(function () {
    // Обновление текста кнопки при выборе файла
    $('#modelFile').on('change', function () {
      var fileName = $(this).val().split('\\').pop();
      if (fileName) {
        $('.custom-file-upload').text('' + fileName);
      } else {
        $('.custom-file-upload').text('Выберите файл модели');
      }
    });
  
    // Подстановка файла в форму загрузки модели
    $('form:eq(0)').on('submit', function (e) {
      if ($('#model')[0].files.length === 0) {
        e.preventDefault();
        $('#model').click(); // эмулируем клик для выбора файла
      }
    });
  
    // То же самое для датасета
    $('form:eq(1)').on('submit', function (e) {
      if ($('#dataset')[0].files.length === 0) {
        e.preventDefault();
        $('#dataset').click();
      }
    });
  });