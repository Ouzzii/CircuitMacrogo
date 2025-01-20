
$('body').on('click', '.compileButton', function(){
    let compilePromises = [];

    $('.willcheck:checked').each(function() {
        const target = $(this).parent().find('.compileas').val()
        const path = $(this).parent().attr('dir')
        const pdfpath = path.substr(0, path.lastIndexOf(".") < 0 ? path.length : path.lastIndexOf(".")) + ".pdf";
            console.log(pdfpath)
            let compilePromise =fetch("/Compile", {
            method: "POST",
            body: JSON.stringify({
                target: target,
                path: path
            })
        }).then(function(err) {

            console.log($(this).parent())
                if (err != "") {
                    window.RefreshSinglePreview(pdfpath)
                    //generateNotification('error', 'Bir hata ile karşılaşıldı', err);
                }else{
                    
                    window.RefreshSinglePreview(pdfpath)
                    //generateNotification('success', 'Başarıyla derlendi', 'Pdf Onizlemesi yenileniyor..');
                }
            });
        compilePromises.push(compilePromise);
        generateNotification('info', 'Derleniyor', `${$(this).parent().attr('dir')} derleniyor...`);
    });


    Promise.all(compilePromises).then(function() {

        //window.RefreshAllPreviews(Array.from($(".previewtab")));
    
    });
})