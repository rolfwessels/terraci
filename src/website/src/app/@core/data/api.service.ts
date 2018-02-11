import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/of';
import 'rxjs/add/operator/map'
import { Http, Response } from '@angular/http';

let counter = 0;

@Injectable()
export class ApiService {
  private apiUrl = 'http://localhost:8000/api'

  constructor(private http: Http) {
  }

  getCurrentState(): Observable<CurrentState> {
    var url = this.urljoin(this.apiUrl,'terra/state');
    return this.http.get(url).map((res: Response) => res.json())
  }

  private urljoin(...urls): string {
    return urls.join("/");
  }
}


export interface Package {
  name: string;
  path: string;
  tfVars: string[];
  tfFiles: string[];
  packages: Package[];
}

export interface CurrentState {
  path: string;
  package: Package;
}



