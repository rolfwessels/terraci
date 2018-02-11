import { Component, OnInit, Input } from '@angular/core';
import { Package } from '../../@core/data/api.service';

@Component({
  selector: 'app-package-list',
  templateUrl: './package-list.component.html',
  styleUrls: ['./package-list.component.scss']
})

export class PackageListComponent implements OnInit {
  @Input() packages: Package[];

  constructor() { }

  ngOnInit() {
  }

}
