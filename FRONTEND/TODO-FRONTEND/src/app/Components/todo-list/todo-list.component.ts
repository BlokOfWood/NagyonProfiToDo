import { Component, OnInit } from '@angular/core';
import { TaskPriority, TodoItem } from 'src/app/interfaces';

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

  constructor() { }

  ngOnInit(): void {

  }

}
