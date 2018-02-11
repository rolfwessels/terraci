import { Component, OnInit,Input,OnChanges } from '@angular/core';
import { Package } from '../../@core/data/api.service';

@Component({
  selector: 'app-package',
  templateUrl: './package.component.html',
  styleUrls: ['./package.component.scss']
})

export class PackageComponent implements OnInit,OnChanges {
  @Input() package: Package;
  packagesWithStates: Package[];
  packagesWithSettings: Package[];
  packagesOther: Package[];

  constructor() { }

  ngOnInit() {
    this.buildPackages();
  }

  ngOnChanges(changes) {
    this.buildPackages();
  }

  buildPackages() {
    if (this.package != null && this.package.packages) {
      this.packagesWithStates = this.package.packages.filter((p) => p.tfFiles );
      this.packagesOther = this.package.packages.filter((p) => !p.tfFiles );
    }
  }

}
