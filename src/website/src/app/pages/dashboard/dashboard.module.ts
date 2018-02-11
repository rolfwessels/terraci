import { NgModule } from '@angular/core';

import { ThemeModule } from '../../@theme/theme.module';
import { DashboardComponent } from './dashboard.component';
import { PackageComponent } from '../../components/package/package.component';

@NgModule({
  imports: [
    ThemeModule,
  ],
  declarations: [
    DashboardComponent,
    PackageComponent
  ],
})
export class DashboardModule { }
