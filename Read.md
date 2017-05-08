#go-homework

To Build:

(via github)
1. go get github.com/markjohnmagallanes/go-homework
2. cd $GOPATH/src/github.com/markjohnmagallanes/go-homework
3. cd openexchange
4. go install

(via downloaded src.zip)
1. unzip
2. cp -R src/* $GOPATH/src/
3. cd $GOPATH/src/github.com/markjohnmagallanes/go-homework
4. cd openexchange
5. go install

To Run:
1. cd $GOPATH/src/github.com/markjohnmagallanes/go-homework
2. go run main.go


## Notes

Registered apiKey = 5f8245af19724a29bddfc5cd23c6596b

Available for free account
https://openexchangerates.org/api/latest.json?app_id=[APP_APP_ID]


Not Available
https://openexchangerates.org/api/latest.json?app_id=[APP_APP_ID]&base=[CURRENCY_CODE]&callback=someCallbackFunction
https://openexchangerates.org/api/convert/[AMOUNT]/[FROM_CURRENCY]]/[TO_CURRENCY]?app_id=[APP_APP_ID]
