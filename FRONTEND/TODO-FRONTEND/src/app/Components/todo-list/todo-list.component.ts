import { Component, DebugElement, OnInit } from '@angular/core';
import { TaskPriority, TodoItem } from 'src/app/interfaces';
import { LocalApiService } from 'src/app/local-api.service';
import { MatDialog } from '@angular/material/dialog';
import { CreateTodoDialogComponent } from 'src/app/create-todo-dialog/create-todo-dialog.component';

interface TodoItemDictionary {
  [todoId: number]: TodoItem;
}

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss']
})
export class TodoListComponent implements OnInit {
  todoItemsDictionary: TodoItemDictionary = {};

  constructor(
    private apiService: LocalApiService,
    private dialog: MatDialog
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
    this.apiService.updateTodoItem(this.todoItemsDictionary[id]).subscribe()
  }

  markTaskAsUndone(id: number): void {
    this.todoItemsDictionary[id].done = false;
    this.apiService.updateTodoItem(this.todoItemsDictionary[id]).subscribe()
  }

  openNewTodoItemDialog(): void {
    const dialogRef = this.dialog.open(CreateTodoDialogComponent);
    dialogRef.afterClosed().subscribe(() => {
      this.apiService.getTodoItems().subscribe(todoItems => {
        console.log(todoItems);
        this.generateDictonaryFromTodoItemList(todoItems);
      })
    })
  }
}
