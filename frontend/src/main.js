import './style.css';
import './app.css';
import './style/mainpage.css';
import './style/tree.css';
import './style/editor.css';
import './style/preview.css';
import './style/notification.css';
import './style/loadingSpinner.css';
import './style/distroSelection.css';

import { Detect_tex_distros, Boxdims_is_installed, ChooseDistro } from '../wailsjs/go/backend/App';
import { backend } from '../wailsjs/go/models';



function Init() {
    console.log('creating loading div')
    var loading = jQuery('<div>', {
        class: 'loading',
    })
    loading.css({
        'height': '100vh',
        'width': '100vw',
        'position': "absolute",
        'top': '0px',
        'backgroundColor': 'rgba(27, 38, 54, 1)',
        'display': 'flex',
        'flexDirection': 'column'
    })
    var loadingSpinner = `
    <div class="loadingspinner">
      <div id="square1"></div>
      <div id="square2"></div>
      <div id="square3"></div>
      <div id="square4"></div>
      <div id="square5"></div>
    </div>
    <a>LaTeX Kütüphaneleri kontrol ediliyor</a>
  `;
    loading.append(loadingSpinner)
    $('body').append(loading)



    Detect_tex_distros().then(function (configuration) {
        if (configuration['last-distro'] == "") {
            Boxdims_is_installed().then(function (boxdims) {
                loading.remove()
                distroSelect(boxdims, configuration)
            })
        }else{
            Boxdims_is_installed()
            loading.remove()
            generateNotification('success', 'Tex dağıtımı seçildi', `Cihazınızda daha önceden ${configuration['last-distro']} seçildiğinden tekrardan bu dosya yolundaki dağıtım seçilmiştir`)

        }
        
    })



}

function distroSelect(boxdims, config) {
    var distroSelectPopup = jQuery('<div/>', {
        id: 'distroSelectPopup',



    }).css({
        'height': '200px',
        'width': '50%',
        'position': 'absolute',
        'backgroundColor': 'black',
        'left': '50%',
        'top': '50%',
        'transform': 'translate(-50%, -50%)'
    })
    var selectionArea = jQuery('<div/>', { class: 'selectionArea' })
    
    var question = jQuery('<a/>', {
        class: 'distroSelectQuestion',
        text: 'Cihazınızda aşağıdaki dağıtımlar bulundu, devam etmek için birini seçmelisiniz.'
    })




    if (Object.keys(config.pdflatexPaths).length > 1) {
        Object.keys(config.pdflatexPaths).forEach(key => {
            if (boxdims[key] == "exist") {
                console.log(key, config.pdflatexPaths[key]);
                var selection = jQuery('<a/>', { class: 'distroSelection dexist', text: `${key}, boxdims kurulum durumu: kurulu` })
            } else {
                var selection = jQuery('<a/>', { class: 'distroSelection dnotexist', text: `${key}, boxdims kurulum durumu: kurulu değil` })

            }
            selectionArea.append(selection)
        });
        distroSelectPopup.append(question)
        distroSelectPopup.append(selectionArea)
        $('body').append(distroSelectPopup)
    }else if (Object.keys(config.pdflatexPaths).length == 1){
        Object.keys(config.pdflatexPaths).forEach(key => {
            if (boxdims[key] == "exist") {
                console.log(key, config.pdflatexPaths[key]);
                ChooseDistro(key).then(function(){})
                generateNotification('success', 'Tex dağıtımı seçildi', `Cihazınızda yalnızca tek bir dağıtım bulunduğundan otomatik olarak ${key} seçilmiştir`)
            } else {
                generateNotification('error', 'Tex dağıtımı hatası', `Cihazınızda ${key} tespit edildi fakat ${key} içerisinde boxdims kütüphanesi bulunmadığından herhangi bir tex dağıtımı seçilememektedir.`)

            }
        });
    }



}

$('body').on('click', '.dexist', function(){
    const key = $(this).text().split(",")[0]
    ChooseDistro(key).then(function(){})
    generateNotification('success', 'Tex dağıtımı seçildi', `Cihazınızda yalnızca tek bir dağıtım bulunduğundan otomatik olarak ${key} seçilmiştir`)
    $('#distroSelectPopup').remove()
})

Init()