import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import {Greet} from '../wailsjs/go/main/App';
import { AskDirectory } from '../wailsjs/go/main/App';


document.querySelector("button").addEventListener("click", function(){
    AskDirectory().then(function(result){
        const workspaceName = result.split("\\")[result.split("\\").length-1]
        document.querySelector(".workspace").innerHTML = workspaceName
    })

    
})





// Setup the greet function
window.greet = function () {
    // Get name
    let name = nameElement.value;

    // Check if the input is empty
    if (name === "") return;

    // Call App.Greet(name)
    try {
        Greet(name)
            .then((result) => {
                // Update result with data back from App.Greet()
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
