import { GetPDF } from "../../wailsjs/go/backend/App"

const zoomLevels = {};

$('body').on('click','.file', function(){ // dosya .pdf ile bitiyorsa preview olarak aç
    if ($(this).text().endsWith('.pdf') && $(`.previewtab[dir="${$(this).attr('dir')}"]`).length == 0){
        console.log($(this).attr('dir'))
        createpreviewTab($(this).attr('dir'))
        previewPdf($(this).attr('dir'))
        createPDFButtons($(this).attr('dir'))
        updateZoomLevelDisplay()
    }
})


$('body').on('click', '.previewtab', function(){ // tablar arası geçiş
    $('.previewtab#active').removeAttr('id')
    $(this).attr('id', 'active')
    changeCanvas()
    updateZoomLevelDisplay()
})

function changeCanvas(){ // tablar arasında geçiş yapılırken canvasları gizler/ortaya çıkarır
    const activeTabPath = $('.previewtab#active').attr('dir')
    $('.pdf-render#active').removeAttr('id')
    $(`.pdf-render[dir="${activeTabPath}"]`).attr('id', 'active')
}

$('body').on('click', '.previewtabs .closeButton', function(event){
    event.stopPropagation();
    closeTab($(this))
})

function closeTab(button){ // Sekme kapama olayı bu fonksiyonda gerçekleşir
    var tab = button.parent()[0]
    var tabText = $(`canvas[dir="${$(tab).attr('dir')}"]`)
    if ($(tab).attr('id') == 'active'){
    


        var currentIndex = $('.previewtabs div.previewtab').toArray().indexOf(tab);

        // Sonraki tab'ı bul
        const nextTab = $('.previewtabs .previewtab').eq(currentIndex + 1);
        // Önceki tab'ı bul
        const prevTab = $('.previewtabs .previewtab').eq(currentIndex - 1);


        console.log(nextTab)
        console.log(prevTab)
        if (nextTab.length){
            nextTab.attr('id', 'active')
            $(`canvas[dir="${nextTab.attr('dir')}"]`).attr('id', 'active')
            tab.remove()
            tabText.remove()
        }else if (prevTab.length){
            prevTab.attr('id', 'active')
            $(`canvas[dir="${prevTab.attr('dir')}"]`).attr('id', 'active')
            tab.remove()
            tabText.remove()
        }
        console.log(nextTab.length)
        console.log(prevTab.length)
    }else{
        $(`canvas[dir="${$(tab).attr('dir')}"]`).remove()
        console.log($(`canvas[dir="${$(tab).attr('dir')}"]`))
        $(tab).remove()

    }
}


function createpreviewTab(path){ // Bir dosyaya tıklandığında o dosyaya uygun canvas ve tab oluşturur
    if ($(`.previewtab[dir="${path}"]`).length != 0) return
    var previewTab = jQuery('<div>', {
        dir: path,
        class: 'previewtab',
        
    })
    var editorTitle = jQuery('<a>', {
        text: path.split('/')[path.split('/').length-1]
    })
    var editorTextArea = jQuery('<canvas>', {
        class: 'pdf-render',
        dir: path
    })

    var closeButton = jQuery('<img>', {
        class: 'closeButton',
        src: 'https://cdn-icons-png.flaticon.com/512/9068/9068699.png'
    })
    if ($('.previewtab').length == 0) {
        previewTab.attr('id', 'active')
        editorTextArea.attr('id', 'active')
    }




    previewTab.append(editorTitle)
    previewTab.append(closeButton)
    $('.previewtabs').append(previewTab)
    $('.previewCanvas').append(editorTextArea)

}

function previewPdf(path){ // canvas oluşturulduktan sonra içerisine pdf verisini işler

    zoomLevels[path] = 1.0; // Varsayılan zoom seviyesi
    renderPdfPage(path, zoomLevels[path]);

}
function renderPdfPage(path, scale) {
    GetPDF(path).then(function(base64Data) {
        if (base64Data) {
            const pdfData = atob(base64Data); // Base64'ten binary'ye çevir

            // PDF.js ile PDF verisini yükleyin
            const loadingTask = pdfjsLib.getDocument({ data: pdfData });
            loadingTask.promise.then(function(pdf) {
                pdf.getPage(1).then(function(page) {
                    const viewport = page.getViewport({ scale: scale });
                    const canvas = $(`canvas[dir="${path}"]`)[0];
                    const ctx = canvas.getContext('2d');

                    canvas.height = viewport.height;
                    canvas.width = viewport.width;

                    const renderContext = {
                        canvasContext: ctx,
                        viewport: viewport,
                    };
                    page.render(renderContext);
                });
            }).catch(function(error) {
                console.error("PDF Yüklenemedi:", error);
            });
        } else {
            console.error("PDF verisi alınamadı");
        }
    });
}



function createPDFButtons(path){
    var zoomOutButton = jQuery('<button>',{
        text: '+',
        class: 'zoomOut',
        path: path
    })

    var zoomLevel = jQuery('<a>', {
        text: '',
        class: 'zoom-level'
    })

    var zoomInButton = jQuery('<button>', {
        text: '-',
        class: 'zoomIn',
        path: path
    })
    if ($('.zoomOut').length == 0){
    $('.previewSettings').append(zoomInButton)

    $('.previewSettings').append(zoomLevel)
    $('.previewSettings').append(zoomOutButton)
}




}


function RefreshAllPreviews(paths){


    for (let path of paths) {

        closeTab($(`.previewtab[dir="${path.getAttribute("dir")}"] button`))
        createpreviewTab(path.getAttribute("dir"))
        previewPdf(path.getAttribute("dir"))
    }


}

$('body').on('click', '.zoomOut', function() {
    const path = $('.previewtab#active').attr('dir'); // Hangi PDF için zoom yapıldığını al
    zoomLevels[path] = (zoomLevels[path] || 1.0) + 0.1; // Zoom seviyesini artır
    console.log(zoomLevels)
    renderPdfPage(path, zoomLevels[path]);
    updateZoomLevelDisplay()
});

$('body').on('click', '.zoomIn', function() {
    const path = $('.previewtab#active').attr('dir'); // Hangi PDF için zoom yapıldığını al
    zoomLevels[path] = Math.max((zoomLevels[path] || 1.0) - 0.1, 0.5); // Zoom seviyesini azalt, 0.5'ten küçük olmasın
    renderPdfPage(path, zoomLevels[path]);
    updateZoomLevelDisplay()
});

function updateZoomLevelDisplay() {
    const zoomLevel = zoomLevels[$('.previewtab#active').attr('dir')] || 1.0; // Varsayılan zoom seviyesi 1.0
    const percentage = Math.round(zoomLevel * 100); // Yüzdeye çevir
    $(`.zoom-level`).text(`${percentage}%`); // İlgili span'a yazdır
}



window.RefreshAllPreviews = RefreshAllPreviews