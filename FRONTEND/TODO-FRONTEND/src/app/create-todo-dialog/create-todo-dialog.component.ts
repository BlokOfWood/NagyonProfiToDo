import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
import { LocalApiService } from '../local-api.service';

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
    private localApiService: LocalApiService
  ) { }

  ngOnInit(): void {
  }

  closeDialog(): void {
    this.dialogRef.close();
  }

  createTodo(): void {
    this.localApiService.createTodoItem(this.formGroup.value.name).subscribe(() => {
      this.closeDialog();
      this.localApiService.getTodoItems();
    })
  }
}