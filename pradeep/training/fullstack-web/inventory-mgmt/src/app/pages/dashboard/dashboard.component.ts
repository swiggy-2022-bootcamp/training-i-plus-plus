import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ProductService } from 'src/app/shared/services/product.service';
import { IProduct } from '../../shared/models/product.model'
import {Sort} from '@angular/material/sort';
import {Router} from '@angular/router';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  errMsg: string;
  pageTitle: string = 'Highest Quantity Product List';
  imageWidth: number = 50;
  imageMargin: number = 2;
  showImage: boolean = false;
  errorMessage: string;

  topProducts: any[] = [];
  constructor(private productsService: ProductService, private router: Router) { }

  ngOnInit(): void {
    this.loadProductData();
  }

  loadProductData(): void {
    this.productsService.getProducts().subscribe({
      next: topProducts => {
        this.topProducts = topProducts.sort((a, b) => {
          return compare(a.quantityOnHand, b.quantityOnHand);
        });
        this.topProducts.reverse();
        this.topProducts = topProducts.slice(0,10);
      },
      error: err => this.errMsg = err
    });
  }
}

function compare(a: number | string, b: number | string) {
  return (a < b ? -1 : 1);
}
