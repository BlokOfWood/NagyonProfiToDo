import { Injectable } from '@angular/core';
import { from, Observable } from 'rxjs';
import { LoginInfo } from './interfaces';

@Injectable({
  providedIn: 'root'
})
export class LocalApiService {
  sessionID: string = "";

  constructor() { }

  attemptLogin(loginInfo: LoginInfo): Observable<void> {
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var requestOptions = {
      body: JSON.stringify(loginInfo),
      method: 'POST',
      headers: myHeaders,
    };

    return from(
      fetch("http://81.182.202.18:4000/login", requestOptions)
        .then(response => response.json().then(x => {
          if (response.status === 200) {
            this.sessionID = x.sessionID;
          }
        }))
        .then(result => console.log(result))
        .catch(error => console.log('error', error)));
  }


}
