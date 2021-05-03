var bootstrapToast;
jQuery(function(){
    initial();
});

function initial() {
    let toastUI = $('#upload-status-toast').get()[0];
    let option = {
        animation: true,
        autohide: true,
        delay: 1000
    };
    bootstrapToast = new bootstrap.Toast(toastUI, option);

    $('#paint-file-input').change(function(){
        previewLocalImage(this, $('#paint-preview img'));
    });

    $('#paint-add-btn').on('click', uploadPaintImage);
}

function previewLocalImage(input, imageUI) {
    if (input.files && input.files[0]) {
        var reader = new FileReader();

        reader.onload = function(e) {
            imageUI.attr('src', e.target.result);
        };
        
        reader.readAsDataURL(input.files[0]);
    }
}

function uploadPaintImage() {
    let addPaintButton = $('#paint-add-btn');
    addPaintButton.prop('disabled', true);

    let paintName = $('#paint-name-input').val();
    let paintFile = $('#paint-file-input')[0].files[0];
    
    let paintFormData = new FormData();
    paintFormData.append('name', paintName);
    paintFormData.append('image', paintFile);
    
    let content = {
        'url': '/api/paint',
        'type': 'POST',
        'cache': false,
        'contentType': false,
        'processData': false,
        'data': paintFormData
    };
    
    $.ajax(content).done(function (response) {        
        if (response.status == 'ok') {
            $('#upload-status-toast div.toast-body').text('Upload paint success');
            bootstrapToast.show();
        } else {
            $('#upload-status-toast div.toast-body').text('Upload paint failed');
            bootstrapToast.show();
        }

        setTimeout(function(){
            $(location).prop('href', 'paint.html');
        }, 1500);
    })
    .fail(function (response) {
        addPaintButton.prop('disabled', false);
        $('#upload-status-toast div.toast-body').text('Upload paint failed');
        bootstrapToast.show();
        console.log('api_post_paint: Fail ' + response.responseText);
    });
}