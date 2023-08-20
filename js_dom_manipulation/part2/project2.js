// Vars

const accordion = document.getElementsByClassName('content-container');

for (let i = 0; i < accordion.length; i++) {
    accordion[i].addEventListener('click', function() {
        // toggle - add class if it's not there, remove if it is.
        this.classList.toggle('active'); // we can nicely use this.
    });
}



