// EVENTS

// more evenets: https://developer.mozilla.org/en-US/docs/Web/Events
// https://www.w3schools.com/jsref/dom_obj_event.asp

// Event Listener

// first way how to respont to event can be done in html file
// <button onclick="alert('hello')">Click me</button> -> like so in our events.html file
// but you can have only one event listener per element like this
// js allow us apply several event listeners to one element like so:

// element.addEventListener('event', function() {})

const buttonTwo = document.querySelector('.btn-2');

function alertBtn () {
    alert('General Kenobi');
}

buttonTwo.addEventListener('click', alertBtn);

// Mouse over here
const boxThree = document.querySelector('.box-3');
boxThree.style.backgroundColor = 'cyan';

function changeBg () {
    boxThree.style.backgroundColor = 'red';
}

boxThree.addEventListener('mouseover', changeBg);

// SND PART

// this is useful for  example for hamburger menu and other stuff...
const revealBtn = document.querySelector('.reveal-btn');
const hiddenContent = document.querySelector('.hidden-content');

function revealContent () {
    if (hiddenContent.classList.contains('reveal-btn')) {
        hiddenContent.classList.remove('reveal-btn');
    } else {
        hiddenContent.classList.add('reveal-btn');
    }
}

revealBtn.addEventListener('click', revealContent);


