// DOM manipulation
// Event delegation

// it allows users to append a SINGLE event listener to a parent element, 
// and that event listener will fire for all descendants matching a selector, 
// whether those descendants exist now or are added in the future 
// - dam tatkovi neco co muze pouzit na svoje deti ktere ma uz ted nebo bude mit v budoucnu
// - velmi ucinne, protoze se nemusi pridavat event listener na kazdeho potomka zvlast

const nextSports = ["football", "basketball", "tennis", "volleyball", "hockey", "baseball", "golf", "rugby", "cricket", "badminton"];

const sports = document.querySelector(".container");

document.querySelector('.container').addEventListener('click', function (e) {
    console.log(e.target); // e.target is the element that was clicked
    if (e.target.className === 'container' || e.target.className === 'heading') {
        if (nextSports.length > 0) {
            addSport(nextSports.pop());
        }
    } else {
        if (e.target.style.backgroundColor === 'white') {    
            e.target.style.color = 'white';
            e.target.style.backgroundColor = 'black';
            console.log(e.target.getAttribute('id') + ' Unclicked');
        } else {
            e.target.style.backgroundColor = 'white';
            e.target.style.color = 'black';
            console.log(e.target.getAttribute('id') + ' Clicked');
        }
    }
});


const addSport = (sport) => {
    const div = document.createElement("div");
    div.classList.add("box");
    div.setAttribute("id", sport);
    div.innerText = sport;
    sports.appendChild(div);
    
    // need to get the height of the container and add 15px to it
    const currHeight = sports.clientHeight;
    sports.style.height = (currHeight + 100).toString() + "px";
    console.log(sports.style.height);
}

