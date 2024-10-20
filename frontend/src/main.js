import './style.css';
import './app.css';
import './style/mainpage.css';
import './style/tree.css';
import './style/editor.css';
import './style/preview.css';
import './style/notification.css';
import './style/loadingSpinner.css';

import { Detect_tex_distros } from '../wailsjs/go/backend/App';



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
    Detect_tex_distros().then(function () {
        loading.remove()
    })
    

}
Init()