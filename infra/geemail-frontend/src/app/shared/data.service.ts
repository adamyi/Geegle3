import { Injectable } from "@angular/core";
import { BehaviorSubject, Observable } from "rxjs";

@Injectable()
export class DataService {
    private checkButton = new BehaviorSubject<boolean>(false);
    currentCheck: Observable<any> = this.checkButton.asObservable();

    changeButton(boolean: boolean) {
        this.checkButton.next(boolean);
    }
}
