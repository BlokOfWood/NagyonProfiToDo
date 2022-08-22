import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { LoginInfo } from 'src/app/interfaces';
import { LocalApiService } from '../../local-api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  formGroup: FormGroup = new FormGroup({
    username: new FormControl(''),
    password: new FormControl('')
  })


  constructor(
    public apiService: LocalApiService
  ) { 

  }

  ngOnInit(): void {
  }

  loginAttempt(): void {
    const loginInfo: LoginInfo = this.formGroup.value;

    this.apiService.attemptLogin(loginInfo)
  }

}
