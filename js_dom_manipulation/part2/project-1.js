// Variables

let btn = document.querySelector('#new-quote');
let quote = document.querySelector('.quote');
let person = document.querySelector('.person');


const quotes = [
  {
    "quote": "The only way to do great work is to love what you do.",
    "person": "Steve Jobs"
  },
  {
    "quote": "In three words I can sum up everything I've learned about life: it goes on.",
    "person": "Robert Frost"
  },
  {
    "quote": "The future belongs to those who believe in the beauty of their dreams.",
    "person": "Eleanor Roosevelt"
  },
  {
    "quote": "Strive not to be a success, but rather to be of value.",
    "person": "Albert Einstein"
  },
  {
    "quote": "The only limit to our realization of tomorrow will be our doubts of today.",
    "person": "Franklin D. Roosevelt"
  },
  {
    "quote": "Success is not final, failure is not fatal: It is the courage to continue that counts.",
    "person": "Winston Churchill"
  },
  {
    "quote": "It does not matter how slowly you go as long as you do not stop.",
    "person": "Confucius"
  },
  {
    "quote": "The best revenge is massive success.",
    "person": "Frank Sinatra"
  },
  {
    "quote": "The only person you are destined to become is the person you decide to be.",
    "person": "Ralph Waldo Emerson"
  },
  {
    "quote": "The way to get started is to quit talking and begin doing.",
    "person": "Walt Disney"
  }
]


btn.addEventListener('click', function() { 
    let random = Math.floor(Math.random() * quotes.length);
    quote.innerHTML = quotes[random].quote;
    person.innerHTML = quotes[random].person;
});



