import { Observable } from "rxjs";

export class APIFunctions {
    apiAddress: string = "";
    sessionID: string = "";

    constructor(_apiAddress: string, _sessionID: string = "") {
        this.apiAddress = _apiAddress
        if (!this.apiAddress.endsWith('/')) {
            this.apiAddress += '/';
        }

        this.sessionID = _sessionID;
    }

    constructAddress = (endpoint: string) => this.apiAddress + endpoint;

    get(endpoint: string, headers: Headers = new Headers()): Observable<string> {
        var requestAddress = this.constructAddress(endpoint)

        if (this.sessionID != "" && headers.get('sessionID') == null)
            headers.append('sessionID', this.sessionID);

        var requestOptions = {
            method: 'GET',
            headers: headers,
        };

        const observable = new Observable<string>(subscriber => {
            fetch(requestAddress, requestOptions)
                .then(response => response.text().then(x => {
                    if (response.status === 200) {
                        subscriber.next(x)
                        subscriber.complete();
                    }
                    else {
                        subscriber.error(new Error(`Failed Request. Code: ${response.status}`))
                    }
                }))
                .then(result => console.log(result))
                .catch(error => console.log(subscriber.error(new Error(`Failed Request. Error: ${error}`))))
        });
        return observable;
    }

    post(endpoint: string, body: any, headers: Headers = new Headers()): Observable<string> {
        const requestAddress = this.constructAddress(endpoint)

        headers.append("Content-Type", "application/json");

        var requestOptions = {
            body: JSON.stringify(body),
            method: 'POST',
            headers: headers,
        };

        const observable = new Observable<string>(subscriber => {
            fetch(requestAddress, requestOptions)
                .then(response => response.text().then(x => {
                    if (response.status === 200) {
                        subscriber.next(x)
                        subscriber.complete();
                    }
                    else {
                        subscriber.error(new Error(`Failed Request. Code: ${response.status}`))
                    }
                }))
                .then(result => console.log(result))
                .catch(error => console.log(subscriber.error(new Error(`Failed Request. Error: ${error}`))))
        });
        return observable;
    }
}