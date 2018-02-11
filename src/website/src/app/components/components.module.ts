import { NgModule } from '@angular/core';

import { ThemeModule } from '../@theme/theme.module';
import { PackageComponent } from './package/package.component';

@NgModule({
  imports: [
    ThemeModule,
  ],
  declarations: [
    PackageComponent
  ],
})
export class ComponentsModule { }
