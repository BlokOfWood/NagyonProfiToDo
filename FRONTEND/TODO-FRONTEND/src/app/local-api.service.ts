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

  updateTodoItem(): Observable<void> {
    return new Observable<void>(
      subscriber => {
        this.apiFunctions.patch('todos', new Headers()).subscribe(() => {
          subscriber.next()
          subscriber.complete()
        })
      }
    );
  }
}
