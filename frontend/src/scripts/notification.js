const notificationItems = [];

/*class SoundManager {
    constructor() {
        this.audioIn = new Audio('https://dsp-studio.ro/slide-sound.wav');
        this.audioOut = new Audio('https://dsp-studio.ro/fast-swipe.mp3');
    }

    playOpenSound() {
        this.playSound(this.audioIn);
    }

    playCloseSound() {
        this.playSound(this.audioOut);
    }

    playSound(audio) {
        audio.volume = 0.75;
        audio.pause();
        audio.currentTime = 0;
        audio.play().catch((err) => console.error('Audio playback failed:', err));
    }
}*/

//const audioManager = new SoundManager();

function generateNotification(type,title,message) {
    const alerts = [
        {
            type: 'error',
            title: 'Lorem Ipsum Error',
            message: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
        },
        {
            type: 'success',
            title: 'Lorem Ipsum Success',
            message: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
        },
        {
            type: 'info',
            title: 'Lorem Ipsum Info',
            message: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
        }
    ];

    //const notification = alerts[Math.floor(Math.random() * alerts.length)];
    const notification = {
        type: type,
        title: title,
        message: message
    }
    
    notification.id = Math.random().toString(36).substr(2, 9);
    notificationItems.push(notification);
    //audioManager.playOpenSound();
    
    displayNotification(notification);
}

function generateNotifications() {
    for (let index = 0; index < 4; index++) {
        setTimeout(() => {
            //generateNotification();
        }, 500 * index);
    }
}

function closeNotifications() {
    const notificationsContainer = document.getElementById('notifications');
    notificationsContainer.innerHTML = '';
    notificationItems.length = 0;
    audioManager.playCloseSound();
}

function displayNotification(notification) {
    const notificationsContainer = document.getElementById('notifications');
    const notificationDiv = document.createElement('div');
    notificationDiv.className = `notification ${notification.type} fade show`;
    notificationDiv.id = notification.id;

    notificationDiv.innerHTML = `
        <div>
          <div class="notification-title">${notification.title}</div>
          <div>${notification.message}</div>
        </div>
        <button onclick="removeNotification('${notification.id}')">X</button>
      `;

    notificationsContainer.appendChild(notificationDiv);

    setTimeout(function() {
        notificationDiv.classList.remove('show');
        setTimeout(function(){removeNotification(notification.id)}, 500);
    }, 5000);
}

function removeNotification(id) {
    const index = notificationItems.findIndex(item => item.id === id);
    console.log(notificationItems)
    if (index !== -1) {
        notificationItems.splice(index, 1);
        const notificationDiv = document.getElementById(id);
        if (notificationDiv) {
            audioManager.playCloseSound();
            notificationDiv.remove();
        }
    }
}

window.generateNotification = //generateNotification;
window.removeNotification = removeNotification;
