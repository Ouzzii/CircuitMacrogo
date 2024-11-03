import { GetContent } from "../../wailsjs/go/backend/App"

$('body').on('click','.file', function(){
    if (!$(this).text().endsWith('.pdf')){
        createTab($(this).attr('dir'))
    }
    
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

    $('#active.compile').removeAttr('id')
    $(`.compile[dir="${activeTabPath}"]`).attr('id', 'active')

}


function createTab(path){
    if ($(`.filetab[dir="${path}"]`).length != 0) return
    var editorTab = jQuery('<div>', {
        dir: path,
        class: 'filetab',
        
    })
    var editorTitle = jQuery('<div>', {
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

    var Types = {
        'latex': 'LaTeX',
        'pdf': 'PDF'
    }
    var targetType = $('<select/>', {class: 'compileas'});
    
    for(var val in Types) {
        $('<option/>', {value: val, text: Types[val]}).appendTo(targetType);
    }

    var compileFlag = {
        'pgf': 'PGF'
    }
    var compileType = $('<select/>', {class: 'compilewith'});
    for(var val in compileFlag) {
        $('<option/>', {value: val, text: compileFlag[val]}).appendTo(compileType);
    }
    var checkbox = $('<input/>', {
        type: 'checkbox',
        class: 'willcheck',
        dir: path
    })
    var compileButton = $('<button/>', {
        class: 'compileButton',
        text: 'Derle'
    })

    const filesettings = $('<div/>', {class: 'filesettings', dir: path})
    const compilediv = $('<div/>', {class: 'compile', dir: path})
    
    if ($('.filetab').length == 0) {
        editorTab.attr('id', 'active')
        editorTextArea.attr('id', 'active')
        filesettings.attr('id', 'active')
        compilediv.attr('id', 'active')
    }
    GetContent(path).then(function(content){
        editorTextArea.text(content)
    })




    editorTab.append(editorTitle)
    editorTab.append(closeButton)
    $('.editortabs').append(editorTab)
    $('.editArea').append(editorTextArea)
    $('.fileOperations').append(filesettings)
    $('.fileOperations').append(compilediv)
    compilediv.append(targetType).append(compileType).append(checkbox).append(compileButton)
    //$('.fileOperations .compile').append(compileType)
    //$('.fileOperations .compile').append(checkbox)
    //$('.fileOperations .compile').append(compileButton)
}

$('body').on('click', '.editortabs .closeButton', function(event){
    event.stopPropagation();
    closeTab($(this))
})

function closeTab(button){
    const tab = button.parent()[0]
    const tabText = $(`textarea[dir="${$(tab).attr('dir')}"]`)
    const tabCompile = $(`.compile[dir="${$(tab).attr('dir')}"]`)
    const tabSettings = $(`.filesettings[dir="${$(tab).attr('dir')}"]`)
    if ($(tab).attr('id') == 'active'){
        var currentIndex = $('.editortabs div.filetab').toArray().indexOf(tab);

        const nextTab = $('.editortabs .filetab').eq(currentIndex + 1);
        const prevTab = $('.editortabs .filetab').eq(currentIndex - 1);
        
        if (nextTab.length){
            nextTab.attr('id', 'active')
            $(`textarea[dir="${nextTab.attr('dir')}"]`).attr('id', 'active')
            tab.remove()
            tabText.remove()
            tabCompile.remove()
            tabSettings.remove()
        }else if (prevTab.length){
            prevTab.attr('id', 'active')
            $(`textarea[dir="${prevTab.attr('dir')}"]`).attr('id', 'active')
            tab.remove()
            tabText.remove()
            tabCompile.remove()
            tabSettings.remove()
        }
    }else{
        $(`textarea[dir="${$(tab).attr('dir')}"]`).remove()
        $(tab).remove()
        tabCompile.remove()
        tabSettings.remove()

    }
}
function checkScrollbarHeight() {
    var w = 0
    $('.editortabs .filetab').each(function(){
        w = w + $(this).width()
    })

    if ($('.editortabs').width() < w ){
        $(':root').css('--scrollbar-height', '0.2rem');
    }else{
        $(':root').css('--scrollbar-height', '0');
    }
}
