//import { SaveContent } from "../../wailsjs/go/backend/App"


document.addEventListener('keydown', function(event){
    console.log("key pressed")
    if (event.ctrlKey && event.key === 's'){
        
        event.preventDefault()

        const path = $('.editTextArea#active').attr('dir')
        const content = $('.editTextArea#active').val()


        fetch("/SaveContent", {
            method: "POST",
            body: JSON.stringify({
                path: path,
                content: content
            })
        }).then(response => response.text()).then(response=>{
            if (response == "ok"){
                console.log("Kaydedildi")
                generateNotification("success", "Dosya Kaydedildi", `Dosya başarıyla kaydedildi.\nKaydedilen Dosya: ${path}`)
            }
        })
    }
})