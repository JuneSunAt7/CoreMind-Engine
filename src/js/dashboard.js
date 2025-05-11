$(document).ready(function () {
    // upd txt
    $('#modelFile').on('change', function () {
      var fileName = $(this).val().split('\\').pop();
      if (fileName) {
        $('.custom-file-upload').text('' + fileName);
      } else {
        $('.custom-file-upload').text('Выберите файл модели');
      }
    });
  
    $('form:eq(0)').on('submit', function (e) {
      if ($('#model')[0].files.length === 0) {
        e.preventDefault();
        $('#model').click(); // emu clc
      }
    });

    $('form:eq(1)').on('submit', function (e) {
      if ($('#dataset')[0].files.length === 0) {
        e.preventDefault();
        $('#dataset').click();
      }
    });
  });