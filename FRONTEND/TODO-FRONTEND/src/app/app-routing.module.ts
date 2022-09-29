import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './Components/login/login.component';
import { RegistrationComponent } from './Components/registration/registration.component';
import { TodoListComponent } from './Components/todo-list/todo-list.component';
import {AuthGuardService} from "./auth-guard.service";

const routes: Routes = [
  { path: 'todolist', component: TodoListComponent, canActivate: [AuthGuardService]},
  { path: 'register', component: RegistrationComponent},
  { path: 'login', component: LoginComponent},
  { path: '', redirectTo: '/todolist', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
