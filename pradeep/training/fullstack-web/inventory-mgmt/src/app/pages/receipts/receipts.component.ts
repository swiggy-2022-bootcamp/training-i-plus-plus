import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { HttpEventType, HttpErrorResponse } from '@angular/common/http';
import { of } from 'rxjs';
import {Sort} from '@angular/material/sort';
import { catchError, map } from 'rxjs/operators';
import { ProductService } from 'src/app/shared/services/product.service';
import { IReceipt } from 'src/app/shared/models/receipt.model';


@Component({
  selector: 'app-receipts',
  templateUrl: './receipts.component.html',
  styleUrls: ['./receipts.component.scss']
})
export class ReceiptsComponent implements OnInit {
  @ViewChild("uploader", { static: false }) uploader: ElementRef;
  constructor(private productsService: ProductService) { }
  ngOnInit(): void {
    this.loadReceiptData();
  }

  receiptUpload: any;
  errMsg: string;
  receipts = [];
  sortedData = [];
  pageTitle: string = " Product Inventory Receipts";


  uploadClick() {
    const uploader = this.uploader.nativeElement; uploader.onchange = () => {
      if(uploader.files.length){
        this.uploadReceipt(uploader.files[0]);
      }
      this.uploader.nativeElement.value = '';
    };
    uploader.click();
  }


  uploadReceipt(file) {
    const formData = new FormData();
    formData.append('file', file);
    file.inProgress = true;
    this.productsService.upload(formData)
    .pipe(
      map(event => {
        switch (event.type) {
          case HttpEventType.Response:
            return event;
        }
      }),
      catchError((error: HttpErrorResponse) => {
        file.inProgress = false;
        return of(`${file.name} upload failed.`);
      })).subscribe((event: any) => {
        if (typeof (event) === 'object') {
          console.log(event.body);
        }
      });
  }

  sortData(sort: Sort) {
    const data = this.receipts.slice();
    if (!sort.active || sort.direction === '') {
      this.sortedData = data;
      return;
    }
    this.sortedData = data.sort((a, b) => {
      const isAsc = sort.direction === 'asc';
      switch (sort.active) {
        case 'name': return compare(a.name, b.name, isAsc);
        case 'uploadDate': return compare(a.uploadDate, b.uploadDate, isAsc);
        default: return 0;
      }
    });
  }


  downloadReciept(receipt: IReceipt): void{
    alert(receipt.receiptId);
  }

  loadReceiptData(): void {
    this.productsService.getReceipts().subscribe({
      next: receipts => {
        this.receipts = receipts;
        this.sortedData = receipts;
      },
      error: err => this.errMsg = err
    });
  }
}


function compare(a: number | string, b: number | string, isAsc: boolean) {
  return (a < b ? -1 : 1) * (isAsc ? 1 : -1);
}
