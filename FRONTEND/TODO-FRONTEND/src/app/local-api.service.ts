import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { LoginInfo, RegistrationInfo, TaskPriority, TodoEditor, TodoItem } from './interfaces';
import { APIFunctions } from './request-helper';

@Injectable({
  providedIn: 'root'
})
export class LocalApiService {
  apiFunctions: APIFunctions = new APIFunctions("http://81.182.202.18:4000/");

  constructor() { }

  attemptLogin(loginInfo: LoginInfo): Observable<void> {
    return new Observable<void>(
      subscriber => {
        this.apiFunctions.post('login', loginInfo).subscribe(sessionIDResponse => {
          localStorage.setItem('sessionID', JSON.parse(sessionIDResponse).sessionID)
          subscriber.next()
          subscriber.complete()
        })
      }
    );
  }

  attemptRegister(registrationInfo: RegistrationInfo): Observable<void> {
    return new Observable<void>(
      subscriber => {
        this.apiFunctions.post('register', registrationInfo).subscribe(sessionIDResponse => {
          subscriber.next()
          subscriber.complete()
        })
      }
    );
  }

  getTodoItems(): Observable<TodoItem[]> {
    return this.apiFunctions.get('todos', new Headers())
      .pipe(map(todoItemList => {
        return JSON.parse(todoItemList)
      }))
  }

  createTodoItem(todoTitle: string): Observable<void> {
    var todoEditor: TodoEditor = {
      name: todoTitle,
      priority: TaskPriority.NORMAL,
      done: false,
      description: "",
      deadline: new Date(),
    }

    return new Observable<void>(
      subscriber => {
        this.apiFunctions.post('todos', todoEditor, new Headers()).subscribe(() => {
          subscriber.next()
          subscriber.complete()
        })
      }
    );
  }

  updateTodoItem(todoItem: TodoItem): Observable<void> {
    return new Observable<void>(
      subscriber => {
        this.apiFunctions.patch('todos/' + todoItem.todoID, todoItem, new Headers()).subscribe(() => {
          subscriber.next()
          subscriber.complete()
        })
      }
    );
  }
}
