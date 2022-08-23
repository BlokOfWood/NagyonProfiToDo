import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { RegistrationInfo } from 'src/app/interfaces';
import { LocalApiService } from 'src/app/local-api.service';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss']
})
export class RegistrationComponent implements OnInit {
  triedToRegister: boolean = false;

  formGroup: FormGroup = new FormGroup({
    username: new FormControl(''),
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl(''),
    repeatedPassword: new FormControl('')
  })
  get email() {
    return this.formGroup.get('email'); 
  }
  currentErrorMessage : string = "";

  constructor(
    private apiService : LocalApiService,
    private router : Router
  ) { }

  ngOnInit(): void {
  }

  registerAttempt(): void {
    if(!this.formGroup.valid) {
      console.log("Form is not valid");
      this.triedToRegister = true;
      return;
    }
      
    const registrationInfo: RegistrationInfo = this.formGroup.value;

    this.apiService.attemptRegister(registrationInfo).subscribe(() => {
      console.log('registration successful');
      this.router.navigate(['login']);
    })
  }
}
