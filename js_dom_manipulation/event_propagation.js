// Event popagation

// the third argument of the addEventListener method is the capture phase, its optional default=false
// capture phase -> from the top of the DOM to the target element
window.addEventListener('click', function () {
    console.log('Window');
}, true);
// to see this comment for while the stop propagation method bellow

document.addEventListener('click', function () {
    console.log('Document');
}, true);

document.querySelector('.div1').addEventListener('click', function () {
    console.log('Box 1');
}, true);

document.querySelector('.div2').addEventListener('click', function () {
    console.log('Box 2');
}, true);

document.querySelector('button').addEventListener('click', function (e) {
    console.log(e); // e is the event object, has many properties see console...
}, true); // true is the capture phase, false is the bubbling phase

// watch the console, the event object is logged in the capture phase, then in the bubbling phase
// bubbling phase -> default, from the target element to the top of the DOM
window.addEventListener('click', function () {
    console.log('Window');
}, false);

document.addEventListener('click', function () {
    console.log('Document');
}, false);

document.querySelector('.div1').addEventListener('click', function () {
    console.log('Box 1');
}, false);

document.querySelector('.div2').addEventListener('click', function () {
    console.log('Box 2');
}, false);



// we can change the default behaviour of the event object
document.querySelector('button').addEventListener('click', function (e) {
    console.log(e.target.innerText = 'Clicked!'); 
}, false); 

// we can stop propagation of the event object
// lets say we dont want to propagate further then div2...
document.querySelector('.div2').addEventListener('click', function (e) {
    e.stopPropagation();
    console.log('Box 2');
});

document.querySelector('.div1').addEventListener('click', function () {
    console.log('div 1'); 
}, {once: true}); // once: true, means that the event will be fired only once, then removed


// because the a tag is trying to redirect to another page,
// it refreshes the page and we lose progress - change to click,
// we can prevent that with the preventDefault method
document.querySelector('.link').addEventListener('click', function (e) {
    e.preventDefault(); // if u comment this u ll see what i mean
    console.log(e.target.innerText = 'Clicked!'); 
}, false);




