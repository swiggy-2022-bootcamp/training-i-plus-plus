import { Component, OnInit } from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';

@Component({
  selector: 'app-ok-dialog',
  templateUrl: './ok-dialog.component.html',
  styleUrls: ['./ok-dialog.component.scss']
})
export class OkDialogComponent implements OnInit {

  public dialogMessage = 'message';
  public okMessage = 'Ok';

  constructor(public dialogRef: MatDialogRef<OkDialogComponent>) {
  }

  ngOnInit() {
  }

}
