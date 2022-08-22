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
  listOfTasks: TodoItem[] = [
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

  listOfTasksDictionary: TodoItemDictionary = {};

  constructor() { }

  ngOnInit(): void {
    this.generateDictonaryFromTodoItemList(this.listOfTasks)
  }

  generateDictonaryFromTodoItemList(todoItemList: TodoItem[]) {
    todoItemList.forEach(task => {
      this.listOfTasksDictionary[task.taskID] = task;
    })
  }

  markTaskAsDone(id: number): void {
    this.listOfTasksDictionary[id].taskDone = true;
  }

  markTaskAsUndone(id: number): void {
    this.listOfTasksDictionary[id].taskDone = false;
  }
}
