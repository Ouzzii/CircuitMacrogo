import { Compile } from "../../wailsjs/go/backend/App"

$('body').on('click', '.compileButton', function(){
    let compilePromises = [];

    $('.willcheck:checked').each(function() {
        let compilePromise = Compile($(this).parent().find('.compileas').val(), $(this).parent().attr('dir'))
            .then(function(err) {
                if (err != "") {
                    //generateNotification('error', 'Bir hata ile karşılaşıldı', err);
                }
            });
        compilePromises.push(compilePromise);
        //generateNotification('info', 'Derleniyor', `${$(this).parent().attr('dir')} derleniyor...`);
    });


    Promise.all(compilePromises).then(function() {

        window.RefreshAllPreviews(Array.from($(".previewtab")));
    });
})