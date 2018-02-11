import { NgModule } from '@angular/core';

import { ThemeModule } from '../@theme/theme.module';
import { PackageComponent } from './package/package.component';
import { PackageListComponent } from './package-list/package-list.component';
import { PackageWithStateComponent } from './package-with-state/package-with-state.component';

@NgModule({
  imports: [
    ThemeModule,
  ],
  declarations: [
    PackageComponent,
    PackageListComponent,
    PackageWithStateComponent
  ],
})
export class ComponentsModule { }
