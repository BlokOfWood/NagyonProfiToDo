import { Injectable } from '@angular/core';
import {Observable, tap} from 'rxjs';
import {LoginInfo, RegistrationInfo, SessionIdResponse, TodoEditor, TodoItem} from './interfaces';
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class LocalApiService {
  apiAddress = "http://localhost:4000/";

  constructor(private httpClient: HttpClient) { }

  attemptLogin(loginInfo: LoginInfo): Observable<SessionIdResponse> {
    return this.httpClient.post<SessionIdResponse>(this.apiAddress + "login", loginInfo)
      .pipe(
        tap(x =>  {
          localStorage.setItem('sessionID', x.sessionID);
        }
      ));
  }

  attemptRegister(registrationInfo: RegistrationInfo): Observable<SessionIdResponse> {
    return this.httpClient.post<SessionIdResponse>(this.apiAddress + "register", registrationInfo)
      .pipe(
        tap(x =>  {
          localStorage.setItem('sessionID', x.sessionID);
        })
      );
  }

  getTodoItems(): Observable<TodoItem[]> {
    return this.httpClient.get<TodoItem[]>(this.apiAddress + "todos");
  }

  createTodoItem(newTodoItem: TodoEditor): Observable<Object> {
    return this.httpClient.post(this.apiAddress + "todos", newTodoItem);
  }

  updateTodoItem(todoId: number, todoEditor: TodoEditor): Observable<Object> {
    return this.httpClient.patch(this.apiAddress + "todos/" + todoId, todoEditor);
  }
}
