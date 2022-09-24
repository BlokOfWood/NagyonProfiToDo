import {Component, OnInit} from '@angular/core';
import {TaskPriority, } from 'src/app/interfaces';
import {MatDialog} from '@angular/material/dialog';
import {CreateTodoDialogComponent} from 'src/app/create-todo-dialog/create-todo-dialog.component';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {TodoItemsService} from "../../todo-items-service";
import {Subscription} from "rxjs";

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss']
})
export class TodoListComponent implements OnInit {
  editedID: number = 0;

  currentFormSubscription: Subscription | null = null;

  formGroup = new FormGroup({
    name: new FormControl('', Validators.required),
    description: new FormControl('', Validators.required),
    priority: new FormControl(TaskPriority.NORMAL, Validators.required),
    done: new FormControl(false),
    deadline: new FormControl(new Date(), Validators.required)
  }, {
      updateOn: "blur"
  });
  taskPriorities = Object.values(TaskPriority);

  constructor(
    public todoItemService: TodoItemsService,
    private dialog: MatDialog
  ) {}

  ngOnInit(): void {
  }

  markTaskAsDone(id: number): void {
    this.todoItemService.changeTodoItemCompletion(id, true);
  }

  markTaskAsUndone(id: number): void {
    this.todoItemService.changeTodoItemCompletion(id, false);
  }

  openTodoItemPanel(id: number): void {
    if(this.currentFormSubscription != null)
      this.currentFormSubscription.unsubscribe();

    this.editedID = id;

    this.formGroup.setValue(this.todoItemService.getTodoEditor(id)!);

    this.currentFormSubscription = this.formGroup.valueChanges.subscribe(_ => {
      this.updateTodoItem(id);
    })
  }

  updateTodoItem(id: number): void {
    const formValue = this.formGroup.value;
    this.todoItemService.updateTodoItem(id, formValue.name!, formValue.description!, formValue.priority!, formValue.deadline!)
  }

  openNewTodoItemDialog(): void {
    this.dialog.open(CreateTodoDialogComponent);
  }
}
