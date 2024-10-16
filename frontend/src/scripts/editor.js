import { GetContent } from "../../wailsjs/go/backend/App"

$('body').on('click','.file', function(){
    createTab($(this).attr('dir'))
    checkScrollbarHeight()
})

$('body').on('click', '.filetab', function(){
    $('.filetab#active').removeAttr('id')
    $(this).attr('id', 'active')
    changeEditor()
})

function changeEditor(){
    const activeTabPath = $('.filetab#active').attr('dir')
    $('.editTextArea#active').removeAttr('id')
    $(`.editTextArea[dir="${activeTabPath}"]`).attr('id', 'active')
}


function createTab(path){
    console.log($(`.filetab[dir="${path}"]`))
    if ($(`.filetab[dir="${path}"]`).length != 0) return
    var editorTab = jQuery('<div>', {
        dir: path,
        class: 'filetab',
        
    })
    var editorTitle = jQuery('<a>', {
        text: path.split('/')[path.split('/').length-1]
    })
    var editorTextArea = jQuery('<textarea>', {
        class: 'editTextArea',
        dir: path
    })

    var closeButton = jQuery('<button>', {
        class: 'closeButton',
        text: 'Close'
    })

    if ($('.filetab').length == 0) {
        editorTab.attr('id', 'active')
        editorTextArea.attr('id', 'active')
    }
    GetContent(path).then(function(content){
        editorTextArea.text(content)
    })

    editorTab.append(editorTitle)
    editorTab.append(closeButton)
    $('.editortabs').append(editorTab)
    $('.editArea').append(editorTextArea)

}

$('body').on('click', '.closeButton', function(event){
    event.stopPropagation();
    closeTab($(this))
})

function closeTab(button){
    var tab = button.parent()[0]
    var tabText = $(`textarea[dir="${$(tab).attr('dir')}"]`)
    if ($(tab).attr('id') == 'active'){
    


        //const nextTab = $('.editortabs .filetab').eq($('.editortabs div.filetab').toArray().indexOf(tab)+1)
        //const prevTab = $('.editortabs .filetab').eq($('.editortabs div.filetab').toArray().indexOf(tab)-1)
        var currentIndex = $('.editortabs div.filetab').toArray().indexOf(tab);

        // Sonraki tab'ı bul
        const nextTab = $('.editortabs .filetab').eq(currentIndex + 1);
        // Önceki tab'ı bul
        const prevTab = $('.editortabs .filetab').eq(currentIndex - 1);


        console.log(nextTab)
        console.log(prevTab)
        if (nextTab.length){
            nextTab.attr('id', 'active')
            $(`textarea[dir="${nextTab.attr('dir')}"]`).attr('id', 'active')
            tab.remove()
            tabText.remove()
        }else if (prevTab.length){
            prevTab.attr('id', 'active')
            $(`textarea[dir="${prevTab.attr('dir')}"]`).attr('id', 'active')
            tab.remove()
            tabText.remove()
        }
        console.log(nextTab.length)
        console.log(prevTab.length)
        
        //console.log(nextTab.length)
        
    }else{
        $(`textarea[dir="${$(tab).attr('dir')}"]`).remove()
        console.log($(`textarea[dir="${$(tab).attr('dir')}"]`))
        $(tab).remove()

    }
}
function checkScrollbarHeight() {
    var w = 0
    $('.editortabs .filetab').each(function(){
        w = w + $(this).width()
    })
    console.log(w)

    if ($('.editortabs').width() < w ){
        $(':root').css('--scrollbar-height', '0.2rem');
    }else{
        $(':root').css('--scrollbar-height', '0');
    }
}
