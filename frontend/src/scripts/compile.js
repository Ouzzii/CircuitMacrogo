import { Compile } from "../../wailsjs/go/backend/App"

$('body').on('click', '.compileButton', function(){
    $('.willcheck:checked').each(function(){
        generateNotification('info', 'Derleniyor', `${$(this).parent().attr('dir')} derleniyor...`)
        Compile($(this).parent().find('.compileas').val() ,$(this).parent().attr('dir')).then(function(err){
            if (err != ""){
                generateNotification('error', 'Bir hata ile karşılaşıldı', err)
            }
            //console.log($(this).parent().find('.compileas').val())
        })
    })

})