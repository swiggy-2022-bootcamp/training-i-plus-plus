import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { ProductListComponent } from './pages/inventory/product-list/product-list.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { ReportingComponent } from './pages/reporting/reporting.component';
import { ReceiptsComponent } from './pages/receipts/receipts.component';
import { ProductDetailsComponent } from './pages/inventory/product-details/product-details.component';


const routes: Routes = [
  {path: 'receipts', component: ReceiptsComponent},
  {path: 'reporting', component: ReportingComponent},
  {path: 'dashboard', component: DashboardComponent},
  {path: 'products', component: ProductListComponent},
  {path: 'products/:id', component: ProductDetailsComponent},
  {path: 'home', component: HomeComponent},
  {path: '', redirectTo: 'home', pathMatch: 'full'},
  {path: '**', redirectTo: 'home', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
