import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from "@angular/common/http";
import {Observable} from "rxjs";

export class AuthInterceptorService implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    let sessionID = localStorage.getItem('sessionID');

    if (sessionID !=  null) {
      const modifiedRequest = req.clone({headers: req.headers.set('sessionID', sessionID)});
      return next.handle(modifiedRequest);
    }

    return next.handle(req);
  }
}
