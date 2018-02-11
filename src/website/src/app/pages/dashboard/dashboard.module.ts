import { NgModule } from '@angular/core';

import { ThemeModule } from '../../@theme/theme.module';
import { DashboardComponent } from './dashboard.component';
import { PackageComponent } from '../../components/package/package.component';
import { PackageListComponent } from '../../components/package-list/package-list.component';

@NgModule({
  imports: [
    ThemeModule,
  ],
  declarations: [
    DashboardComponent,
    PackageComponent,
    PackageListComponent
  ],
})
export class DashboardModule { }
