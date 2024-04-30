### Getting started

To start this project, make sure you have Go installed on your computer and clone the repository.
Create an json file type with an empty `{}` in it. This will serve as your database.

The command to run this application requires two variables to be passed in with the run command: `FILEPATH` and `PORT`.
Run command: `PORT=<port number> FILEPATH=<file path to json file> go run main.go`

Example of json data:
```
  {
    "username": "testUser",
    "password": "testPassword",
    "habits": [
      {
        "name": "Reading",
        "goal": "Read at least 30 minutes per day",
        "grids": [
          {
            "month": "January",
            "days": [
              {
                "number": 1,
                "completed": true,
                "level": {
                  "level": 1,
                  "text": "Read less than ten minutes",
                  "color": "#A3E4D7"
                }
              },
              {
                "number": 2,
                "completed": false,
                "level": {
                  "level": 0,
                  "text": "",
                  "color": ""
                }
              },
              ...
            ]
            ...
          }
        ],
        "key": {
          "levels": [
            {
              "level": 1,
              "text": "Read less than ten minutes",
              "color": "#A3E4D7"
            },
            {
              "level": 2,
              "text": "Read more than ten minutes",
              "color": "#48C9B0"
            },
            {
              "level": 3,
              "text": "Read 30 minutes or less",
              "color": "#17A589"
            },
            {
              "level": 4,
              "text": "Read more than 30 minutes",
              "color": "#0E6251"
            }
          ]
        }
      }
    ]
  }
```
