import {Injectable} from "@angular/core";
import {LocalApiService} from "./local-api.service";
import {LoginInfo, RegistrationInfo, SessionIdResponse} from "./interfaces";
import {Observable, tap} from "rxjs";
import {Router} from "@angular/router";

@Injectable({
  providedIn: "root"
})
export class AuthService {
  public loggedIn: boolean = false;
  public sessionID: string | null = null;

  constructor(
    private localApiService: LocalApiService,
    private router: Router
  ) {
    let localSessionID = sessionStorage.getItem('sessionID');
    if(localSessionID == null)
      localSessionID = localStorage.getItem('sessionID')

    if(localSessionID != null) {
      this.loggedIn = true;
      this.sessionID = localSessionID;

      sessionStorage.setItem('sessionID', localSessionID);
      /*localApiService.checkSessionId(localSessionID).subscribe(
        (response) => {
          if (response.isCorrectSessionID) {
            this.authData.loggedIn = true;
            this.authData.sessionID = localSessionID;
          }
        }
      )*/
    }
  }

  saveSessionId(sessionID: string) {
    localStorage.setItem('sessionID', sessionID)
  }

  login(loginInfo: LoginInfo, rememberInfo: boolean): Observable<SessionIdResponse> {
    return this.localApiService.attemptLogin(loginInfo).pipe(
      tap((response) => {
        this.loggedIn = true;
        this.sessionID = response.sessionID;
        sessionStorage.setItem('sessionID', response.sessionID);

        if(rememberInfo)
          this.saveSessionId(response.sessionID);
      }));
  }

  register(registrationInfo: RegistrationInfo): Observable<SessionIdResponse> {
    return this.localApiService.attemptRegister(registrationInfo).pipe(
      tap((response) => {
        this.loggedIn = true;
        this.sessionID = response.sessionID;
        sessionStorage.setItem('sessionID', response.sessionID);
      }));
  }

  logout(): void {
    this.loggedIn = false;
    this.sessionID = null;

    sessionStorage.removeItem('sessionID');
    localStorage.removeItem('sessionID');

    this.router.navigate(['login']);
  }
}
