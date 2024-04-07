# Golang Client for the [Cheapshark API](https://apidocs.cheapshark.com/)
## Examples
### Instantiating a new Client
```
client, err := gocheapshark.NewClient(gocheapshark.NewClientOpts{
  BaseUrl: os.Getenv("CHEAPSHARK_API_BASE_URL"),
})
```
### DealList
```
deals, err := client.DealList(gocheapshark.DealListOpts{
  StoreID:    utils.ValueToPointer("1"),
  UpperPrice: utils.ValueToPointer(uint(15)),
  OnSale:     utils.ValueToPointer(false),
})
```
### DealLookup
```
deal, err := client.DealLookup(gocheapshark.DealLookupOpts{
  ID: "X8sebHhbc1Ga0dTkgg59WgyM506af9oNZZJLU9uSrX8%3D",
})
```
### GameList
```
gameList, err := client.GameList(gocheapshark.GameListOpts{
  Title: utils.ValueToPointer("The Witcher"),
})
```
### GameLookup
```
game, err := client.GameLookup(gocheapshark.GameLookupOpts{
  ID: 612,
})
```
### GameLookupMultiple
```
games, err := client.GameLookupMultiple(gocheapshark.GameLookupMultipleOpts{
  IDs: []string{"612", "128"},
})
```
### StoreList
```
stores, err := client.StoreList()
```
### StoreLastChangedList
```
storesLastChanged, err := client.StoreLastChangeList()
```
### AlertEdit
```
alertEdit, err := client.AlertEdit(gocheapshark.AlertEditOpts{
  Action: "set",
  Email:  "example@example.com",
  GameID: 59,
  Price:  50,
})
```
### AlertManage
```
alertManage, err := client.AlertManage(gocheapshark.AlertManageOpts{
  Action: "manage",
  Email:  "drofferni2@gmail.com",
})
```

NOTES: 
- Anything that Cheapshark API expects a `string` with comma-separated ids, the client expects a `slice` of strings and this will be parsed accordingly for the query params
- Anything that Cheapshark API expects `1` or `0` for flag filters, the client expects a `boolean` and this will be parsed accordingly for the query params
- This follows all the types and payloads available on the [Cheapshark API Docs](https://apidocs.cheapshark.com/), all requests and responses are working, have validations for required fields and error handling from their API.
Since this is a work in progress, it's not guaranteed it will always return the correct HTTP status code when validations or errors happen.
