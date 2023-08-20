// DOM BASIC MANIPULATION
// how to select element in dom? most basic functions:

// document.getElementById('id')
const tit = document.getElementById('main-heading');

// document.getElementsByClassName('class')
const items = document.getElementsByClassName('list-items');
console.log(items);

// document.getElementsByTagName('tag')
const l = document.getElementsByTagName('li');

// document.querySelector('css selector')
const container = document.querySelector('.container'); // could be 'div' or #id...
console.log(container);

// document.querySelectorAll('css selector')
const container2 = document.querySelectorAll('div');
console.log(container2);

// how to create element in dom?
// document.createElement('tag')

// how to add element to dom?
// parent.appendChild(child)

// how to remove element from dom?
// parent.removeChild(child)

// how to add event listener to element?
// element.addEventListener('event', function() {})

// okay, now we have the summary of basic dom manipulation, lets go..


// STYLING
const title = document.getElementById('main-heading');
title.style.color = 'red' ;

const listItems = document.querySelectorAll('.list-items');


for (let i = 0; i < listItems.length; i++) {
    listItems[i].style.fontSize = '15px';
}

// CREATING ELEMENT
const ul = document.querySelector('.ul');
ul.style.color = 'blue';
const li = document.createElement('li');
li.textContent = 'The Spiderman';
li.addClassName = 'list-items';
ul.appendChild(li);


// MODIFYING ELEMENT
const firstListstItem = document.querySelector('.list-items');
console.log(firstListstItem.innerText);  // without tags in one line
console.log(firstListstItem.textContent);  // just text on multiple lines, w indents
console.log(firstListstItem.innerHTML);  // with <span> tag, with indents, multiple lines

li.innerText = 'X-men'; // now we changed just text in variable


// MODIFYING ATTRIBUTES
// - edits/sets attribute of element ('name of attribute', 'value of attribute')
li.setAttribute('id', 'main-heading'); // now we changed class in variable
li.removeAttribute('id'); // now we removed id in variable, so it looks normal again
console.log(title.getAttribute('id')); // returns value of attribute

li.classList.add('list-items'); // adds class to element
// li.classList.remove('list-items'); // removes class from element

// REMOVE ELEMENTS
//li.remove(); // removes element from dom
