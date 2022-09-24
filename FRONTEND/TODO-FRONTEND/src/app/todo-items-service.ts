import {Injectable} from "@angular/core";
import {TaskPriority, TodoEditor, TodoItem} from "./interfaces";
import {LocalApiService} from "./local-api.service";
import {Observable} from "rxjs";

@Injectable({
  providedIn: "root"
})
export class TodoItemsService {
  todoItemsDictionary: Map<number, TodoItem> = new Map();

  constructor(private localApiService: LocalApiService) {
    localApiService.getTodoItems().subscribe(
      todoItems => {
        this.initalizeTodoDictionary(todoItems);
      }
    )
  }

  initalizeTodoDictionary(todoItems: TodoItem[]) {
    todoItems.forEach(task => {
      this.todoItemsDictionary.set(task.todoID, task);
    })
  }

  getTodoItem(id: number): TodoItem | undefined {
    if(!this.todoItemsDictionary.has(id)) {
      console.error(`Failure to get todo item. No todo item with id ${id} present.`);
      return undefined;
    }
    return this.todoItemsDictionary.get(id);
  }

  getTodoEditor(id: number): TodoEditor | undefined {
    if(!this.todoItemsDictionary.has(id)) {
      console.error(`Failure to get todo item. No todo item with id ${id} present.`);
      return undefined;
    }

    return this.todoItemToTodoEditor(this.todoItemsDictionary.get(id)!);
  }

  createTodoItem(todoTitle: string): Observable<any> {
    let todoEditor: TodoEditor = {
      name: todoTitle,
      description: "",
      priority: TaskPriority.NORMAL,
      done: false,
      deadline: new Date(),
    }

    return this.localApiService.createTodoItem(todoEditor);
  }

  changeTodoItemCompletion(id: number, completion: boolean) {
    if(!this.todoItemsDictionary.has(id)) {
      console.error(`Failure to change task completion. No todo item with id ${id} present.`);
      return;
    }

    let editedTodoItem = this.todoItemsDictionary.get(id)!;
    editedTodoItem.done = completion;

    this.localApiService.updateTodoItem(id, editedTodoItem);
  }

  updateTodoItem(id: number, name: string, description: string, priority: TaskPriority, deadline: Date){
    if(!this.todoItemsDictionary.has(id)) {
      console.error(`Failure to update todo item. No todo item with id ${id} present.`);
      return;
    }

    let editedTodoItem = this.todoItemsDictionary.get(id)!
    editedTodoItem.name = name;
    editedTodoItem.description = description;
    editedTodoItem.priority = priority;
    editedTodoItem.deadline = deadline;

    this.todoItemsDictionary.set(id, editedTodoItem);
    this.localApiService.updateTodoItem(id, editedTodoItem).subscribe();
  }

  todoItemToTodoEditor(todoItem: TodoItem): TodoEditor {
    return {
      name: todoItem.name,
      description: todoItem.description,
      priority: todoItem.priority,
      done: todoItem.done,
      deadline: todoItem.deadline
    }
  }
}
