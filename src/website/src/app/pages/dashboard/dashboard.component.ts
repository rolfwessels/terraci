import { Component, ChangeDetectionStrategy } from '@angular/core';
import { ApiService, CurrentState } from '../../@core/data/api.service';
import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/do';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'ngx-dashboard',
  templateUrl: './dashboard.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DashboardComponent {
  private loading: boolean = true;
  public currentState$: Observable<CurrentState> ;

  constructor(
    protected stateService: ApiService,
    private activatedRoute: ActivatedRoute
  )
  {
      console.log(this.activatedRoute.queryParams)
      this.activatedRoute.queryParams.subscribe(params => {
          console.log("o",params)
          let date = params['startdate'];
          console.log(date); // Print the parameter to the console.
      });
  }


  ngOnInit() {
    this.loadData();
  }

  loadData() {
    this.currentState$ = this.stateService.getCurrentState()
    .do(_ => this.loading = false);
  }
}
