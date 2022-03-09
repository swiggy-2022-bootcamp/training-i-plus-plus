import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient, HttpEvent, HttpErrorResponse, HttpEventType} from '@angular/common/http';
import {IProduct} from '../../shared/models/product.model'
import { map } from  'rxjs/operators';
import { IReceipt } from '../models/receipt.model';



@Injectable({
  providedIn: 'root'
})
export class ProductService {
  API_URL: string = "http://localhost:5000";
  private productUri = 'api/mock/products.json';
  private receiptsUri = 'api/mock/receipts.json';

  constructor(private http: HttpClient) { }

  getProducts(): Observable<IProduct[]> {
    return this.http.get<IProduct[]>(this.productUri)
  }

  getReceipts(): Observable<IReceipt[]> {
    return this.http.get<IReceipt[]>(this.receiptsUri)
  }

  

  public upload(data) {
    return this.http.post<any>(this.API_URL, data, {  
        reportProgress: true,
        observe: 'events'  
      });
  }

}
