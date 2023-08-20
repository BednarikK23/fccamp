// DOM Manipulation
// Traversal DOM
// Parent Node Traversal

let ul = document.querySelector('.ul');

console.log(ul.parentNode);
console.log(ul.parentElement); 
// same as parentNode it only differs 
// when the parent is not an element node but a text node

// can be chained
console.log(ul.parentNode.parentNode);
console.log(ul.parentElement.parentElement);

const html = document.documentElement;
console.log(html);
console.log(html.parentNode); // #document
console.log(html.parentElement); // null
// -> thats why parentNode is more reliable, and used more often


// Child Node Traversal

console.log(ul.childNodes); 
// indentations and line breaks are also considered as text nodes!!
// NodeList(15) [text, li.list-items, text, li.list-items, text, li.list-items, text, li.list-items, text, li.list-items, text]
ul.removeChild(ul.childNodes[12]); // removes the first text node -> just removing the indentation...
console.log(ul.childNodes); // NoedeList(14)...
// but if i were to delete this... the Avengers disappear... (snap, i've seen this one this one is a classic)
ul.removeChild(ul.childNodes[12]);

ul.childNodes[1].style.backgroundColor = 'cyan';
 
console.log(ul.children); 
console.log(ul.firstElementChild);
console.log(ul.lastElementChild);

// Sibling Node Traversal
// watchout -> comment <!-- ---> is also a node and therefore sibling, child...

const div = document.querySelector('div');

console.log(ul.previousSibling);
console.log(ul.nextSibling);
// these gives us the text nodes,
console.log(div.childNodes); // here u can see why...

// THEREFORE its better to use:
console.log(ul.previousElementSibling);
console.log(ul.nextElementSibling); // null, when no such sibling
// it ignore #texts and also comments!! when cannot 
