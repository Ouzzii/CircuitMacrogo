const CheckSocket = new WebSocket("ws://localhost:8080/RunDirectoryCheck");


CheckSocket.addEventListener("message", (event) => {

    if (event.data == "false"){
        $('.filesfolders div,.filesfolders a').each(function(item){
            if ($(this).attr('class') != 'askdirectory'){
                $(this).remove()
            }
        })
        console.log(event.data)
        fetch("/CheckWorkspace").then(response => response.json()).then(workspace => {
                
            console.log(workspace)
            if (workspace != ""){
                loadWorkspace(workspace)
            }
          })

    }
});


function checkInit(){

    fetch("/CheckWorkspace").then(response => response.json()).then(workspace => {
        if (workspace != ""){
            loadWorkspace(workspace)
            ToggleWorkspace()
        }
      })
}


checkInit()

$("body").on("click", "#pickFolder", function(){

    fetch("/AskDirectory").then(async function(workspace){
        res = await workspace.json()
        loadWorkspace(res)
        ToggleWorkspace()
    })
})

function loadWorkspace(rootDir) {
    let workspaceName = rootDir.split("\\")[rootDir.split("\\").length - 1];
    workspaceName = workspaceName.split("/")[workspaceName.split("/").length - 1];
    console.log(workspaceName)
    $(".workspace .workspaceName").html(workspaceName);
    $(".filesfolders").attr("dir", rootDir);

    fetch("/GetDirectory", {
        method: "POST",
        body: JSON.stringify(rootDir)
    }).then(async function(result){
        res = await result.json()
        console.log(res)

        for (let i = 0; i < res.files.length; i++){
            file = res.files[i];
            filetype = res.filetypes[i]
            await replaceFiles(file, rootDir, filetype)
        }



    })

    /*
    GetDirectory(result).then(async function (files) {
        for (const item of files) {
            await replaceFiles(item, result);
        }
    });
    */


    $(".askdirectory").css('display', 'none');
    //CheckSocket.send("CheckDirectory")
    if (window.DirectoryCheckInterval) {
        clearInterval(window.DirectoryCheckInterval);
    }

    window.DirectoryCheckInterval = setTimeout(function() {
        RunDirectoryCheck();

        window.DirectoryCheckInterval = setInterval(function() {
            RunDirectoryCheck();
        }, 2000);

    }, 5000);
}

function RunDirectoryCheck(){
    var files = [];
    var dirs = [];
    
    $(".file").each(function() {
        files.push($(this).attr("dir"));
    });
    
    $(".directoryName").each(function() {
        dirs.push($(this).attr("dir"));
    });
    
    files.sort();
    dirs.sort();
    

    CheckSocket.send(JSON.stringify([files, dirs]))


    //EventsEmit("RunCheckDirectory", { files: files, directories: dirs });

}

async function replaceFiles(file, rootDir, filetype) {
    //const filetype = await IsFile(file);
    /*const response = await (await fetch("/IsFile", {method: 'POST', body: JSON.stringify(file)}))
    console.log("Raw Response:", response);
    const filetype = await response.json()
    console.log(filetype)*/


    const parentDir = file.split(file.split("/")[file.split("/").length - 2])[0] + file.split("/")[file.split("/").length - 2];

    if (!filetype) {
        createDirectory(file, parentDir, rootDir);
    } else {
        var directory = jQuery('<div>', {
            dir: file,
            class: 'file',
        });
        directory.css('padding-left', '10px');
        directory.text(file.split('/')[file.split('/').length - 1]);
        $(`div[dir='${parentDir}']`).append(directory);
    }
}
function createDirectory(dirName , parentDir, rootDir){
    if ($(`div[dir='${dirName}']`).length == 0){
        const name = dirName.split('/')[dirName.split('/').length-1]
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
    clearTimeout(window.DirectoryCheckInterval)
    clearInterval(window.DirectoryCheckInterval)
    window.DirectoryCheckInterval = null
}