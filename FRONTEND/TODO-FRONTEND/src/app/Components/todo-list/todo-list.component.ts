import { Component, DebugElement, OnInit } from '@angular/core';
import { TaskPriority, TodoItem } from 'src/app/interfaces';
import { LocalApiService } from 'src/app/local-api.service';

interface TodoItemDictionary{
  [todoId: number] : TodoItem;
}

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss']
})
export class TodoListComponent implements OnInit {
  todoItemsDictionary: TodoItemDictionary = {};

  constructor(
    private apiService: LocalApiService
    ) { 

  }

  ngOnInit(): void {
    this.apiService.getTodoItems().subscribe(todoItems => {
      console.log(todoItems)
      this.generateDictonaryFromTodoItemList(todoItems);
    })
  }

  generateDictonaryFromTodoItemList(todoItemList: TodoItem[]) {
    todoItemList.forEach(task => {
      this.todoItemsDictionary[task.todoID] = task;
    })
  }

  markTaskAsDone(id: number): void {
    this.todoItemsDictionary[id].done = true;
  }

  markTaskAsUndone(id: number): void {
    this.todoItemsDictionary[id].done = false;
  }
}
