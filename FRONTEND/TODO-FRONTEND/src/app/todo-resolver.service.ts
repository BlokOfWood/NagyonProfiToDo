import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';
import { TodoItem } from './interfaces';
import { TodoItemsService } from './todo-items-service';

@Injectable({
  providedIn: 'root'
})
export class TodoResolverService implements Resolve<TodoItem[]> {
  constructor(
    private todoItemsService: TodoItemsService
  ) { }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<TodoItem[]> {
    return this.todoItemsService.fetchTodoItems();
  }
}
