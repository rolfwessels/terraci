import { Component, OnInit,Input,OnChanges } from '@angular/core';
import { Package, PackageState } from '../../@core/data/api.service';
import { ActivatedRoute, Router, NavigationExtras } from '@angular/router';
import { Observable } from 'rxjs/Observable';

@Component({
  selector: 'app-package',
  templateUrl: './package.component.html',
  styleUrls: ['./package.component.scss']
})

export class PackageComponent implements OnInit,OnChanges {
  @Input() package: Package;
  @Input() options: string;
  optionsToPass: string;
  selectedState : PackageState;
  packagesWithStates: Package[];
  packagesWithSettings: Package[];
  packagesOther: Package[];
  selectedOption$: Observable<string>;


  constructor(
    private route: ActivatedRoute,
    private router: Router
  ) { }

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
      this.optionsToPass = this.options

      if (this.package.configOptions != null && this.package.configOptions.length > 1) {
        var keyName = this.getKeyName(this.package.name)
        this.selectedOption$ = this.route.queryParamMap.map(params => params.get(keyName) || this.package.configOptions[0]);
        this.selectedOption$.subscribe(x=> {
          this.optionsToPass = this.options? this.options+"_"+x: x;
        });
      }
    }
  }

  selectOption(option: string) {

    var keyName = this.getKeyName(this.package.name)
    var queryParams = {}
    queryParams[keyName] = option;
    let navigationExtras: NavigationExtras = {
      queryParams: queryParams,
      queryParamsHandling : 'merge'
    };
    console.log(`change ${keyName} to ${option}.`)
    this.router.navigate(['/dashboard'], navigationExtras);
  }

  getKeyName(option: string) : string {
    return option.toLocaleLowerCase().replace(" ","");
  }
}
