import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import {TodoItemsService} from "../todo-items-service";

@Component({
  selector: 'app-create-todo-dialog',
  templateUrl: './create-todo-dialog.component.html',
  styleUrls: ['./create-todo-dialog.component.scss']
})
export class CreateTodoDialogComponent implements OnInit {
  formGroup: FormGroup = new FormGroup({
    name: new FormControl('')
  })

  constructor(
    private dialogRef: MatDialogRef<CreateTodoDialogComponent>,
    private todoItemsService: TodoItemsService
  ) { }

  ngOnInit(): void {
  }

  closeDialog(): void {
    this.dialogRef.close();
  }

  createTodo(): void {
    this.todoItemsService.createTodoItem(this.formGroup.value.name).subscribe(() => {
      this.closeDialog();
    })
  }
}
