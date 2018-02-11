import { Component, OnInit,Input } from '@angular/core';
import { Package } from '../../@core/data/api.service';

@Component({
  selector: 'app-package',
  templateUrl: './package.component.html',
  styleUrls: ['./package.component.scss']
})

export class PackageComponent implements OnInit {
  @Input() package: Package;
  packagesWithStates: Package[];
  packagesWithSettings: Package[];
  packagesOther: Package[];

  constructor() { }

  ngOnInit() {
    console.log(this.package);
    if (this.package != null && this.package.packages) {
      this.packagesWithStates = this.package.packages.filter((p) => p.tfFiles );
      this.packagesWithSettings = this.package.packages.filter((p) => p.tfVars );
      this.packagesOther = this.package.packages.filter((p) => !p.tfFiles && !p.tfVars );
    }
  }

}
