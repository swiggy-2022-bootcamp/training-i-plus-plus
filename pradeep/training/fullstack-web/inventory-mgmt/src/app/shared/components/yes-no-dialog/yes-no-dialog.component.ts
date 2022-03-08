import { Component, OnInit } from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';

@Component({
  selector: 'app-yes-no-dialog',
  templateUrl: './yes-no-dialog.component.html',
  styleUrls: ['./yes-no-dialog.component.scss']
})
export class YesNoDialogComponent implements OnInit {

  public dialogPromptMessage = 'Are You Sure?';
  public yesMessage = 'Yes';
  public cancelMessage = 'Cancel';

  constructor(public dialogRef: MatDialogRef<YesNoDialogComponent>) {
  }

  ngOnInit() {
  }

}
