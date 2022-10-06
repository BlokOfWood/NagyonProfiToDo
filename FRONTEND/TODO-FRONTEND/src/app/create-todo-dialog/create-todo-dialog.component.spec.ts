import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateTodoDialogComponent } from './create-todo-dialog.component';
import {HarnessLoader} from "@angular/cdk/testing";
import {TestbedHarnessEnvironment} from "@angular/cdk/testing/testbed";
import {MatDialogHarness} from "@angular/material/dialog/testing";
import {MatDialogRef} from "@angular/material/dialog";


describe('CreateTodoDialogComponent', () => {
  let fixture: ComponentFixture<CreateTodoDialogComponent>;
  let loader: HarnessLoader;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MatDialogRef],
      declarations: [ CreateTodoDialogComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateTodoDialogComponent);
    fixture.detectChanges();
    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('a', () => expect(true).toBeTrue())
});
