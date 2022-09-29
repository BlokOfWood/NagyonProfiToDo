import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginInfo } from 'src/app/interfaces';
import {AuthService} from "../../auth-service";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  formGroup: FormGroup = new FormGroup({
    username: new FormControl(''),
    password: new FormControl('')
  })

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  loginAttempt(): void {
    const loginInfo: LoginInfo = this.formGroup.value;

    this.authService.login(loginInfo, this.formGroup.value.rememberInfo).subscribe(() => {
      console.log('login successful');
      this.router.navigate(['todolist']);
    })
  }
}
