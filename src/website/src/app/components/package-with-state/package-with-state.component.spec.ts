import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PackageWithStateComponent } from './package-with-state.component';

describe('PackageWithStateComponent', () => {
  let component: PackageWithStateComponent;
  let fixture: ComponentFixture<PackageWithStateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PackageWithStateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PackageWithStateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
