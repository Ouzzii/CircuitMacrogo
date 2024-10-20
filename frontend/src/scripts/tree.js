import { Detect_tex_distros,AskDirectory, GetDirectory, IsFile, CheckWorkspace, CloseConfWorkspace } from '../../wailsjs/go/backend/App';


function checkInit(){

    CheckWorkspace().then(function(workspace){
        console.log(workspace)
        if (workspace != ""){
            loadWorkspace(workspace)
        }

    })
}
checkInit()
$("body").on("click", "#pickFolder", function(){
            AskDirectory().then(function(result){
                loadWorkspace(result)
            })
})


function loadWorkspace(result){
    let workspaceName = result.split("\\")[result.split("\\").length-1]
    workspaceName = workspaceName.split("/")[workspaceName.split("/").length-1]
    $(".workspace .workspaceName").html(workspaceName)
    $(".filesfolders").attr("dir", result)
    GetDirectory(result).then(function(files){
        files.forEach(function(item){
            replaceFiles(item, result)
        })
    })
    $(".askdirectory").css('display', 'none') 
    ToggleWorkspace()
}

function replaceFiles(file, rootDir){
    IsFile(file).then(function(filetype){
        const parentDir = file.split(file.split("/")[file.split("/").length-2])[0]+file.split("/")[file.split("/").length-2]
    //console.log(file.split(rootDir)[1], filetype)
    if (!filetype){

        createDirectory(file,  parentDir, rootDir)
    }else{

        var directory = jQuery('<div>',{
            dir: file,
            class: 'file',
        })
        directory.css('padding-left', '10px')
        directory.text(file.split('/')[file.split('/').length-1])
        $(`div[dir='${parentDir}']`).append(directory)
    }
    

    })

}

function createDirectory(dirName , parentDir, rootDir){
    if ($(`div[dir='${dirName}']`).length == 0){
        const name = dirName.split('/')[dirName.split('/').length-1]
        console.log(dirName, parentDir, rootDir)
        var directory = jQuery('<div>',{
            dir: dirName,
            class: 'directory Closed'

        })
        var directoryText = jQuery('<a>',{
            dir: dirName,
            class: 'directoryName'
        })

        directoryText.css('padding-left', '10px')
        directory.css('padding-left', '10px')
        //directory.text(sname)
        directoryText.text(name)
        $(`div[dir='${parentDir}']`).append(directoryText)
        $(`div[dir='${parentDir}']`).append(directory)
    }
}






$('body').on('click', '.filesfolders a', function(){
    const dirName = $(this).attr('dir')
    $(`div[dir='${dirName}']`).toggleClass('Closed')
    $(this).toggleClass('Downarrow')
})





function ToggleWorkspace(){
    if ($('.workspace').css('visibility') == 'visible'){
        $('.workspace').css('visibility', 'hidden');
    }else{
        $('.workspace').css('visibility', 'visible');
    }

}


$('body').on('click', '.closeWorkspace', CloseWorkspace)

function CloseWorkspace(){
    $('.workspaceName').html('')
    $('.filesfolders').removeAttr('dir');

    $('.filesfolders div,.filesfolders a').each(function(item){
        if ($(this).attr('class') != 'askdirectory'){
            $(this).remove()
        }
    })

    ToggleWorkspace()
    $(".askdirectory").css('display', 'flex')
    CloseConfWorkspace()
}


