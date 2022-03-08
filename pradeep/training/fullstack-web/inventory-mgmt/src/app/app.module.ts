import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './pages/home/home.component';
import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { ProductListComponent } from './pages/inventory/product-list/product-list.component';
import { ProductDetailsComponent } from './pages/inventory/product-details/product-details.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { ReceiptsComponent } from './pages/receipts/receipts.component';
import { ReportingComponent } from './pages/reporting/reporting.component';
import { OkDialogComponent } from './shared/components/ok-dialog/ok-dialog.component';
import { YesNoDialogComponent } from './shared/components/yes-no-dialog/yes-no-dialog.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { SideNavComponent } from './side-nav/side-nav.component';
import { MatSidenavModule } from '@angular/material/sidenav';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './shared/modules/material-module/material-module.module';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    HeaderComponent,
    FooterComponent,
    ProductListComponent,
    ProductDetailsComponent,
    DashboardComponent,
    ReceiptsComponent,
    ReportingComponent,
    OkDialogComponent,
    YesNoDialogComponent,
    SideNavComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    BrowserModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
