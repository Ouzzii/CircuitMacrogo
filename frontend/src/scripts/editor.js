$('body').on('click','.file', function(){
    createTab($(this).attr('dir'))
})

$('body').on('click', '.filetab', function(){
    $('.filetab#active').removeAttr('id')
    $(this).attr('id', 'active')
})


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
    editorTab.append(editorTitle)
    $('.editortabs').append(editorTab)
}