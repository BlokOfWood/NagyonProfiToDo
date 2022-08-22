import { Component, DebugElement, OnInit } from '@angular/core';
import { TaskPriority, TodoItem } from 'src/app/interfaces';

interface TodoItemDictionary{
  [todoId: number] : TodoItem;
}

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss']
})
export class TodoListComponent implements OnInit {
  listOfTodoItems: TodoItem[] = [
    {
      taskCreatedAt:new Date(), 
      taskDeadline: new Date(), 
      taskDone: true, 
      taskID: 1, 
      taskName: "asdasda", 
      taskPriority: TaskPriority.CRITICAL, 
      taskText:"as"
    },
    {
      taskCreatedAt:new Date(), 
      taskDeadline: new Date(), 
      taskDone: false, 
      taskID: 2, 
      taskName: "asdasdaasdasd", 
      taskPriority: TaskPriority.EVENTUALLY, 
      taskText:"as"
    }
  ];

  todoItemsDictionary: TodoItemDictionary = {};

  constructor() { }

  ngOnInit(): void {
    this.generateDictonaryFromTodoItemList(this.listOfTodoItems)
  }

  generateDictonaryFromTodoItemList(todoItemList: TodoItem[]) {
    todoItemList.forEach(task => {
      this.todoItemsDictionary[task.taskID] = task;
    })
  }

  markTaskAsDone(id: number): void {
    this.todoItemsDictionary[id].taskDone = true;
  }

  markTaskAsUndone(id: number): void {
    this.todoItemsDictionary[id].taskDone = false;
  }
}
