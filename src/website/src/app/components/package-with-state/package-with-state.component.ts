import { Component, OnInit,Input } from '@angular/core';
import { Package } from '../../@core/data/api.service';

@Component({
  selector: 'app-package-with-state',
  templateUrl: './package-with-state.component.html',
  styleUrls: ['./package-with-state.component.scss']
})
export class PackageWithStateComponent implements OnInit {
  @Input() package: Package;

  constructor() {

  }

  ngOnInit() {  }

}
