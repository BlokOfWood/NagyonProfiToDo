import { Injectable } from '@angular/core';
import { from, Observable } from 'rxjs';
import { LoginInfo } from './interfaces';
import { APIFunctions } from './request-helper';

@Injectable({
  providedIn: 'root'
})
export class LocalApiService {
  sessionID: string = "";
  apiFunctions: APIFunctions = new APIFunctions("http://81.182.202.18:4000/");

  constructor() { }

  attemptLogin(loginInfo: LoginInfo): Observable<string> {
    return this.apiFunctions.post('login', loginInfo)
  }

  getTodoItems(): void {
    APIFunctions 
  }
}
