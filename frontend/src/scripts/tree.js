import { AskDirectory, GetDirectory, IsFile } from '../../wailsjs/go/backend/App';



$("body").on("click", "#pickFolder", function(){
    AskDirectory().then(function(result){
        let workspaceName = result.split("\\")[result.split("\\").length-1]
        workspaceName = result.split("/")[result.split("/").length-1]

        $(".workspace").html(workspaceName)
        $(".filesfolders").attr("dir", result)
        GetDirectory(result).then(function(files){
            files.forEach(function(item){
                replaceFiles(item, result)
            })
        })
        $(".askdirectory").remove() 

    })
})


function replaceFiles(file, rootDir){
    IsFile(file).then(function(filetype){
        const parentDir = file.split(file.split("/")[file.split("/").length-2])[0]+file.split("/")[file.split("/").length-2]
    //console.log(file.split(rootDir)[1], filetype)
    if (!filetype){

        createDirectory(file,  parentDir, rootDir)
    }else{

        var directory = jQuery('<div>',{
            dir: file
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
            class: 'directory'

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
    $(this).toggleClass('Uparrow')
})