package app

import(
  "lambda-func/api"
  "lambda-func/database"
)


type App struct {
  ApiHandler api.ApiHandler
}


func NewApp()  App {
  // we actually initialize our DB store
  db := database.NewDynamoDBClient()
  // gets pass down to api handler
  apiHandler :=  api.NewApiHandler(db)

  return App{
    ApiHandler : apiHandler ,
  }
}
