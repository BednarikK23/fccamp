
const addTask = document.getElementById('add-task');
const taskContainer = document.getElementById('task-container');
const inputTask = document.getElementById('input-task');

addTask.addEventListener('click', () => {   
    const task = document.createElement('div');
    task.classList.add('task');
    
    let li = document.createElement('li');
    li.innerText = inputTask.value;
    task.appendChild(li);

    let checkButton = document.createElement('button');
    checkButton.innerHTML = '<i class="fa-check fa-solid"></i>';
    checkButton.classList.add('checkTask');
    task.appendChild(checkButton);
    
    let deleteButton = document.createElement('button');
    deleteButton.innerHTML = '<i class="fa-solid fa-trash"></i>';
    deleteButton.classList.add('deleteTask');
    task.appendChild(deleteButton);

    if (inputTask.value === '') {
        alert('Please enter a task');
    } else {
        taskContainer.appendChild(task);
    }

    inputTask.value = '';


    checkButton.addEventListener('click', () => {
        checkButton.parentElement.style.textDecoration = 'line-through';
    })

    
    deleteButton.addEventListener('click', (e) => {
        console.log(e.target.classList, e.target);
        let target = e.target;
        if (target.classList.contains('deleteTask')) {
            target.parentElement.remove();
        }
        if (target.classList.contains('fa-solid')) {
            target.parentElement.parentElement.remove();
        }
    });

} );





