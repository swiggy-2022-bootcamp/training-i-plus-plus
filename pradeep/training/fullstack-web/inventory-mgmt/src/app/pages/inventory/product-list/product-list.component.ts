import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ProductService } from 'src/app/shared/services/product.service';
import { IProduct } from '../../../shared/models/product.model'
import {Sort} from '@angular/material/sort';
import {Router} from '@angular/router';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})
export class ProductListComponent implements OnInit {

  errMsg: string;
  pageTitle: string = 'Product List';
  imageWidth: number = 50;
  imageMargin: number = 2;
  showImage: boolean = false;
  errorMessage: string;

  products: IProduct[] = [];
  sortedData: IProduct[] = [];
  constructor(private productsService: ProductService, private router: Router) { }

  ngOnInit(): void {
    this.loadProductData();
  }

  loadProductData(): void {
    this.productsService.getProducts().subscribe({
      next: products => {
        this.products = products;
        this.sortedData = products;
      },
      error: err => this.errMsg = err
    });
  }

  navigateToProduct(product: IProduct){
    this.router.navigate(['/products', product.productId], {state: {product: product}});
  }

  createNew(): void{
    this.router.navigate(['/products/new']);
  }

  sortData(sort: Sort) {
    const data = this.products.slice();
    if (!sort.active || sort.direction === '') {
      this.sortedData = data;
      return;
    }

    this.sortedData = data.sort((a, b) => {
      const isAsc = sort.direction === 'asc';
      switch (sort.active) {
        case 'sku': return compare(a.sku, b.sku, isAsc);
        case 'upc': return compare(a.upc, b.upc, isAsc);
        case 'name': return compare(a.productName, b.productName, isAsc);
        case 'manufacturer': return compare(a.manufacturer, b.manufacturer, isAsc);
        case 'price': return compare(a.pricePerUnit, b.pricePerUnit, isAsc);
        case 'onHand': return compare(a.quantityOnHand, b.quantityOnHand, isAsc);
        default: return 0;
      }
    });
  }
}

function compare(a: number | string, b: number | string, isAsc: boolean) {
  return (a < b ? -1 : 1) * (isAsc ? 1 : -1);
}
