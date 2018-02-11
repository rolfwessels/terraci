import { Component, ChangeDetectionStrategy } from '@angular/core';
import { ApiService, CurrentState } from '../../@core/data/api.service';
import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/do';

@Component({
  selector: 'ngx-dashboard',
  templateUrl: './dashboard.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DashboardComponent {
  private loading: boolean = true;
  public currentState$: Observable<CurrentState> ;

  constructor(protected stateService: ApiService)
  {

  }

  ngOnInit() {
    this.loadData();
  }

  loadData() {
    this.currentState$ = this.stateService.getCurrentState()
    .do(_ => this.loading = false);
  }
}
