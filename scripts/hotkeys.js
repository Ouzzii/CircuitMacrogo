import { SaveContent } from "../../wailsjs/go/backend/App"


document.addEventListener('keydown', function(event){
    console.log("key pressed")
    if (event.ctrlKey && event.key === 's'){
        
        event.preventDefault()

        const path = $('.editTextArea#active').attr('dir')
        const content = $('.editTextArea#active').val()
        SaveContent(path, content).then(function(result){
            if (result == true){
                //generateNotification("success", "Dosya Kaydedildi", `Dosya başarıyla kaydedildi.\nKaydedilen Dosya: ${path}`)
            }

        })

    }
})