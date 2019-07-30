import { Injectable } from "@angular/core";
import { BehaviorSubject } from "rxjs";

@Injectable()
export class DataService {
    private checkButton = new BehaviorSubject<boolean>(false);
    currentCheck = this.checkButton.asObservable();

    changeButton(boolean: boolean) {
        this.checkButton.next(boolean);
    }
}