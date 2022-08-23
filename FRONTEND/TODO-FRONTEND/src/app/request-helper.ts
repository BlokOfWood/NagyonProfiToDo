import { Observable } from "rxjs";

export class APIFunctions {
    private apiAddress: string = "";

    constructor(_apiAddress: string) {
        this.apiAddress = _apiAddress
        if (!this.apiAddress.endsWith('/')) {
            this.apiAddress += '/';
        }
    }

    constructAddress = (endpoint: string) => this.apiAddress + endpoint;

    get(endpoint: string, headers: Headers = new Headers()): Observable<string> {
        let sessionID = localStorage.getItem('sessionID')

        if (sessionID !== null && headers.get('sessionID') == null)
            headers.append('sessionID', sessionID);

        const requestAddress = this.constructAddress(endpoint)

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
        let sessionID = localStorage.getItem('sessionID')

        if (sessionID !== null && headers.get('sessionID') == null)
            headers.append('sessionID', sessionID);

        const requestAddress = this.constructAddress(endpoint)

        headers.append("Content-Type", "application/json");

        console.log(JSON.stringify(body))
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