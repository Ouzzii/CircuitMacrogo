import { AskDirectory, GetDirectory, IsFile } from '../../wailsjs/go/main/App';



$("body").on("click", "#pickFolder", function(){
    AskDirectory().then(function(result){
        let workspaceName = result.split("\\")[result.split("\\").length-1]
        workspaceName = workspaceName.split("/")[workspaceName.split("/").length-1]

        $(".workspace").html(workspaceName)
        GetDirectory(result).then(function(files){
            files.forEach(function(item){
                replaceFiles(item, result)
            })
        })

    })
})


function replaceFiles(file, rootDir){
    var filetype
    IsFile(file).then(function(result){

    console.log(file.split(rootDir)[1], result)
    })

}





